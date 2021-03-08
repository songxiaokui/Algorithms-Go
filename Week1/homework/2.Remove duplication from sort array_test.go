package homework

import "testing"

/*
@Time    : 2021/3/8 12:54
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 2.Remove duplication from sort array_test.go
@Software: GoLand
*/

func TestRemoveDuplication(t *testing.T) {
	testCase := []struct {
		l      []int
		target int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{}, 0},
		{[]int{0, 1, 1, 3, 3, 3, 6}, 4},
		{[]int{1, 3, 5, 6, 7}, 5},
	}
	for _, v := range testCase {
		data := RemoveDuplication(v.l)
		if data != v.target {
			t.Errorf("data: %#v, target: %#v, error result: %#v", v.l, v.target, data)
		}
	}
}

func TestRemoveDuplication2(t *testing.T) {
	testCase := []struct {
		l      []int
		target int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{}, 0},
		{[]int{0, 1, 1, 3, 3, 3, 6}, 4},
		{[]int{1, 3, 5, 6, 7}, 5},
	}
	for _, v := range testCase {
		data := RemoveDuplication2(v.l)
		if data != v.target {
			t.Errorf("data: %#v, target: %#v, error result: %#v", v.l, v.target, data)
		}
	}
}
