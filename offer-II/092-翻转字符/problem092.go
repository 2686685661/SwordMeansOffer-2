/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/9 2:33 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _92_翻转字符

/**
动态规划
https://leetcode-cn.com/problems/cyJERH/solution/jian-zhi-offer-2-mian-shi-ti-92-shu-zhon-4oz8/
 */

func minFlipsMonoIncr(s string) int {
	if s == "" || len(s) < 1 {
		return 0
	}

	dp := make([][]int, len(s))
	for i, _ := range dp {
		dp[i] = make([]int, 2)
	}

	if s[0] == '0' {
		dp[0][1] = 1
	} else {
		dp[0][0] = 1
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i := 1; i < len(s); i++ {

		dp[i][0] = dp[i - 1][0]
		if s[i] != '0' {
			dp[i][0] += 1
		}

		dp[i][1] = min(dp[i - 1][0], dp[i - 1][1])
		if s[i] != '1' {
			dp[i][1] += 1
		}
	}

	n := len(s) - 1
	return min(dp[n][0], dp[n][1])

}