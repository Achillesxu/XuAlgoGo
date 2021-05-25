// Package homogeneous_data_structure
// Time    : 2021/5/22 10:51 上午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package homogeneous_data_structure

import "fmt"

func GetIntRowMatrix(n int) []int {
	return make([]int, n)
}

func GetIntColumnMatrix(n int) [][]int {
	var colMatrix = make([][]int, n)
	for i := 0; i < n; i++ {
		colMatrix[i] = make([]int, 1)
	}
	return colMatrix
}

func GetTriangularMatrix(row, col int, val int) [][]int {
	tMatrix := make([][]int, row)
	for i := 0; i < row; i++ {
		tMatrix[i] = make([]int, col)
		for j := 0; j < col; j++ {
			if j >= i {
				tMatrix[i][j] = val
			}
		}
	}
	return tMatrix
}

func GetIdentityMatrix(row, col int, val int) [][]int {
	iMatrix := make([][]int, row)
	for i := 0; i < row; i++ {
		iMatrix[i] = make([]int, col)
		for j := 0; j < col; j++ {
			if j == i {
				iMatrix[i][j] = val
			}
		}
	}
	return iMatrix
}

// PrintZigZag cant understand this func,
func PrintZigZag(n int) []int {
	zigzag := make([]int, n*n)
	for p, i := 1, 0; p <= n*2; p++ {

		var x = p - n
		if x < 0 {
			x = 0
		}

		var y = p - 1
		if y > n-1 {
			y = n - 1
		}

		var j = n*2 - p
		if j > p {
			j = p
		}

		for k := 0; k < j; k++ {
			if p&1 == 0 {
				zigzag[(x+k)*n+y-k] = i
				fmt.Println(x+k, y-k)
			} else {
				zigzag[(y-k)*n+x+k] = i
				fmt.Println(y-k, x+k)
			}
			i++
		}
	}
	return zigzag
}

// PrintSpiral method
func PrintSpiral(n int) []int {
	left := 0
	top := 0
	right := n - 1
	bottom := n - 1
	s := make([]int, n*n)
	i := 0
	for left < right {
		var c int
		for c = left; c <= right; c++ {
			s[top*n+c] = i
			i++
		}
		top++
		var r int
		for r = top; r <= bottom; r++ {
			s[r*n+right] = i
			i++
		}
		right--
		if top == bottom {
			break
		}
		for c = right; c >= left; c-- {
			s[bottom*n+c] = i
			i++
		}
		bottom--
		for r = bottom; r >= top; r-- {
			s[r*n+left] = i
			i++
		}
		left++
	}
	s[top*n+left] = i
	return s
}
