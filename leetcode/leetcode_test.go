// Package leetcode
// Time    : 2022/6/15 09:19
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		result []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10}, -11, []int{4, 5}},
	}
	for _, test := range tests {
		if result := towSum(test.nums, test.target); !reflect.DeepEqual(result, test.result) {
			t.Errorf("towSum(%v, %v) = %v, want %v", test.nums, test.target, result, test.result)
		}
	}
}
