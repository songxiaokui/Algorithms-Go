package pkg

import "testing"

/*
@Time    : 2021/3/8 19:37
@Author  : austsxk
@Email   : austsxk@163.com
@File    : utility_test.go.go
@Software: GoLand
*/

func TestTwoSliceEqual(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"equal", struct {
			l1 []int
			l2 []int
		}{l1: []int{1, 2, 3}, l2: []int{1, 2, 3}}, true},

		{"x", struct {
			l1 []int
			l2 []int
		}{l1: []int{1, 2, 3}, l2: []int{1, 3, 2}}, false},

		{"null", struct {
			l1 []int
			l2 []int
		}{l1: []int{}, l2: []int{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSliceEqual(tt.args.l1, tt.args.l2); got != tt.want {
				t.Errorf("TwoSliceEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
