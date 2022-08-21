// Package validators
// Time    : 2022/8/17 23:04
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package validators

import "database/sql"

// User contains user information
type User struct {
	FirstName      string  `json:"fname" validate:"required"`
	LastName       string  `json:"lname" validate:"required"`
	Age            uint8   `validate:"gte=0,lte=130"`
	Email          string  `json:"e-mail" validate:"required,email"`
	FavouriteColor string  `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Addr `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Addr houses a users address information
type Addr struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// DbBackedUser User struct
type DbBackedUser struct {
	Name sql.NullString `validate:"required,gte=0"`
	Age  sql.NullInt64  `validate:"required"`
}
