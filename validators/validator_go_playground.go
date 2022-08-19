// Package validators
// Time    : 2022/8/17 23:04
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package validators

type person struct {
	Name                string `validate:"required,min=4,max=15"`
	Email               string `validate:"required,email"`
	Age                 int    `validate:"required,numeric,min=18"`
	DriverLicenseNumber string `validate:"omitempty,len=12,numeric"`
}

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
