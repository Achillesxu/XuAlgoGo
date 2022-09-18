// Package main
// Time    : 2022/9/18 12:35
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"context"
	"fmt"
)

type keyType1 string

func tryAnotherKeyType(ctx context.Context, keyToConvert string) {

	k := keyType1(keyToConvert)
	if v := ctx.Value(k); v != nil {
		fmt.Println("found a value for key type 2:", v)
	} else {
		fmt.Println("no value for key type 2")
	}
}

func main() {
	keyString := "foo"

	k := keyType1(keyString)
	ctx := context.WithValue(context.Background(), k, "bar")

	if v := ctx.Value(k); v != nil {
		fmt.Println("found a value for key type 1:", v)
	} else {
		fmt.Println("no value for key type 1")
	}

	tryAnotherKeyType(ctx, keyString)
}
