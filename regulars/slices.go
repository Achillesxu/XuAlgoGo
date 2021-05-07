// Package regulars
// Time    : 2021/5/7 9:26 上午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package regulars

import "strings"

func StrSliceJoin(list []string, sep string) string {
	return strings.Join(list, sep)
}
