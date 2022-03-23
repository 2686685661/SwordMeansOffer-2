/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/14 1:53 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _98_路径的数目

/**
动态规划
https://leetcode-cn.com/problems/2AoeFn/solution/jian-zhi-offer-2-mian-shi-ti-98-shu-zhon-vo2b/
 */

func uniquePaths(m int, n int) int {

	if m < 1 || n < 1 || m > 100 || n > 100 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}

	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j - 1] + dp[i - 1][j]
		}
	}
	return dp[m - 1][n - 1]
}