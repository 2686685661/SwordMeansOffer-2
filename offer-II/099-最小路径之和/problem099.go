/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/15 11:09 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _99_最小路径之和

/**
动态规划
 */

func minPathSum(grid [][]int) int {

	if len(grid) < 1 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	if m > 200 || n < 1 || n > 200 {
		return 0
	}

	dp := make([][]int, m)
	for i,_ := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i - 1][0] + grid[i][0]
	}

	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j - 1] + grid[0][j]
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j]
		}
	}

	return dp[m - 1][n - 1]
}