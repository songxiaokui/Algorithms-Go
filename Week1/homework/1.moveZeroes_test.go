package homework

import (
	. "Algorithms-Go/pkg"
	"testing"
)

/*
@Time    : 2021/3/8 11:12
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 1.moveZeroes_test.go
@Software: GoLand
*/

func TestMoveZeroes(t *testing.T) {
	testCase := []struct {
		l      []int
		target []int
	}{
		{[]int{1, 0, 3, 12, 0}, []int{1, 3, 12, 0, 0}},
		{[]int{}, []int{}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{0, 0, 0, 1}, []int{1, 0, 0, 0}},
	}

	for _, v := range testCase {
		d := moveZeroes(v.l)
		// compare two list is equal
		result := TwoSliceEqual(v.target, d)
		if !result {
			t.Errorf("slice: %#v --> target: %#v,"+
				" run error result is: %#v", v.l, v.target, d)
		}
	}
}
