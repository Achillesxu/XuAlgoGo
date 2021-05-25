// Package XuAlgoGo
// Time    : 2021/5/6 2:10 下午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package main

import (
	"fmt"
	hds "github.com/Achillesxu/XuAlgoGo/homogeneous_data_structure"
)

const (
	example1 = "this is a example"
	example2 = "second example"
)

// main project entrance
func main() {
	spiral := hds.PrintZigZag(4)
	fmt.Printf("%v", spiral)

}
