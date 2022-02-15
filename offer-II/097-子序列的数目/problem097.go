/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/14 1:31 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _97_子序列的数目

/**
动态规划
https://leetcode-cn.com/problems/21dk04/solution/jian-zhi-offer-2-mian-shi-ti-97-shu-zhon-mnfg/
 */

func numDistinct(s string, t string) int {
	sn, tn := len(s), len(t)
	if sn < 0 || tn < 0 {
		return 0
	}

	dp := make([][]int, sn + 1)
	for i, _ := range dp {
		dp[i] = make([]int, tn + 1)
	}
	dp[0][0] = 1
	for i := 0; i < sn; i++ {
		dp[i + 1][0] = 1
		for j := 0; j < tn; j++ {
			if s[i] == t[j] {
				dp[i + 1][j + 1] = dp[i][j + 1] + dp[i][j]
			} else {
				dp[i + 1][j + 1] = dp[i][j + 1]
			}
		}
	}
	return dp[sn][tn]

}