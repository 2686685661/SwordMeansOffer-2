/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/27 9:05 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _15_字符串中的所有变位词

// 滑动窗口
func findAnagrams(s string, p string) []int {
	if s == "" || p == "" || len(p) > len(s) {
		return nil
	}
	var cnt, con [26]int
	var res []int
	for i, v := range p {
		cnt[v - 'a']++
		cnt[s[i] - 'a']--
	}
	if cnt == con {
		res = append(res, 0)
	}

	for i := len(p); i < len(s); i++ {
		cnt[s[i] - 'a']--
		cnt[s[i - len(p)] - 'a']++
		if cnt == con {
			res = append(res, i - len(p) + 1)
		}
	}
	return res
}