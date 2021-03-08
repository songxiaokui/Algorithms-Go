package homework

import . "Algorithms-Go/pkg"

/*
@Time    : 2021/3/8 18:44
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 4.Merge two linkList.go
@Software: GoLand
*/

/*
problems description:
21. 合并两个有序链表
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：

输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：

输入：l1 = [], l2 = []
输出：[]
示例 3：

输入：l1 = [], l2 = [0]
输出：[0]
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// null value process
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		return l2
	}
	if l1 != nil && l2 == nil {
		return l1
	}
	// not empty process
	// defined new linkList
	newList := &ListNode{}
	// save as result
	result := newList
	// loop l1 and l2, until one of l1 or l2 is nil
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			newList.Next = l2
			l2 = l2.Next
		} else {
			newList.Next = l1
			l1 = l1.Next
		}
		newList = newList.Next
	}
	// process null_empty link list
	if l1 != nil {
		newList.Next = l1
	}
	if l2 != nil {
		newList.Next = l2
	}
	return result.Next
}
