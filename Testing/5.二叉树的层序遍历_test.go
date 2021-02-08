package Testing

import (
	"testing"
)

/*
@Time    : 2021/2/8 23:06
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 5.二叉树的层序遍历_test.go
@Software: GoLand
*/
func NewTreeList() *TreeNode {
	a := TreeNode{}
	a.Val = 1

	b := TreeNode{}
	b.Val = 2

	c := TreeNode{}
	c.Val = 3

	a.Left = &b
	a.Right = &c

	e := TreeNode{}
	e.Val = 4

	f := TreeNode{}
	f.Val = 5

	g := TreeNode{}
	g.Val = 6

	h := TreeNode{}
	h.Val = 7

	b.Left = &e
	b.Right = &f
	c.Right = &g
	c.Left = &h

	return &a
}

func TestLevelOrder(t *testing.T) {
	a := NewTreeList()
	levelOrder(a)

}

func TestLevelOrder2(t *testing.T) {
	a := NewTreeList()
	LevelOrder(a)
}
