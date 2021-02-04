package Testing

/*
@Time    : 2021/2/4 12:14
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 4.括号生成.go
@Software: GoLand
*/

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：["()"]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
// Related Topics 字符串 回溯算法
// 👍 1542 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 回溯法 + 递归

func generateParenthesis(n int) []string {
	var result []string
	//recursion(0, 0, n, &result, "")
	recursion2(n, n, &result, "")
	return result
}

func recursion(l int, r int, n int, result *[]string, s string) {
	if l == n && r == n {
		*result = append(*result, s)
	}
	if l < n {
		recursion(l+1, r, n, result, s+"(")
	}
	if l > r {
		recursion(l, r+1, n, result, s+")")
	}
}

func recursion2(l, r int, result *[]string, s string) {
	if l == 0 && r == 0 {
		*result = append(*result, s)
	}
	if l > r || r < 0 || l < 0 {
		return
	}
	recursion2(l-1, r, result, s+"(")
	recursion2(l, r-1, result, s+")")
}
