// Package basics
// Time    : 2022/7/26 10:15
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
// description : Go语言官方库：os、io/iouti(弃用)、bufio涵盖了文件操作的所有场景
// os库中的方法对文件都是直接的IO操作，频繁的IO操作会增加CPU的中断频率，所以我们可以使用内存缓存区来减少IO操作，
// 在写字节到硬盘前使用内存缓存，当内存缓存区的容量到达一定数值时在写内存数据buffer到硬盘
package basics

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CreateFile(fileName string) (*os.File, error) {
	f, err := os.Create(fileName)
	return f, err
}

func GetFileInfo(fileName string) (os.FileInfo, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}
	return fInfo, nil
}

func FileChMod(fileName string, mode os.FileMode) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	err = f.Chmod(mode)
	if err != nil {
		return err
	}
	return nil
}

func FileChown(fileName string, uid, gid int) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	err = f.Chown(uid, gid)
	if err != nil {
		return err
	}
	return nil
}

func RmFile(fileName string) error {
	err := os.Remove(fileName)
	return err
}

func WriteAll(fileName string, content []byte) error {
	err := os.WriteFile(fileName, content, 0644)
	return err
}

func WriteLine(fileName string, content []byte) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(content)
	_, err = f.WriteString("\n")
	return err
}

func WriteLine2(fileName string, content []byte) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	bufWriter := bufio.NewWriter(f)
	_, err = bufWriter.Write(content)
	if err != nil {
		return err
	}
	err = bufWriter.Flush()
	if err != nil {
		return err
	}
	return nil
}

func WriteAt(fileName string, content []byte, offset int64) error {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteAt(content, offset)
	return err
}

func WriteBuffer(fileName string, content []byte) error {
	f, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	bufWriter := bufio.NewWriter(f)
	_, err = bufWriter.Write(content)
	if err != nil {
		return err
	}
	unFlushBufSize := bufWriter.Buffered()
	log.Printf("Bytes buf : %d\n", unFlushBufSize)

	byteAvailable := bufWriter.Available()
	log.Printf("Bytes available : %d\n", byteAvailable)

	err = bufWriter.Flush()
	if err != nil {
		return err
	}
	return nil
}

func ReadFileAll(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ReadLine(fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bufReader := bufio.NewReader(f)
	data, err := bufReader.ReadBytes('\n')
	line := strings.TrimSpace(string(data))
	if err != nil && err != io.EOF {
		return nil, err
	}

	return []byte(line), nil
}

func ReadBytes(fileName string, size int) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bufReader := bufio.NewReader(f)
	buf := make([]byte, size)
	_, err = bufReader.Read(buf)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return buf, nil
}

func ReadScanWord(fileName string, size int) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	success := scanner.Scan()
	if !success {
		err = scanner.Err()
		if err != nil {
			return nil, err
		} else {
			return scanner.Bytes(), nil
		}
	} else {
		return scanner.Bytes(), nil
	}
}

func UnZipFiles(zipName string) error {
	zipReader, err := zip.OpenReader(zipName)
	if err != nil {
		return err
	}
	defer zipReader.Close()
	for _, file := range zipReader.File {
		rc, err := file.Open()
		if err != nil {
			rc.Close()
			return err
		}
		f, err := os.Create(file.Name)
		if err != nil {
			f.Close()
			return err
		}
		_, err = io.Copy(f, rc)
		if err != nil {
			return err
		}
		fmt.Println(file.Name)
	}
	return nil
}

func ZipFiles(zipName string, zFiles []FileBody) error {
	zFile, err := os.Create(zipName)
	if err != nil {
		return err
	}
	defer zFile.Close()
	zipWriter := zip.NewWriter(zFile)

	for _, file := range zFiles {
		zWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			return err
		}
		_, err = zWriter.Write(file.Data)
		if err != nil {
			return err
		}
	}
	return nil
}
