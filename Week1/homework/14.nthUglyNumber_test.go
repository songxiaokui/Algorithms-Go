package homework

import "testing"

/*
@Time    : 2021/3/12 17:06
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 14.nthUglyNumber_test.go.go
@Software: GoLand
*/

func Test_nthUglyNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", struct{ n int }{n: 10}, 12},
		{"case2", struct{ n int }{n: 1}, 1},
		{"case3", struct{ n int }{n: 2}, 2},
		{"case4", struct{ n int }{n: 4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.args.n); got != tt.want {
				t.Errorf("nthUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
