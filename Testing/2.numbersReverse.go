package Testing

import (
	"math"
)

/*
@Time    : 2021/1/29 14:37
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 2.numbersReverse.go
@Software: GoLand
*/

// 数字反转

// description: Given a 32-bit signed integer, reverse digits of an integer.

/*
1. 数据反转，要注意是否会溢出
2. int32有符号最大: 2147483648
3. 处理完毕后需要判断

解题思路: 除10取余 在成10相加
*/
func NumbersReverse(number int) int {
	if number == 0 {
		return 0
	}
	var raw, flag = 0, 1
	if number < 0 {
		flag = -1
		number = flag * number
	}
	for number != 0 {
		// 新数据的整数部分
		raw = raw*10 + number%10
		number /= 10
	}
	if raw > math.MaxInt32 {
		return 0
	}
	return flag * raw
}

// 第一次的解执行效果更佳
func reverse(x int) int {
	var num, mid = 0, 0
	for x != 0 {
		mid = x % 10
		num = num*10 + mid
		x = x / 10
	}
	if math.Abs(float64(num)) > math.Pow(2, 31) {
		return 0
	}
	return num
}
