// Package basics
// Time    : 2022/8/7 11:46
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
// describe: show custom flag type in go 1.19
package main

import (
	"flag"
	"fmt"
	"strings"
)

type Person struct {
	FirstName, LastName string
}

func (p *Person) String() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

// UnmarshalText unmarshal text interface to Person
func (p *Person) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return nil
	}
	s := string(text)
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return fmt.Errorf("invalid person: %s, must firstname and lastname", s)
	}
	*p = Person{parts[0], parts[1]}
	return nil
}

func main() {
	var p Person
	defaultVal := &Person{"xushiyin", "yuqing"}
	flag.TextVar(&p, "person", defaultVal, "a person")
	flag.Parse()
	fmt.Println(p)
}
