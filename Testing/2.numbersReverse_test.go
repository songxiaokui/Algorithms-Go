package Testing

import (
	"fmt"
	"testing"
)

/*
@Time    : 2021/1/29 15:06
@Author  : austsxk
@Email   : austsxk@163.com
@File    : numbersReverse_test.go
@Software: GoLand
*/

func TestNumbersReverse(t *testing.T) {
	var tests = []struct {
		n int
		t int
	}{
		{123, 321},
		{-123, -321},
		{-120, -21},
		{120, 21},
		{2147483647, 0},
		{2147483641, 1463847412},
		{-2147483647, 0},
		{-2147483617, 0},
	}
	for _, v := range tests {
		result := NumbersReverse(v.n)
		fmt.Println("最后的结果: ", result)
		if result != v.t {
			t.Errorf("%d 反转应该为: %d", v.n, v.t)
		}
		result2 := reverse(v.n)
		fmt.Println("最后的结果: ", result2)
		if result2 != v.t {
			t.Errorf("%d 反转应该为: %d", v.n, v.t)
		}
	}
}
