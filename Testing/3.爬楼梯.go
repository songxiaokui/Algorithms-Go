package Testing

/*
@Time    : 2021/2/3 10:46
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 3.爬楼梯.go
@Software: GoLand
*/

import (
	"math"
)

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
// 注意：给定 n 是一个正整数。
//
// 示例 1：
//
// 输入： 2
//输出： 2
//解释： 有两种方法可以爬到楼顶。
//1.  1 阶 + 1 阶
//2.  2 阶
//
// 示例 2：
//
// 输入： 3
//输出： 3
//解释： 有三种方法可以爬到楼顶。
//1.  1 阶 + 1 阶 + 1 阶
//2.  1 阶 + 2 阶
//3.  2 阶 + 1 阶
//
// Related Topics 动态规划
// 👍 1450 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 方案一 斐波那契数列在原地交换
// 内存消耗:1.9 MB,击败了77.65% 的Go用户
func climbStairs1(n int) int {
	if n == 0 {
		return 0
	}
	var prev, next, result = 0, 1, 1
	for i := 1; i < n; i++ {
		prev, next = next, result
		result = prev + next
	}
	return result
}

// 方案二 斐波那契通向公式， 第n项，有 n+1，从0开始
// 内存消耗:1.8 MB,击败了99.75% 的Go用户
func climbStairs2(n int) int {
	sqrt5 := math.Sqrt(5)
	pow1 := math.Pow((1+sqrt5)/2, float64(n+1))
	pow2 := math.Pow((1-sqrt5)/2, float64(n+1))
	result := math.Round((pow1 - pow2) / sqrt5)
	return int(result)
}

// 方案三 动态规划，自底向上
// 内存消耗:1.9 MB,击败了44.02% 的Go用户 可能存在内存拷贝问题，追加O（1）
func climbStairs3(n int) int {
	dp := []int{1, 1}
	for i := 2; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}

// 动态规划，但是使用map进行保存
// 内存消耗:2 MB,击败了6.28% 的Go用户
func climbStairs(n int) int {
	hashMap := make(map[int]int)
	hashMap[0] = 1
	hashMap[1] = 1

	for i := 2; i <= n; i++ {
		hashMap[i] = hashMap[i-1] + hashMap[i-2]
	}
	return hashMap[n]
}

//leetcode submit region end(Prohibit modification and deletion)
