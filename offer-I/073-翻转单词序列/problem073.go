/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/18 8:38 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _73_翻转单词序列

func ReverseSentence( str string ) string {

	if str == "" {
		return str
	}


	str = Reverse(str, 0, len(str) - 1)

	front := -1
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			next := i





			str = Reverse(str, front + 1, next - 1)
			front = next
		}
	}
	str = Reverse(str, front + 1, len(str) - 1)
	return str
}


func Reverse(str string, start, end int) string {
	if len(str) < end {
		return str
	}
	ch := []byte(str)


	for start < end {
		ch[start], ch[end] = str[end], ch[start]
		start++
		end--
	}
	return string(ch)
}