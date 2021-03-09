package homework

import (
	"reflect"
	"testing"
)

/*
@Time    : 2021/3/9 13:28
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 7. plus one_test.go.go
@Software: GoLand
*/

func Test_plusOne(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", struct{ nums []int }{nums: []int{1, 2, 3}}, []int{1, 2, 4}},
		{"case2", struct{ nums []int }{nums: []int{4, 2, 9}}, []int{4, 3, 0}},
		{"case3", struct{ nums []int }{nums: []int{5, 9, 9}}, []int{6, 0, 0}},
		{"case4", struct{ nums []int }{nums: []int{9, 9, 9}}, []int{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
