package homework

import (
	. "Algorithms-Go/pkg"
	"testing"
)

/*
@Time    : 2021/3/8 16:17
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 3.Rotate Array_test.go.go
@Software: GoLand
*/

func TestRotateArray(t *testing.T) {
	type args struct {
		nums   []int
		k      int
		target []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "RotateArray1", args: struct {
			nums   []int
			k      int
			target []int
		}{
			nums: []int{1, 2, 3, 4, 5}, k: 1, target: []int{5, 1, 2, 3, 4}},
		},
		{name: "RotateArray2", args: struct {
			nums   []int
			k      int
			target []int
		}{
			nums: []int{1, 2, 3, 4, 5}, k: 0, target: []int{1, 2, 3, 4, 5}},
		},
		{name: "RotateArray3", args: struct {
			nums   []int
			k      int
			target []int
		}{
			nums: []int{1, 2, 3, 4, 5}, k: 5, target: []int{1, 2, 3, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RotateArray(tt.args.nums, tt.args.k)
			if !TwoSliceEqual(tt.args.nums, tt.args.target) {
				t.Errorf("test error")
			}
		})
	}
}
