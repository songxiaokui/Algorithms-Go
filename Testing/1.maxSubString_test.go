package Testing

import (
	"testing"
)

/*
@Time    : 2021/1/29 00:47
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 1.maxSubString_test.go
@Software: GoLand
*/

func TestLengthOfLongestSubStringBitMap(t *testing.T) {
	var testCase = []struct {
		s      string
		target string
	}{
		{"abcabcbbytrsdf", "bytrsdf"},
		{"pwwasd", "wasd"},
		{"asdfggggg", "asdfg"},
		{"ggggg", "g"},
	}
	for _, v := range testCase {
		_, _, _, st := lengthOfLongestSubStringBitMap(v.s)
		if st != v.target {
			t.Errorf("ErrInfo: %s的值应该是:%s,不是：%s", v.s, v.target, st)
		}
	}
}
