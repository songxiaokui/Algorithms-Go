package pkg

/*
@Time    : 2021/3/8 19:06
@Author  : austsxk
@Email   : austsxk@163.com
@File    : list.go
@Software: GoLand
*/

// array convert to linkList
type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayConvertLinkList(array []int) *ListNode {
	n := &ListNode{}
	result := n
	if len(array) == 0 {
		return nil
	}
	for i := 0; i < len(array); i++ {
		node := &ListNode{}
		node.Val = array[i]
		n.Next = node
		n = n.Next
	}
	return result.Next
}

// linkList convert to array
func LinkListConvertArray(n *ListNode) []int {
	if n == nil {
		return nil
	}
	var result []int
	for n != nil {
		d := n.Val
		result = append(result, d)
		n = n.Next
	}
	return result
}
