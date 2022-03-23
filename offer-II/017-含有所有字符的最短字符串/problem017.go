/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/27 9:36 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _17_含有所有字符的最短字符串


//滑动窗口
//https://leetcode-cn.com/problems/M1oyTv/solution/jian-zhi-offer-ii-017-han-you-suo-you-zi-ttsd/
func minWindow(s string, t string) string {

	if s == "" || t == "" || len(t) > len(s) {
		return ""
	}

	var m [60]int
	min := len(s) + 1
	res := ""

	for i, v := range t {
		m[v - 'A']++
		m[s[i] - 'A']--
	}
	if allZero(m) {
		return s[:len(t)]
	}

	for l, r := 0, len(t); r < len(s); r++ {
		m[s[r] - 'A']--
		for allZero(m) {
			if r - l + 1 < min {
				min = r - l + 1
				res = s[l:r+1]
			}
			m[s[l] - 'A']++
			l++
		}
	}
	return res
}

func allZero(m [60]int) bool {
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}