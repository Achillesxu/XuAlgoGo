// Package go_req
// Time    : 2022/7/24 16:11
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package go_req

import (
	"context"
	"github.com/carlmjohnson/requests"
)

type placeholder struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func ExampleGetJson() (jBody *placeholder, err error) {
	id := 1
	var body placeholder
	err = requests.
		URL("https://jsonplaceholder.typicode.com").
		Pathf("/posts/%d", id).
		ToJSON(&body).
		Fetch(context.Background())
	jBody = &body
	return
}

func ExamplePostJson() (jBody *placeholder, err error) {
	req := placeholder{
		Title:  "foo",
		Body:   "baz",
		UserID: 1,
	}
	var body placeholder
	err = requests.
		URL("https://jsonplaceholder.typicode.com/posts").
		BodyJSON(req).
		ToJSON(&body).
		Fetch(context.Background())
	jBody = &body
	return
}
