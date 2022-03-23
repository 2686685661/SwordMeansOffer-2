/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/11 12:07 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _96_字符串交织

/**
动态规划
https://leetcode-cn.com/problems/IY6buf/solution/jian-zhi-offer-2-mian-shi-ti-96-shu-zhon-5kc7/
 */

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1, n2, n3 := len(s1), len(s2), len(s3)
	if n1 > 100 || n2 > 100 || n3 > 200 {
		return false
	}
	if n1 + n2 != n3 {
		return false
	}

	dp := make([]map[int]bool, n1 + 1)
	for i, _ := range dp {
		dp[i] = make(map[int]bool, n2 + 1)
	}

	dp[0][0] = true
	for j := 0; j < n2 && s2[j] == s3[j]; j++ {
		dp[0][j + 1] = true
	}

	for i := 0; i < n1 && s1[i] == s3[i]; i++ {
		dp[i + 1][0] = true
	}

	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if s3[i + j + 1] == s1[i] {
				dp[i + 1][j + 1] = dp[i][j + 1]
			}
			if s3[i + j + 1] == s2[j] {
				dp[i + 1][j + 1] = dp[i + 1][j]
			}
		}
	}

	return dp[n1][n2]
}