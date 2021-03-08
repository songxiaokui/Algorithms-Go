package homework

import (
	. "Algorithms-Go/pkg"
	"reflect"
	"testing"
)

/*
@Time    : 2021/3/8 19:23
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 4.Merge two linkList_test.go.go
@Software: GoLand
*/

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"l1 && l2", struct {
			l1 *ListNode
			l2 *ListNode
		}{l1: ArrayConvertLinkList([]int{1, 3, 5}),
			l2: ArrayConvertLinkList([]int{2, 4, 6}),
		}, []int{1, 2, 3, 4, 5, 6}},

		{"l2", struct {
			l1 *ListNode
			l2 *ListNode
		}{l1: ArrayConvertLinkList([]int{}),
			l2: ArrayConvertLinkList([]int{2, 4, 6}),
		}, []int{2, 4, 6}},

		{"l1", struct {
			l1 *ListNode
			l2 *ListNode
		}{l1: ArrayConvertLinkList([]int{1, 3, 5}),
			l2: ArrayConvertLinkList([]int{}),
		}, []int{1, 3, 5}},

		{"l1 l2 == nil", struct {
			l1 *ListNode
			l2 *ListNode
		}{l1: ArrayConvertLinkList([]int{}),
			l2: ArrayConvertLinkList([]int{}),
		}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinkListConvertArray(mergeTwoLists(tt.args.l1, tt.args.l2)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
