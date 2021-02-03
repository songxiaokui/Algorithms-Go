package Testing

import "testing"

/*
@Time    : 2021/2/3 10:50
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 3.爬楼梯_test.go
@Software: GoLand
*/

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		target int
		value  int
	}{
		{2, 2},
		{3, 3},
	}
	for _, v := range tests {
		r := climbStairs(v.target)
		if r != v.value {
			t.Errorf("n=%d, 值错误!", v.target)
		}
	}
}
