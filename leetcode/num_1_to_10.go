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
https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0001.Two-Sum/
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
https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0002.Add-Two-Numbers/

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
https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0003.Longest-Substring-Without-Repeating-Characters/
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
https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0004.Median-of-Two-Sorted-Arrays/
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

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 == 0 && len2 == 0 {
		return 0
	}

	numsT := nums1
	if len1 == 0 {
		numsT = nums2
	}

	if len1 == 0 || len2 == 0 {
		len3 := len(numsT)
		div, mod := len3/2, len3%2
		if mod == 0 {
			return float64(numsT[div-1]+numsT[div]) / 2
		} else {
			return float64(numsT[div])
		}

	} else {
		// 思路：
		// 已知nums1和nums2均是有序的，可设计一种切换遍历的机制。
		// 通过div和mod来判断切换遍历提前结束，并计算中位数。
		//
		// 结束遍历的判断条件：
		// 	遍历次数count 等于 div
		// 计算中位数：
		// 	1、mod等于0时，中位数等于当前遍历的数；
		//	2、mod等于1时，中位数等于(前一个遍历的数+当期遍历的数)/2；
		//
		div := (len1 + len2) / 2
		mod := (len1 + len2) % 2
		n := 0 // nums1_index
		m := 0 // nums2_index
		count := 0
		pre := 0
		curr := 0

		for {
			// 条件1：nums1已经遍历完，但nums2还没有遍历完，numsT切片切换到nums2。
			// 条件2：nums2已经遍历完，但nums1还没有遍历完，numsT切片切换到nums1。
			// 条件3：nums1[n]大于等于nums2[n], numsT切片切换到nums2。
			// 条件4：nums1[n]小于nums2[n], numsT切片切换到nums1。
			if n >= len1 && m < len2 {
				numsT = nums2[m:]
				m++
			} else if m >= len2 && n < len1 {
				numsT = nums1[n:]
				n++
			} else if nums1[n] >= nums2[m] {
				numsT = nums2[m:]
				m++
			} else if nums1[n] < nums2[m] {
				numsT = nums1[n:]
				n++
			}

			curr = numsT[0] // 取出0号值
			count++         // 遍历了一次
			if (count - 1) == div {
				break
			}
			pre = curr
		}

		// 计算中位数
		if mod == 0 {
			return float64(pre+curr) / 2
		} else {
			return float64(curr)
		}
	}
}

/*
https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0005.Longest-Palindromic-Substring/
Given a string s, return the longest palindromic substring in s
Example 1:
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:
Input: "cbbd"
Output: "bb"
Example 3:
Input: "a"
Output: "a"
Example 4:
Input: "ac"
Output: "a"
Constraints:
1 <= s.length <= 1000
s contains only lowercase English letters.

给你一个字符串 s，找到 s 中最长的回文子串。

解法一，动态规划。定义 dp[i][j] 表示从字符串第 i 个字符到第 j 个字符这一段子串是否是回文串。
由回文串的性质可以得知，回文串去掉一头一尾相同的字符以后，剩下的还是回文串。
所以状态转移方程是 dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])，
注意特殊的情况，j - i == 1 的时候，即只有 2 个字符的情况，只需要判断这 2 个字符是否相同即可。
j - i == 2 的时候，即只有 3 个字符的情况，只需要判断除去中心以外对称的 2 个字符是否相等。每次循环动态维护保存最长回文串即可。
时间复杂度 O(n^2)，空间复杂度 O(n^2)。

解法二，中心扩散法。动态规划的方法中，我们将任意起始，终止范围内的字符串都判断了一遍。
其实没有这个必要，如果不是最长回文串，无需判断并保存结果。
所以动态规划的方法在空间复杂度上还有优化空间。判断回文有一个核心问题是找到“轴心”。
如果长度是偶数，那么轴心是中心虚拟的，如果长度是奇数，那么轴心正好是正中心的那个字母。
中心扩散法的思想是枚举每个轴心的位置。然后做两次假设，假设最长回文串是偶数，
那么以虚拟中心往 2 边扩散；假设最长回文串是奇数，那么以正中心的字符往 2 边扩散。
扩散的过程就是对称判断两边字符是否相等的过程。这个方法时间复杂度和动态规划是一样的，
但是空间复杂度降低了。时间复杂度 O(n^2)，空间复杂度 O(1)。

解法三，滑动窗口。这个写法其实就是中心扩散法变了一个写法。中心扩散是依次枚举每一个轴心。
滑动窗口的方法稍微优化了一点，有些轴心两边字符不相等，下次就不会枚举这些不可能形成回文子串的轴心了。
不过这点优化并没有优化时间复杂度，时间复杂度 O(n^2)，空间复杂度 O(1)。

解法四，马拉车算法。这个算法是本题的最优解，也是最复杂的解法。时间复杂度 O(n)，空间复杂度 O(n)。
中心扩散法有 2 处有重复判断，第一处是每次都往两边扩散，不同中心扩散多次，实际上有很多重复判断的字符，
能否不重复判断？第二处，中心能否跳跃选择，不是每次都枚举，是否可以利用前一次的信息，跳跃选择下一次的中心？
马拉车算法针对重复判断的问题做了优化，增加了一个辅助数组，将时间复杂度从 O(n^2) 优化到了 O(n)，
空间换了时间，空间复杂度增加到 O(n)。
*/

