package pkg

/*
@Time    : 2021/3/8 11:54
@Author  : austsxk
@Email   : austsxk@163.com
@File    : utility.go
@Software: GoLand
*/

// compare two slice is equal
func TwoSliceEqual(l1 []int, l2 []int) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil || l2 == nil {
		return false
	}
	if len(l1) != len(l2) {
		return false
	}
	for e := range l1 {
		if l1[e] != l2[e] {
			return false
		}
	}
	return true
}
