// Package basics
// Time    : 2021/5/7 9:26 上午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package basics

import "strings"

func StrSliceJoin(list []string, sep string) string {
	return strings.Join(list, sep)
}

func GetIntTwoSlice(rows int, cols int) [][]int {
	var twoSlices = make([][]int, rows)
	for i := range twoSlices {
		twoSlices[i] = make([]int, cols)
	}
	return twoSlices
}
