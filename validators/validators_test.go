// Package validators
// Time    : 2022/8/17 23:13
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package validators

import (
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestStrValidator(t *testing.T) {
	req := require.New(t)
	data := "exa"
	err := validation.Validate(data,
		validation.Required,       // not empty
		validation.Length(5, 100), // length between 5 and 100
	)
	fmt.Println(err)
	req.Error(err)
}

func TestMapValidator(t *testing.T) {
	m := map[string]interface{}{
		"Name":  "Qiang Xue",
		"Email": "q@qq.com",
		"Address": map[string]interface{}{
			"Street": "123456",
			"City":   "Unknown",
			"State":  "Virginia",
			"Zip":    "12345",
		},
	}
	t.Logf("%#v", m)
	infoRule := validation.Map(
		// Name cannot be empty, and the length must be between 5 and 20.
		validation.Key("Name", validation.Required, validation.Length(5, 20)),
		// Email cannot be empty and should be in a valid email format.
		validation.Key("Email", validation.Required, is.Email),
		// Validate Address using its own validation rules
		validation.Key("Address", validation.Map(
			// Street cannot be empty, and the length must between 5 and 50
			validation.Key("Street", validation.Required, validation.Length(5, 50)),
			// City cannot be empty, and the length must between 5 and 50
			validation.Key("City", validation.Required, validation.Length(5, 50)),
			// State cannot be empty, and must be a string consisting of two letters in upper case
			validation.Key("State", validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Z]{2,20}$"))),
			// State cannot be empty, and must be a string consisting of five digits
			validation.Key("Zip", validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
		)),
	)
	err := ValidateMap(m, infoRule)
	require.NoError(t, err)
}

func TestCustomerValidate(t *testing.T) {
	c := Customer{
		Name:   "Qiang Xue",
		Email:  "q@qq.com",
		Gender: "Male",
		Address: Address{
			Street: "123 Main Street",
			City:   "Unknown",
			State:  "VA",
			Zip:    "12345",
		},
	}
	err := c.Validate()
	require.NoError(t, err)
}

func TestAddresses(t *testing.T) {
	addresses := []Address{
		{State: "MD", Zip: "12345"},
		{Street: "123 Main St", City: "Vienna", State: "VA", Zip: "12345"},
		{City: "Unknown", State: "NC", Zip: "123"},
	}
	err := validation.Validate(addresses)
	require.NoError(t, err)
}

// TestArticleValidate has error type of email
func TestArticleValidate(t *testing.T) {
	c := Eps{
		Name: "Qiang Xue",
		Emails: []string{
			"valid@example.com",
			"invalid",
		},
	}
	err := c.Validate()
	require.Error(t, err)
}
