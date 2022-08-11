// Package generics
// Time    : 2022/8/11 09:38
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package generics

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSumIntsOrFloats(t *testing.T) {
	m := map[string]int64{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	n := map[string]float64{
		"a": 1.1,
		"b": 2.2,
		"c": 3.1,
	}
	req := require.New(t)
	req.Equal(int64(6), SumIntsOrFloats(m))
	req.Equal(6.4, SumIntsOrFloats(n))
}