// 解法一 Manacher's algorithm，时间复杂度 O(n)，空间复杂度 O(n)
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	newS := make([]rune, 0)
	newS = append(newS, '#')
	for _, c := range s {
		newS = append(newS, c)
		newS = append(newS, '#')
	}
	// dp[i]:    以预处理字符串下标 i 为中心的回文半径(奇数长度时不包括中心)
	// maxRight: 通过中心扩散的方式能够扩散的最右边的下标
	// center:   与 maxRight 对应的中心字符的下标
	// maxLen:   记录最长回文串的半径
	// begin:    记录最长回文串在起始串 s 中的起始下标
	dp, maxRight, center, maxLen, begin := make([]int, len(newS)), 0, 0, 1, 0
	for i := 0; i < len(newS); i++ {
		if i < maxRight {
			// 这一行代码是 Manacher 算法的关键所在
			dp[i] = min(maxRight-i, dp[2*center-i])
		}
		// 中心扩散法更新 dp[i]
		left, right := i-(1+dp[i]), i+(1+dp[i])
		for left >= 0 && right < len(newS) && newS[left] == newS[right] {
			dp[i]++
			left--
			right++
		}
		// 更新 maxRight，它是遍历过的 i 的 i + dp[i] 的最大者
		if i+dp[i] > maxRight {
			maxRight = i + dp[i]
			center = i
		}
		// 记录最长回文子串的长度和相应它在原始字符串中的起点
		if dp[i] > maxLen {
			maxLen = dp[i]
			begin = (i - maxLen) / 2 // 这里要除以 2 因为有我们插入的辅助字符 #
		}
	}
	return s[begin : begin+maxLen]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 解法二 滑动窗口，时间复杂度 O(n^2)，空间复杂度 O(1)
func longestPalindrome1(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		// 移动到相同字母的最右边（如果有相同字母）
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		// 找到回文的边界
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		// 重置到下一次寻找回文的中心
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}

// 解法三 中心扩散法，时间复杂度 O(n^2)，空间复杂度 O(1)
func longestPalindrome2(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}

func maxPalindrome(s string, i, j int, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}

// 解法四 DP，时间复杂度 O(n^2)，空间复杂度 O(n^2)
func longestPalindrome3(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

/*
https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0006.ZigZag-Conversion/
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
(you may want to display this pattern in a fixed font for better legibility)
P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"
Write the code that will take a string and make this conversion given a number of rows:
string convert(string text, int nRows);
Example 1:
Input: text = "PAYPALISHIRING", nRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:
Input: text = "PAYPALISHIRING", nRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
Example 3:
Input: text = "A", nRows = 1
Output: "A"
Constraints:
1 <= text.length <= 100
s consists of English letters (lower-case and upper-case), ',' and '.'
1 <= numRows <= 1000
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

这一题没有什么算法思想，考察的是对程序控制的能力。用 2 个变量保存方向，当垂直输出的行数达到了规定的目标行数以后，需要从下往上转折到第一行，循环中控制好方向ji
*/

func convert(s string, numRows int) string {
	return ""
}
