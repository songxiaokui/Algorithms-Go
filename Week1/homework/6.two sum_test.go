package homework

import (
	"reflect"
	"testing"
)

/*
@Time    : 2021/3/9 13:08
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 6.two sum_test.go.go
@Software: GoLand
*/

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", struct {
			nums   []int
			target int
		}{nums: []int{2, 7, 11, 15}, target: 9}, []int{0, 1}},

		{"case2", struct {
			nums   []int
			target int
		}{nums: []int{2, 7, 11, 15}, target: 33}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
