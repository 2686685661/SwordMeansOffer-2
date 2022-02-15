/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/8 4:19 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _86_分割回文子字符串


/**
回溯
https://leetcode-cn.com/problems/M99OJA/solution/jian-zhi-offer-2-mian-shi-ti-86-shu-zhon-gstd/
 */

func partition(s string) [][]string {

	if len(s) < 1 || len(s) > 16 {
		return nil
	}

	var res [][]string
	var cur []string

	var isPalindrome func(s string) bool
	isPalindrome = func(s string) bool {
		n := len(s)
		l, r := 0, n-1
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(s) {
			res = append(res, append([]string{}, cur...))
			return
		}
		for i := idx + 1; i <= len(s); i++ {
			if isPalindrome(s[idx:i]) {
				cur = append(cur, s[idx:i])
				dfs(i)
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs(0)
	return res
}