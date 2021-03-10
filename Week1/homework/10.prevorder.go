package homework

/*
@Time    : 2021/3/10 21:03
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 10.prevorder.go
@Software: GoLand
*/

/*
589. N 叉树的前序遍历
给定一个 N 叉树，返回其节点值的 前序遍历 。

N 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

进阶：

递归法很简单，你可以使用迭代法完成此题吗?

示例 1：

输入：root = [1,null,3,2,4,null,5,6]
输出：[1,3,5,6,2,4]
示例 2：

输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
输出：[1,2,3,6,7,11,14,4,8,12,5,9,13,10]
*/

func PrevOrder(root *Nodes) []int {
	var result []int
	prevRecursion(root, &result)
	return result
}

func prevRecursion(node *Nodes, result *[]int) {
	if node == nil {
		return
	}
	// prev add value in array
	*result = append(*result, node.Val)

	for _, r := range node.Children {
		prevRecursion(r, result)
	}
}
