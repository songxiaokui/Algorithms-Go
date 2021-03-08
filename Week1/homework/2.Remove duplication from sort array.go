package homework

import "fmt"

/*
@Time    : 2021/3/8 12:54
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 2.Remove duplication from sort array.go
@Software: GoLand
*/

/*
problems description

给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。



示例 1:

给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。
示例 2:

给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。


说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

*/

func RemoveDuplication(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// Save the index of non duplicate elements with variables
	var k int
	for i := 0; i < len(nums); i++ {
		if nums[k] != nums[i] {
			// first element is must equal
			k++
			nums[k] = nums[i]
		}
	}
	fmt.Println("result: ", nums[:k+1])
	return k + 1

}

// use hashMap to save looked element
func RemoveDuplication2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var k int
	hashMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := hashMap[nums[i]]; ok || nums[i] == nums[k] {
			continue
		} else {
			k++
			nums[k] = nums[i]
			hashMap[nums[i]] = true
		}
	}
	fmt.Println("result: ", nums[:k+1])
	return k + 1
}
