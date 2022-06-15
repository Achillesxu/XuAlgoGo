// Package leetcode
// Time    : 2022/6/15 09:14
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package leetcode

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。

这道题最优的做法时间复杂度是 O(n)。

顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，如果找到了，直接返回 2 个数字的下标即可。
如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返回结果。
*/

func towSum(nums []int, target int) []int {
	var result []int
	var mapNums = make(map[int]int)
	for i, v := range nums {
		if _, ok := mapNums[v]; ok {
			result = append(result, mapNums[v], i)
			return result
		}
		mapNums[target-v] = i
	}
	return result
}

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
Example:
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
2 个逆序的链表，要求从低位开始相加，得出结果也逆序输出，返回值是逆序结果链表的头结点。

需要注意的是各种进位问题。
极端情况，例如
Input: (9 -> 9 -> 9 -> 9 -> 9) + (1 -> )
Output: 0 -> 0 -> 0 -> 0 -> 0 -> 1

为了处理方法统一，可以先建立一个虚拟头结点，这个虚拟头结点的 Next 指向真正的 head，
这样 head 不需要单独处理，直接 while 循环即可。
另外判断循环终止的条件不用是 p.Next ！= nil，这样最后一位还需要额外计算，循环终止条件应该是 p != nil。
*/

func addTwoNumbers() {

}
