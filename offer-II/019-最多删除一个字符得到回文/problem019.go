/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/28 10:55 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _19_最多删除一个字符得到回文


func validPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	l, r := 0, len(s) - 1
	h := func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	for l < r {
		if s[l] != s[r] {
			return h(l+1, r) || h(l, r-1)
		}
		l++
		r--
	}
	return true
}