// Package map2struct
// Time    : 2022/8/14 09:22
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package map2struct

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Extra  map[string]string
}

func PersonMap2Struct(input interface{}) (*Person, error) {
	var person Person
	err := mapstructure.Decode(input, &person)
	if err != nil {
		return nil, errors.Wrap(err, "map2struct.PersonMap2Struct")
	}
	return &person, nil
}

func (p Person) String() string {
	return fmt.Sprintf("%#v", p)
}

type Family struct {
	LastName string `mapstructure:"LastName"`
}

type Location struct {
	City string `mapstructure:"City"`
}

type Person2 struct {
	Family    `mapstructure:"Family,squash"`
	Location  `mapstructure:"Location,squash"`
	FirstName string `mapstructure:"FirstName"`
}

func Person2Map2StructSquash(input interface{}) (*Person2, error) {
	var person Person2
	err := mapstructure.Decode(input, &person)
	if err != nil {
		return nil, errors.Wrap(err, "map2struct.Person2Map2StructSquash")
	}
	return &person, nil
}

func (p Person2) String() string {
	return fmt.Sprintf("%#v", p)
}

type Person3 struct {
	Name string
	Age  int
}

func Person3Map2StructConf(input interface{}) (*Person3, error) {
	// For metadata, we make a more advanced DecoderConfig so we can
	// more finely configure the decoder that is used. In this case, we
	// just tell the decoder we want to track metadata.
	var md mapstructure.Metadata
	var person Person3
	config := &mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &person,
	}
	decoder, _ := mapstructure.NewDecoder(config)

	err := decoder.Decode(input)
	if err != nil {
		return nil, errors.Wrap(err, "map2struct.Person3Map2StructConf")
	}
	return &person, nil
}

func (p Person3) String() string {
	return fmt.Sprintf("%#v", p)
}

func PersonMap2StructWeak(input interface{}) (*Person, error) {
	var person Person

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &person,
	})

	err := decoder.Decode(input)
	if err != nil {
		return nil, errors.Wrap(err, "map2struct.PersonMap2StructWeak")
	}
	return &person, nil
}
