// Package basics
// Time    : 2022/8/6 16:55
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package basics

import (
	"context"
	"io"
	"net"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestListener(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = l.Close() }()
	t.Logf("listening on %s", l.Addr())
	for {
		c, err := l.Accept()
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			defer func() { _ = c.Close() }()
			t.Logf("accepted %s", c.RemoteAddr())
			_, err := c.Write([]byte("hello"))
			if err != nil {
				t.Error(err)
				return
			}
		}()
	}
}

func TestDial(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	done := make(chan struct{})
	go func() {

		defer func() { done <- struct{}{} }()
		for {
			t.Log("listening on", l.Addr())
			conn, err := l.Accept()
			if err != nil {
				t.Log(err)
				return
			}
			go func(c net.Conn) {
				defer func() {
					_ = c.Close()
					done <- struct{}{}
				}()
				b := make([]byte, 1024)
				for {
					n, err := c.Read(b)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}
					t.Logf("received: %q", b[:n])
				}
			}(conn)
		}
	}()
	conn, err := net.Dial("tcp", l.Addr().String())
	_, err = conn.Write([]byte("hello"))
	if err != nil {
		t.Fatal(err)
	}

	_ = conn.Close()

	<-done
	_ = l.Close()
	<-done
}

func TestDialTimeout(t *testing.T) {
	c, err := DialWithTimeout("tcp", "10.0.0.0:http", time.Second*5)
	if err == nil {
		_ = c.Close()
		t.Fatal("connection not timeout")
	}
	nErr, ok := err.(net.Error)
	if !ok {
		t.Fatal("error is not net.Error")
	}
	if !nErr.Timeout() {
		t.Fatal("error is not timeout")
	}
}

func TestDialWithContext(t *testing.T) {
	dl := time.Now().Add(time.Second * 5)
	ctx, cancel := context.WithDeadline(context.Background(), dl)
	defer cancel()

	var d net.Dialer
	d.Control = func(_, _ string, _ syscall.RawConn) error {
		time.Sleep(time.Second * 6)
		return nil
	}
	conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:80")
	if err == nil {
		_ = conn.Close()
		t.Fatal("connection not timeout")
	}
	nErr, ok := err.(net.Error)
	if !ok {
		t.Fatal(err, "error is not net.Error")
	} else {
		if !nErr.Timeout() {
			t.Fatal(err, "error is not timeout")
		}
	}
	if ctx.Err() != context.DeadlineExceeded {
		t.Fatal(ctx.Err(), "error is not DeadlineExceeded")
	}
}

func TestDialContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	sync := make(chan struct{})

	go func() {
		defer func() { sync <- struct{}{} }()
		var d net.Dialer
		d.Control = func(_, _ string, _ syscall.RawConn) error {
			time.Sleep(time.Second)
			return nil
		}
		conn, err := d.DialContext(ctx, "tcp", "10.0.0.1:80")
		if err != nil {
			t.Log(err)
			return
		}
		_ = conn.Close()
		t.Error("connection not timeout")
	}()
	cancel()
	<-sync
	if ctx.Err() != context.Canceled {
		t.Fatal(ctx.Err(), "error is not Canceled")
	}
}

func TestDialContextCancelFanOut(t *testing.T) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*10))
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	go func() {
		conn, err := listener.Accept()
		if err == nil {
			_ = conn.Close()
		}
	}()

	dial := func(ctx context.Context, address string, response chan int, id int, wg *sync.WaitGroup) {
		defer wg.Done()
		var d net.Dialer
		c, err := d.DialContext(ctx, "tcp", address)
		if err != nil {
			return
		}
		_ = c.Close()
		select {
		case <-ctx.Done():
		case response <- id:
		}
	}
	res := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go dial(ctx, listener.Addr().String(), res, i+1, &wg)
	}
	response := <-res
	cancel()
	wg.Wait()
	close(res)

	if ctx.Err() != context.Canceled {
		t.Errorf("error is not Canceled: %v", ctx.Err())
	}

	t.Logf("dialer %d retieved the response", response)
}
