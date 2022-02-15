/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/8 5:17 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _88_爬楼梯的最少成本

/**
动态规划
https://donleetcode-cn.com/problems/GzCJIP/solution/jian-zhi-offer-2-mian-shi-ti-88-shu-zhon-2dno/
 */

func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return 0
	}
	dp := make([]int, len(cost) + 1)

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2])
	}
	return dp[len(cost)]

}