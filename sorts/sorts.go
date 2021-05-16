// Package sorts
// Time    : 2021/5/10 10:44 下午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package sorts

func BubbleSort(src []int) []int {
	exchange := true
	passNum := len(src) - 1
	for passNum > 0 && exchange {
		exchange = false
		for i := 0; i < passNum; i++ {
			if src[i] > src[i+1] {
				exchange = true
				tmp := src[i]
				src[i] = src[i+1]
				src[i+1] = tmp
			}
		}
		passNum -= 1
	}
	return src
}

func SelectionSort(src []int) []int {
	for i := len(src) - 1; i > 0; i-- {
		posMax := 0
		for j := 1; j <= i; j++ {
			if src[j] > src[posMax] {
				posMax = j
			}
		}
		tmp := src[i]
		src[i] = src[posMax]
		src[posMax] = tmp
	}
	return src
}

func InsertSort(src []int) []int {
	return src
}

func ShellSort(src []int) []int {
	return src
}

func MergeSort(src []int) []int {
	return src
}

func QuickSort(src []int) []int {
	return src
}
