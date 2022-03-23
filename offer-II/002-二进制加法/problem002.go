/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/22 12:00 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _02_二进制加法

import "strconv"

func addBinary(a string, b string) string {

	ans := ""
	carry := 0
	la, lb := len(a), len(b)
	n := max(la, lb)

	for i := 0; i < n; i++ {
		if i < la {
			carry += int(a[la - i - 1] - '0')
		}
		if i < lb {
			carry += int(b[lb - i - 1] - '0')
		}

		ans = strconv.Itoa(carry % 2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}