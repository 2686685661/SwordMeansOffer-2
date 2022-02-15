/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/15 11:45 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _00_三角形中最小路径之和


/**
动态规划
https://leetcode-cn.com/problems/IlPe0q/solution/jian-zhi-offer-2-mian-shi-ti-100-shu-zho-ejjf/
 */

func minimumTotal(triangle [][]int) int {
	tl := len(triangle)
	if tl < 1 || tl > 200 {
		return 0
	}

	dp := make([][]int, tl)
	for i, _ := range dp {
		dp[i] = make([]int, tl)
	}

	dp[0][0] = triangle[0][0]

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i := 1; i < tl; i++ {
		dp[i][0] = dp[i - 1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i - 1][j], dp[i - 1][j - 1]) + triangle[i][j]
		}
		dp[i][i] = dp[i - 1][i - 1] + triangle[i][i]
	}

	res := 1<<31 - 1

	for _, v := range dp[tl - 1] {
		res = min(res, v)
	}
	return res

}