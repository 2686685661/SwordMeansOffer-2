/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/10 2:46 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _94_最少回文分割


/**
动态规划
https://leetcode-cn.com/problems/omKAoA/solution/jian-zhi-offer-2-mian-shi-ti-94-shu-zhon-vbe6/
 */

func minCut(s string) int {

	if len(s) < 1 {
		return 0
	}

	n := len(s)
	isPalindrome := make([]map[int]bool, n)
	for i, _ := range isPalindrome {
		isPalindrome[i] = make(map[int]bool)
	}

	cenL, cenR := 0, 0
	for cenL < n {
		left, right := cenL, cenR
		for left >= 0 && right < n && s[left] == s[right] {
			isPalindrome[left][right] = true
			left--
			right++
		}
		if cenR > cenL {
			cenL++
		} else {
			cenR++
		}
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if isPalindrome[0][i] {
			dp[i] = 0
		} else {
			dp[i] = i
			for j := 1; j <= i; j++ {
				if isPalindrome[j][i] {
					dp[i] = min(dp[i], dp[j - 1] + 1)
				}
			}
		}
	}
	return dp[n - 1]
}