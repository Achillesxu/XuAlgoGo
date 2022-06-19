// Package leetcode
// Time    : 2022/6/15 09:14
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package leetcode

import (
	"strconv"
	"strings"
)

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

type ListNode struct {
	val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	head := ln
	var sb strings.Builder
	for head != nil {
		sb.WriteString(strconv.Itoa(head.val))
		sb.WriteString("->")
		head = head.Next
	}
	return sb.String()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{val: 0}
	n1, n2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.val
			l2 = l2.Next
		}
		current.Next = &ListNode{val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}

/*
Given a string, find the length of the longest substring without repeating characters.
Example 1:
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.

在一个字符串重寻找没有重复字母的最长子串。

这一题和第 438 题，第 3 题，第 76 题，第 567 题类似，用的思想都是"滑动窗口”。

滑动窗口的右边界不断的右移，只要没有重复的字符，就持续向右扩大窗口边界。一旦出现了重复字符，
就需要缩小左边界，直到重复的字符移出了左边界，然后继续移动滑动窗口的右边界。
以此类推，每次移动需要计算当前长度，并判断是否需要更新最大长度，最终最大的值就是题目中的所求。
*/
// lengthOfLongestSubstring0 解法一 位图
func lengthOfLongestSubstring0(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [256]bool
	res, left, right := 0, 0, 0
	for left < len(s) {
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if res < right-left {
			res = right - left
		}
		if right >= len(s) {
			break
		}
	}
	return res
}

// lengthOfLongestSubstring1 解法二 滑动窗口
func lengthOfLongestSubstring1(s string) int {
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1

	for left < len(s) {
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

// lengthOfLongestSubstring2 解法三	 滑动窗口-哈希桶
func lengthOfLongestSubstring2(s string) int {
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if len(s) == 0 {
		return 0
	}
	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))
	for left < len(s) {
		if idx, ok := indexes[s[left]]; ok && idx >= right {
			right = idx + 1
		}
		indexes[s[left]] = left
		left++
		res = max(res, left-right)
	}
	return res
}

/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
You may assume nums1 and nums2 cannot be both empty.
Example 1:
nums1 = [1, 3]
nums2 = [2]
The median is 2.0
Example 2:
nums1 = [1, 2]
nums2 = [3, 4]
The median is (2 + 3)/2 = 2.5

给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。
给出两个有序数组，要求找出这两个数组合并以后的有序数组中的中位数。要求时间复杂度为 O(log (m+n))。
这一题最容易想到的办法是把两个数组合并，然后取出中位数。但是合并有序数组的操作是 O(m+n) 的，不符合题意。
看到题目给的 log 的时间复杂度，很容易联想到二分搜索。
由于要找到最终合并以后数组的中位数，两个数组的总大小也知道，所以中间这个位置也是知道的。
只需要二分搜索一个数组中切分的位置，另一个数组中切分的位置也能得到。为了使得时间复杂度最小，所以二分搜索两个数组中长度较小的那个数组。
关键的问题是如何切分数组 1 和数组 2 。
其实就是如何切分数组 1 。先随便二分产生一个 midA，切分的线何时算满足了中位数的条件呢？
即，线左边的数都小于右边的数，即，nums1[midA-1] ≤ nums2[midB] && nums2[midB-1] ≤ nums1[midA] 。
如果这些条件都不满足，切分线就需要调整。
如果 nums1[midA] < nums2[midB-1]，说明 midA 这条线划分出来左边的数小了，切分线应该右移；如果 nums1[midA-1] > nums2[midB]，
说明 midA 这条线划分出来左边的数大了，切分线应该左移。经过多次调整以后，切分线总能找到满足条件的解。
假设现在找到了切分的两条线了，数组 1 在切分线两边的下标分别是 midA - 1 和 midA。
数组 2 在切分线两边的下标分别是 midB - 1 和 midB。最终合并成最终数组，如果数组长度是奇数，那么中位数就是 max(nums1[midA-1], nums2[midB-1])。
如果数组长度是偶数，那么中间位置的两个数依次是：max(nums1[midA-1], nums2[midB-1]) 和 min(nums1[midA], nums2[midB])，那么中位数就是 (max(nums1[midA-1], nums2[midB-1]) + min(nums1[midA], nums2[midB])) / 2
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	low, high, k, n1Mid, n2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		n1Mid = low + (high-low)>>1
		n2Mid = k - n1Mid
		if n1Mid > 0 && nums1[n1Mid-1] > nums2[n2Mid] {
			high = n1Mid - 1
		} else if n1Mid != len(nums1) && (nums1[n1Mid] < nums2[n2Mid-1]) {
			low = n1Mid + 1
		} else {
			break
		}
	}
	midLeft, midRight := 0, 0
	if n1Mid == 0 {
		midLeft = nums2[n2Mid-1]
	} else if n2Mid == 0 {
		midLeft = nums1[n1Mid-1]
	} else {
		midLeft = max(nums1[n1Mid-1], nums2[n2Mid-1])
	}

	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}

	if n1Mid == len(nums1) {
		midRight = nums2[n2Mid]
	} else if n2Mid == len(nums2) {
		midRight = nums1[n1Mid]
	} else {
		midRight = min(nums1[n1Mid], nums2[n2Mid])
	}
	return float64(midLeft+midRight) / 2.0
}
