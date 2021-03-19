package homework

import "testing"

/*
@Time    : 2021/3/19 14:23
@Author  : austsxk
@Email   : austsxk@163.com
@File    : majorityElement_test.go.go
@Software: GoLand
*/

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", struct{ nums []int }{nums: []int{1, 2, 2, 2, 1, 2, 1}}, 2},
		{"case2", struct{ nums []int }{nums: []int{3, 2, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
