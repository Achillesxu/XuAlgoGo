// Package validators
// Time    : 2022/8/17 23:20
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package validators

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"regexp"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

func (a Address) Validate() error {
	return validation.ValidateStruct(&a,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Street, validation.Required, validation.Length(5, 50)),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&a.City, validation.Required, validation.Length(5, 50)),
		// State cannot be empty, and must be a string consisting of two letters in upper case
		validation.Field(&a.State, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		// State cannot be empty, and must be a string consisting of five digits
		validation.Field(&a.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	)
}

type Customer struct {
	Name    string  `json:"name"`
	Gender  string  `json:"gender"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

func (c Customer) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, validation.Length(5, 20)),
		validation.Field(&c.Gender, validation.Required, validation.In("Female", "Male")),
		validation.Field(&c.Email, validation.Required, is.Email),
		validation.Field(&c.Address),
	)
}

func ValidateMap(m map[string]interface{}, mapRule validation.MapRule) error {
	return validation.Validate(m, mapRule)
}

type Eps struct {
	Name   string   `json:"name"`
	Emails []string `json:"emails"`
}

func (e Eps) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Name, validation.Required, validation.Length(5, 20)),
		validation.Field(&e.Emails, validation.Each(is.Email)),
	)
}

type Employee struct {
	Name string
}

type Manager struct {
	Employee
	Level int
}
