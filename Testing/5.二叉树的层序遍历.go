package Testing

import "fmt"

/*
@Time    : 2021/2/8 22:42
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 5.二叉树的层序遍历.go
@Software: GoLand
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	// 按层输出
	var result [][]int
	if root == nil {
		return result
	}
	stackList := []*TreeNode{root}

	for i := 0; len(stackList) > 0; i++ {
		result = append(result, []int{})
		// 用来存储该节点的子节点
		var secondList []*TreeNode
		for j := 0; j < len(stackList); j++ {
			currentNode := stackList[j]
			result[i] = append(result[i], currentNode.Val)
			// 处理左右节点
			if currentNode.Left != nil {
				secondList = append(secondList, currentNode.Left)
			}
			if currentNode.Right != nil {
				secondList = append(secondList, currentNode.Right)
			}
		}
		// 不满足条件时，进行数据替换，进行下一轮操作
		stackList = secondList
	}
	fmt.Println(result)
	return result
}

func LevelOrder(root *TreeNode) []int {
	// DFS广度遍历
	var result []int
	if root == nil {
		return result
	}
	var stackList []*TreeNode
	stackList = append(stackList, root)
	for len(stackList) != 0 {
		// 弹出第一个节点
		currentNode := stackList[0]
		stackList = stackList[1:len(stackList)]
		// 修改节点stack
		if currentNode != nil {
			result = append(result, currentNode.Val)
			if currentNode.Left != nil {
				stackList = append(stackList, currentNode.Left)
			}
			if currentNode.Right != nil {
				stackList = append(stackList, currentNode.Right)
			}
		}
	}
	fmt.Println(result)
	return result
}
