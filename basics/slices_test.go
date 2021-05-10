// Package basics
// Time    : 2021/5/7 9:31 上午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package basics

import "testing"

func TestStrSliceJoin(t *testing.T) {
	var tests = []struct {
		inputList []string
		inputSep  string
		want      string
	}{
		{[]string{"pizza", "pasta", "sushi"}, ",", "pizza,pasta,sushi"},
		{[]string{"pizza", "pasta", "sushi"}, "$$", "pizza$$pasta$$sushi"},
		{[]string{"pizza", "pasta", "sushi"}, ", ", "pizza, pasta, sushi"},
		{[]string{"yes", "no"}, "--", "yes--no"},
	}

	for _, test := range tests {
		if got := StrSliceJoin(test.inputList, test.inputSep); got != test.want {
			t.Errorf("StrSliceJoin(%q, %q) = %v", test.inputList, test.inputSep, got)
		}
	}
}

func TestGetIntTwoSlice(t *testing.T) {
	var tests = []struct {
		rows     int
		cols     int
		wantRows int
		wantCols int
	}{
		{7, 8, 7, 8},
		{7, 9, 7, 9},
		{9, 10, 9, 10},
	}
	for _, test := range tests {
		twoSlices := GetIntTwoSlice(test.rows, test.cols)
		if len(twoSlices) != test.wantRows || len(twoSlices[0]) != test.wantCols {
			t.Errorf("\"GetIntTwoSlice(%q, %q) slice row or col not right", test.rows, test.cols)
		}
	}
}
