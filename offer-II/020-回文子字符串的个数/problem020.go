/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/28 12:17 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _20_回文子字符串的个数


// 中心扩展
// https://leetcode-cn.com/problems/a7VOhD/solution/hui-wen-zi-zi-fu-chuan-de-ge-shu-by-leet-ejfv/
func countSubstrings(s string) int {
	n := len(s)
	count := 0
	for i := 0; i< 2 * n - 1; i++ {
		l, r := i / 2, i / 2 + i % 2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			count += 1
		}
	}
	return count

}