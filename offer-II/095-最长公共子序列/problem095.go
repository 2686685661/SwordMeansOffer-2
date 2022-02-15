/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/11 11:42 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _95_最长公共子序列

/**
动态规划
https://leetcode-cn.com/problems/qJnOS7/solution/jian-zhi-offer-2-mian-shi-ti-95-shu-zhon-43k5/
 */


func longestCommonSubsequence(text1 string, text2 string) int {
	n1, n2 := len(text1), len(text2)
	if n1 < 1 || n2 < 1 {
		return 0
	}


	dp := make([][]int, n1 + 1)
	for i, _ := range dp {
		dp[i] = make([]int, n2 + 1)
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 0;i < n1; i++ {
		for j := 0; j < n2; j++ {
			if text1[i] == text2[j] {
				dp[i + 1][j + 1] =  dp[i][j] + 1
			} else {
				dp[i + 1][j + 1] = max(dp[i + 1][j], dp[i][j + 1])
			}
		}
	}
	return dp[n1][n2]

}