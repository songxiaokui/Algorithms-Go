package Testing

import (
	"fmt"
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
	var x = "abcabcbbytrsdf"
	start, end, max, st := lengthOfLongestSubStringBitMap(x)
	if st != "bytrsdf" {
		t.Fail()
	}
	fmt.Println(start, end, max)
}
