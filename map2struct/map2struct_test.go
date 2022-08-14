// Package map2struct
// Time    : 2022/8/14 09:41
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package map2struct

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPersonMap2Struct(t *testing.T) {
	req := require.New(t)
	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"emails": []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mitchellh",
		},
	}
	person, err := PersonMap2Struct(input)
	req.NoError(err)
	fmt.Println(person)
}

func TestPersonMap2Struct2(t *testing.T) {
	req := require.New(t)
	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := map[string]interface{}{
		"name":   123,
		"age":    "bad value",
		"emails": []int{1, 2, 3},
	}
	person, err := PersonMap2Struct(input)
	req.NoError(err)
	fmt.Println(person)
}

func TestPersonMap2StructWeak(t *testing.T) {
	req := require.New(t)
	// This input can come from anywhere, but typically comes from
	// something like decoding JSON, generated by a weakly typed language
	// such as PHP.
	input := map[string]interface{}{
		"name":   123,                      // number => string
		"age":    "42",                     // string => number
		"emails": map[string]interface{}{}, // empty map => empty array
	}
	person, err := PersonMap2StructWeak(input)
	req.NoError(err)
	fmt.Println(person)

}

func TestPerson2Map2StructSquash(t *testing.T) {
	req := require.New(t)
	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := map[string]interface{}{
		"FirstName": "Mitchell",
		"LastName":  "Hashimoto",
		"City":      "San Francisco",
	}

	person, err := Person2Map2StructSquash(input)
	req.NoError(err)
	fmt.Println(person)
}

func TestPerson3Map2StructConf(t *testing.T) {
	req := require.New(t)
	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := map[string]interface{}{
		"name": "Mitchell",
		"age":  91,
	}

	person, err := Person3Map2StructConf(input)
	req.NoError(err)
	fmt.Println(person)
}