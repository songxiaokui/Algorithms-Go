package Testing

import "testing"

/*
@Time    : 2021/2/4 12:18
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 4.括号生成_test.go
@Software: GoLand
*/

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		k int
		v []string
	}{
		{1, []string{"()"}},
		{2, []string{"(())", "()()"}},
	}
	for _, v := range tests {
		r := generateParenthesis(v.k)
		for i, k := range v.v {
			if k != r[i] {
				t.Fail()
			}
		}
	}
}
