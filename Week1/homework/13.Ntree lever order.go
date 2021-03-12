package homework

/*
@Time    : 2021/3/12 11:18
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 13.Ntree lever order.go
@Software: GoLand
*/

/*
429. N 叉树的层序遍历
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。
示例 1：
输入：root = [1,null,3,2,4,null,5,6]
输出：[[1],[3,2,4],[5,6]]
示例 2：

输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
输出：[[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]

提示：

树的高度不会超过 1000
树的节点总数在 [0, 10^4] 之间
*/

func levelOrder(root *Nodes) [][]int {
	// expect handler
	var result [][]int
	if root == nil {
		return result
	}
	// use a stackList save current level node, put root as first level
	stackList := []*Nodes{root}
	for i := 0; len(stackList) > 0; i++ {
		result = append(result, []int{})
		// save next level nodes
		var secondList []*Nodes
		for j := 0; j < len(stackList); j++ {
			// get current node
			n := stackList[j]
			result[i] = append(result[i], n.Val)
			if n.Children != nil {
				for _, newNode := range n.Children {
					if newNode.Children != nil {
						secondList = append(secondList, newNode)
					}
				}
			}
		}
		stackList = secondList
	}
	return result
}
