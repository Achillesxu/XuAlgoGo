// Package validators
// Time    : 2022/8/20 09:57
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package validators

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"reflect"
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	v := validator.New()
	address := &Addr{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
		City:   "12",
	}
	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            10,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "rgb",
		Addresses:      []*Addr{address},
	}
	err := v.Struct(user)
	require.NoError(t, err)
}

func TestStringVar(t *testing.T) {
	myEmail := "joeybloggs@gmail.com"

	v := validator.New()

	err := v.Var(myEmail, "required,email")

	require.NoError(t, err)
}

// ValidateValuer implements validator.CustomTypeFunc
func ValidateValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err == nil {
			return val
		}
		// handle the error how you want
	}
	return nil
}

// TestCustomValidate NullInt64 has error!!!!!!!!!!
func TestCustomValidate(t *testing.T) {
	v := validator.New()

	v.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})

	x := DbBackedUser{Name: sql.NullString{String: "123", Valid: true}, Age: sql.NullInt64{Int64: 0, Valid: true}}

	err := v.Struct(x)

	require.NoError(t, err)
}

func TestStructLevelValidate(t *testing.T) {
	v := validator.New()

	// register function to get tag name from json tags.
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		names := strings.SplitN(fld.Tag.Get("json"), ",", 2)
		t.Log(names)
		name := names[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// register validation for 'User'
	// NOTE: only have to register a non-pointer type for 'User', validator
	// internally dereferences during it's type checks.
	// UserStructLevelValidation contains custom struct level validations that don't always
	// make sense at the field validation level. For Example this function validates that either
	// FirstName or LastName exist; could have done that with a custom field validation but then
	// would have had to add it to both fields duplicating the logic + overhead, this way it's
	// only validated once.
	//
	// NOTE: you may ask why wouldn't I just do this outside of validator, because doing this way
	// hooks right into validator and you can combine with validation tags and still have a
	// common error output format.
	v.RegisterStructValidation(func(sl validator.StructLevel) {
		user := sl.Current().Interface().(User)
		if len(user.FirstName) == 0 && len(user.LastName) == 0 {
			sl.ReportError(user.FirstName, "fname", "FirstName", "fnameorlname", "")
			sl.ReportError(user.LastName, "lname", "LastName", "fnameorlname", "")
		}
		// plus can do more, even with different tag than "fnameorlname"
	}, User{})

	// build 'User' info, normally posted data etc...
	address := &Addr{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
		City:   "Unknown",
	}

	user := &User{
		FirstName:      "",
		LastName:       "",
		Age:            45,
		Email:          "Badger.Smith@gmail",
		FavouriteColor: "#000",
		Addresses:      []*Addr{address},
	}
	// returns InvalidValidationError for bad validation input, nil or ValidationErrors ( []FieldError )
	err := v.Struct(user)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			require.NoError(t, err)
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.Field())     // by passing alt name to ReportError like below
			// fmt.Println(err.StructNamespace())
			// fmt.Println(err.StructField())
			// fmt.Println(err.Tag())
			// fmt.Println(err.ActualTag())
			// fmt.Println(err.Kind())
			// fmt.Println(err.Type())
			// fmt.Println(err.Value())
			// fmt.Println(err.Param())
			// fmt.Println()
		}
		// from here you can create your own error messages in whatever language you wish
	}
}
