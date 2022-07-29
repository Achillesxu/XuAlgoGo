// Package basics
// Time    : 2022/7/26 11:11
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package basics

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"
)

func CreateTar(tarName string, files []FileBody, isGzip bool) error {
	tFile, err := os.Create(tarName)
	if err != nil {
		return err
	}
	defer tFile.Close()
	var tw *tar.Writer
	if isGzip {
		gz := gzip.NewWriter(tFile)
		tw = tar.NewWriter(gz)
		defer gz.Close()
	} else {
		tw = tar.NewWriter(tFile)
	}
	defer tw.Close()

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Data)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		if _, err := tw.Write(file.Data); err != nil {
			return err
		}
	}
	return nil
}

func ExtractTar(tarName string, isGzip bool) ([]FileBody, error) {
	tFile, err := os.Open(tarName)
	if err != nil {
		return nil, err
	}
	defer tFile.Close()
	var tr *tar.Reader
	if isGzip {
		gz, err := gzip.NewReader(tFile)
		if err != nil {
			return nil, err
		}
		defer gz.Close()
		tr = tar.NewReader(gz)
	} else {
		tr = tar.NewReader(tFile)
	}

	files := make([]FileBody, 0)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		data := make([]byte, hdr.Size)
		buf := bytes.NewBuffer(data)
		if _, err := io.Copy(buf, tr); err != nil {
			return nil, err
		}
		files = append(files, FileBody{hdr.Name, data})
	}
	return files, nil
}
