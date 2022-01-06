/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/5 2:47 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _50_第一个只出现一次的字符


func FirstNotRepeatingChar( str string ) int {

	m := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		m[str[i]] += 1
	}

	for i := 0; i < len(str); i++ {
		if m[str[i]] == 1 {
			return i
		}
	}
	return -1

}