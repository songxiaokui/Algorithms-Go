package homework

/*
@Time    : 2021/3/17 19:11
@Author  : austsxk
@Email   : austsxk@163.com
@File    : majorityElement.go
@Software: GoLand
*/
/*
169. 多数元素
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1：

输入：[3,2,3]
输出：3
示例 2：
输入：[2,2,1,1,1,2,2]
输出：2
*/

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	hashMap := make(map[int]int)
	for _, i := range nums {
		if _, ok := hashMap[i]; ok {
			hashMap[i]++
		} else {
			hashMap[i] = 1
		}
	}

	for k, v := range hashMap {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}
