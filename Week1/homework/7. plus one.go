package homework

/*
@Time    : 2021/3/9 13:12
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 7. plus one.go
@Software: GoLand
*/
/*
66. 加一
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
示例 3：

输入：digits = [0]
输出：[1]
*/

func plusOne(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// from last element，if e = 9 then record carry and mod 10 to save current value
	var carry int
	for i := len(nums) - 1; i >= 0; i-- {
		current := nums[i]
		current += carry
		// init carry
		carry = 0
		// last element
		if i == len(nums)-1 {
			current++
		}
		// judge current is 10
		if current == 10 {
			nums[i] = current % 10
			carry = current / 10
		} else {
			nums[i] = current
		}
	}
	// judge carry is 0, extend slice to head
	if carry == 1 {
		nums = append([]int{1}, nums...)
	}
	return nums

}
