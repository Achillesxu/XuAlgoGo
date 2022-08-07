// Package basics
// Time    : 2022/8/6 16:50
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package basics

import (
	"net"
	"syscall"
	"time"
)

func DialWithTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := net.Dialer{
		Control: func(_, address string, _ syscall.RawConn) error {
			return &net.DNSError{
				Err:         "connection timeout",
				Name:        address,
				Server:      "127.0.0.1",
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
		Timeout: timeout,
	}
	return d.Dial(network, address)
}
