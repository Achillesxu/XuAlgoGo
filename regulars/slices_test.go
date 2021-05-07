// Package regulars
// Time    : 2021/5/7 9:31 上午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package regulars

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
