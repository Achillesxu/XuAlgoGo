// Package constants
// Time    : 2022/8/18 10:47
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
//
//go:generate stringer -type=OrderStatus -linecomment
package constants

type OrderStatus int

const (
	CREATE OrderStatus = iota + 1
	PAID
	DELIVERING
	COMPLETE
	CANCELLED
)
