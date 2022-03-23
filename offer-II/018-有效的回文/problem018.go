/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/27 9:54 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _18_有效的回文

import (
	"fmt"
	"strings"
)


// 双指针
func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	s = strings.ToLower(s)
	fmt.Println(s)
	left, right := 0, len(s) - 1

	h := func(b byte) bool {
		if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
			return true
		}
		return false
	}

	for left < right {
		for left < right && !h(s[left]) {
			left++
		}
		for left < right && !h(s[right]) {
			right--
		}
		if left < right && s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
