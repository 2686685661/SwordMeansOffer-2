/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/20 11:29 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _19_正则表达式匹配


func Match( str string ,  pattern string ) bool {
	return matchCore(str, pattern, 0, 0)
}


func matchCore(str, pattern string, i, j int) bool {

	if i == len(str) && j == len(pattern) {
		return true
	}
	if i != len(str) && j == len(pattern) {
		return false
	}

	// 模式下一个字符是'*'
	if j + 1 < len(pattern) && pattern[j + 1] == '*' {
		// 字符串当前位与模式当前位匹配
		if i != len(str) && (str[i] == pattern[j] || pattern[j] == '.') {
			// 模式后移2，视为x*匹配0个字符 || 视为模式匹配1个字符 || 匹配1个，再匹配str中的下一个
			return matchCore(str, pattern, i, j + 2) || matchCore(str, pattern, i + 1, j + 2) || matchCore(str, pattern, i + 1, j)

		} else { // 字符串当前位与模式当前位不匹配
			return matchCore(str, pattern, i, j + 2)
		}
	}

	// 当前字符和模式匹配，且模式下一个字符不是'*'
	if i != len(str) && (str[i] == pattern[j] || pattern[j] == '.') {
		return matchCore(str, pattern, i + 1, j + 1)
	}
	// 不匹配，直接返回false
	return false
}