// Package basics
// Time    : 2022/7/26 17:23
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package basics

import "testing"

func TestCreateTar(t *testing.T) {
	var tests = []struct {
		tarName string
		files   []FileBody
		isGzip  bool
	}{
		{"test.tar", []FileBody{{"test.txt", []byte("test 123")}}, false},
		{"test.tar.gz", []FileBody{{"test.txt", []byte("test 456")}}, true},
	}
	for _, test := range tests {
		err := CreateTar(test.tarName, test.files, test.isGzip)
		if err != nil {
			t.Errorf("CreateTar(%q, %q, %t) error: %v", test.tarName, test.files, test.isGzip, err)
		}
	}
}

func TestExtractTar(t *testing.T) {
}
