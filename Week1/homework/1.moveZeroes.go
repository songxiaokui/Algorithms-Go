package homework

/*
@Time    : 2021/3/8 11:12
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 1.moveZeroes.go
@Software: GoLand
*/

/*
problems description:

Given an array nums, write a function to move all 0's to the end of it while
maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/

func moveZeroes(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}
	// record current not zeroes index
	var k int
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			if k != i {
				nums[i] = 0
			}
			k++
		}
	}
	return nums
}
