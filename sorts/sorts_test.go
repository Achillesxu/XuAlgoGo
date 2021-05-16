// Package sorts
// Time    : 2021/5/10 10:54 下午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package sorts

import "testing"

func TestBubbleSort(t *testing.T) {
	var tests = []struct {
		src  []int
		want []int
	}{
		{[]int{1, 3, 4, 5, 2}, []int{1, 2, 3, 4, 5}},
		{[]int{11, 23, 54, 15, 2}, []int{2, 11, 15, 23, 54}},
	}
	for _, test := range tests {
		dst := BubbleSort(test.src)
		isEqual := true
		for idx, item := range dst {
			if test.want[idx] != item {
				isEqual = false
				break
			}
		}
		if !isEqual {
			t.Errorf("BubbleSort(%v) failed", test.src)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	var tests = []struct {
		src  []int
		want []int
	}{
		{[]int{1, 3, 4, 5, 2}, []int{1, 2, 3, 4, 5}},
	}
	for _, test := range tests {
		dst := SelectionSort(test.src)
		isEqual := true
		for idx, item := range dst {
			if test.want[idx] != item {
				isEqual = false
				break
			}
		}
		if !isEqual {
			t.Errorf("SelectionSort(%v) failed", test.src)
		}
	}
}
