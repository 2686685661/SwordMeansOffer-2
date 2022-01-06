/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/20 12:08 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _20_表示数值的字符串

import (
	"strings"
)

func IsNumeric( str string ) bool {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return false
	}

	return isScientific(str) || isInteger(str,0, len(str) - 1) || isFloat(str,0, len(str) - 1)
}

// 字符串是否是科学计数法
func isScientific(str string) bool {
	s, m := 0, false
	for i, v := range str {
		if v == 'e' || v == 'E' {
			s, m = i, true
			break
		}
	}
	if m {
		return (isInteger(str, 0, s - 1) || isFloat(str, 0, s - 1)) && isInteger(str, s + 1, len(str) - 1)
	}
	return false

}

// 字符串是否是整数
func isInteger(str string, l, r int) bool {
	if l > r {
		return false
	}
	if str[l] == '+' || str[l] == '-' {
		l += 1
		if l > r {
			return false
		}
	}
	return containsOneNum(str, l, r)
}

// 字符串是否是小数
func isFloat(str string, l, r int) bool {
	if l > r {
		return false
	}
	if str[l] == '+' || str[l] == '-' {
		l++
		if l > r {
			return false
		}
	}

	s, m := l, false

	for i := l; i <= r; i++ {
		if str[i] == '.' {

			s, m = i, true
			break
		}
	}
	if m {
		if s == r {
			return containsOneNum(str, l, s - 1)
		} else if s == l {
			return containsOneNum(str, l + 1, r)
		} else {
			return containsOneNum(str, l, s - 1) && containsOneNum(str, s + 1, r)
		}
	}
	return false
}

// 字符串是否包含至少一位数字
func containsOneNum(str string, l, r int) bool {
	if l > r {
		return false
	}
	for l <= r {
		if str[l] < '0' || str[l] > '9' {
			return false
		}
		l++
	}
	return true
}