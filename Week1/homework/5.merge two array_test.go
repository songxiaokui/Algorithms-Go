package homework

import (
	"reflect"
	"testing"
)

/*
@Time    : 2021/3/8 23:13
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 5.merge two array_test.go.go
@Software: GoLand
*/

func Test_merge2(t *testing.T) {
	type args struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		target []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", struct {
			nums1  []int
			m      int
			nums2  []int
			n      int
			target []int
		}{nums1: []int{1, 2, 2, 4, 0, 0, 0},
			m:      4,
			nums2:  []int{3, 7, 8},
			n:      3,
			target: []int{1, 2, 2, 3, 4, 7, 8}}},

		{"case2", struct {
			nums1  []int
			m      int
			nums2  []int
			n      int
			target []int
		}{nums1: []int{1, 2, 2, 4},
			m:      4,
			nums2:  []int{},
			n:      0,
			target: []int{1, 2, 2, 4}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge2(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.args.target) {
				t.Errorf(tt.name, "test error!")
			}
		})
	}
}
