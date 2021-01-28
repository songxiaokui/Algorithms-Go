package Testing

/*
@Time    : 2021/1/29 00:24
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 1.maxSubString.go
@Software: GoLand
*/

// 1.最长子字串

// description: Given a string, find the length of the longest substring without repeating characters.

/*
思路:
O(n):
1. 使用位图或者hashmap
2. 双指针left right
3. 遍历长度，指针右移动，如果不存在，右指针++，并设置为true 如果存在，左指针++，置为false
*/

func lengthOfLongestSubStringBitMap(x string) (int, int, int, string) {
	if len(x) == 0 {
		return 0, 0, 0, ""
	}
	// 定义两个指针left right 和最大值
	var bitMap [128]bool
	var left, right, max = 0, 0, 0
	for right < len(x) {
		// 如果右指针为true,说明存在该点，置为false，并且左指针++
		if bitMap[x[right]] {
			bitMap[x[left]] = false
			left++
		} else {
			bitMap[x[right]] = true
			right++
		}
		// 修改max数值
		if max < right-left {
			max = right - left
		}
		// 判断是否是越界
		if right >= len(x) || left+max >= len(x) {
			left = right - max
			break
		}
	}
	return left, right, max, x[left:right]
}
