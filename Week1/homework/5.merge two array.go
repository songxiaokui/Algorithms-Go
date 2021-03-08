package homework

import "sort"

/*
@Time    : 2021/3/8 19:04
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 5.merge two array.go
@Software: GoLand
*/

/*
88. 合并两个有序数组
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它
就有足够的空间保存来自 nums2 的元素。

示例 1：

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
示例 2：

输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
*/

// use plus two list, then sort
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}

// use double pointer merge slice
func merge2(nums1 []int, m int, nums2 []int, n int) {
	// 1. copy a [:m] slice
	nums1Copy := make([]int, m)
	// copy
	copy(nums1Copy, nums1)
	var p1, p2, index int
	for p1 < m && p2 < n {
		if nums1Copy[p1] < nums2[p2] {
			nums1[index] = nums1Copy[p1]
			p1++
		} else {
			nums1[index] = nums2[p2]
			p2++
		}
		index++
	}
	if p1 < m {
		copy(nums1[index:], nums1Copy[p1:])
	}

	if p2 < n {
		copy(nums1[index:], nums2[p2:])
	}
}
