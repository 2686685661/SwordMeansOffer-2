/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/9 11:09 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _91_粉刷房子

/**
动态规划
https://leetcode-cn.com/problems/JEj789/solution/tong-su-yi-dong-de-dpzuo-fa-su-kan-by-sm-2mt3/
 */

func minCost(costs [][]int) int {
	if costs == nil || len(costs) < 1 {
		return 0
	}
	dp := make([][]int, len(costs))
	// dp[i][0] 第i间房子涂红色时，前i-1间房子累计所需的最小成本
	// dp[i][1] 第i间房子涂蓝色时，前i-1间房子累计所需的最小成本
	// dp[i][2] 第i间房子涂绿色时，前i-1间房子累计所需的最小成本
	for i, _ := range dp {
		dp[i] = make([]int, 3)
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i, _ := range costs[0] {
		dp[0][i] = costs[0][i]
	}

	for i := 1; i < len(costs); i++ {
		dp[i][0] = min(dp[i - 1][1], dp[i - 1][2]) + costs[i][0]
		dp[i][1] = min(dp[i - 1][0], dp[i - 1][2]) + costs[i][1]
		dp[i][2] = min(dp[i - 1][0], dp[i - 1][1]) + costs[i][2]
	}
	n := len(costs) - 1
	return min(dp[n][2], min(dp[n][0], dp[n][1]))
}