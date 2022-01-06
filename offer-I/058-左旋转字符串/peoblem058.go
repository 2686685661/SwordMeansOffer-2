/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/16 4:53 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _58_左旋转字符串

func LeftRotateString( str string ,  n int ) string {
	if str == "" || n <= 0 {
		return str
	}
	char := []byte(str)

	if n > len(char) {
		n = n % len(char)
	}

	char = append(char[n:], char[:n]...)

	return string(char)
}