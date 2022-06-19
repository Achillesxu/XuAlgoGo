// Package leetcode
// Time    : 2022/6/15 09:19
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package leetcode

import (
	"github.com/stretchr/testify/require"
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

func TestAddTwoNumbers(t *testing.T) {
	var tests = []struct {
		l1  *ListNode
		l2  *ListNode
		res *ListNode
	}{
		{&ListNode{val: 2, Next: &ListNode{val: 4, Next: &ListNode{val: 3}}}, &ListNode{val: 5, Next: &ListNode{val: 6, Next: &ListNode{val: 4}}}, &ListNode{val: 7, Next: &ListNode{val: 0, Next: &ListNode{val: 8}}}},
		{&ListNode{val: 2, Next: &ListNode{val: 1, Next: &ListNode{val: 3}}}, &ListNode{val: 5, Next: &ListNode{val: 6, Next: &ListNode{val: 4}}}, &ListNode{val: 7, Next: &ListNode{val: 7, Next: &ListNode{val: 7}}}},
		{nil, &ListNode{val: 5, Next: &ListNode{val: 6, Next: &ListNode{val: 4}}}, &ListNode{val: 5, Next: &ListNode{val: 6, Next: &ListNode{val: 4}}}},
		{&ListNode{val: 5, Next: &ListNode{val: 6, Next: &ListNode{val: 4}}}, nil, &ListNode{val: 5, Next: &ListNode{val: 6, Next: &ListNode{val: 4}}}},
	}
	req4 := require.New(t)
	for _, test := range tests {
		res := addTwoNumbers(test.l1, test.l2)
		resHead := res
		testResHead := test.res
		for res != nil {
			if test.res != nil {
				if res.val != test.res.val {
					t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", test.l1, test.l2, resHead, testResHead)
					break
				}
			} else {
				req4.Nil(test.res)
				break
			}
			res = res.Next
			test.res = test.res.Next
		}
		if test.res != nil {
			t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", test.l1, test.l2, resHead, testResHead)
		}
	}
}

var LengthOfLongestSubstringTests = []struct {
	input string
	res   int
}{
	{"abccbcbb", 3},
	{"bbbbbc", 2},
	{"pwwkew", 3},
	{"", 0},
}

func TestLengthOfLongestSubstring0(t *testing.T) {
	req4 := require.New(t)
	for _, test := range LengthOfLongestSubstringTests {
		if res := lengthOfLongestSubstring0(test.input); res != test.res {
			req4.Equal(res, test.res)
		}
	}
}

func TestLengthOfLongestSubstring1(t *testing.T) {
	req4 := require.New(t)
	for _, test := range LengthOfLongestSubstringTests {
		if res := lengthOfLongestSubstring1(test.input); res != test.res {
			req4.Equal(res, test.res)
		}
	}
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	req4 := require.New(t)
	for _, test := range LengthOfLongestSubstringTests {
		if res := lengthOfLongestSubstring2(test.input); res != test.res {
			req4.Equal(res, test.res)
		}
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	var tests = []struct {
		num1 []int
		num2 []int
		res  float64
	}{
		{[]int{1, 2, 3, 4}, []int{3, 4, 5, 6, 7}, 4.0},
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{}, []int{4, 5, 6}, 5.0},
		{[]int{4, 5}, []int{}, 4.5},
	}
	req4 := require.New(t)
	for _, test := range tests {
		res := findMedianSortedArrays(test.num1, test.num2)
		req4.Equal(res, test.res)
	}
}
