package homework

/*
@Time    : 2021/3/12 16:57
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 14.nthUglyNumber.go
@Software: GoLand
*/

/*
丑数判断:
循环除2
再循环除3
再循环除5
每个阶段除非最后一个因数为1，则不是丑数

剑指 Offer 49. 丑数
我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。



示例:

输入: n = 10
输出: 12
解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
说明:

1 是丑数。
n 不超过1690。
*/

func nthUglyNumber(n int) int {
	var dp = make([]int, n, n)
	dp[0] = 1

	var a, b, c int
	for i := 1; i < n; i++ {
		ra, rb, rc := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = Min(Min(ra, rb), rc)
		if dp[i] == ra {
			a++
		}
		if dp[i] == rb {
			b++
		}
		if dp[i] == rc {
			c++
		}
	}
	return dp[n-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
