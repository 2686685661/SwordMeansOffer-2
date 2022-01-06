/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/30 11:02 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _75_字符流中第一个不重复的字符

var (
	m = make(map[byte]int, 128)
	q []byte
)
func Insert(ch byte) {
	if _, ok := m[ch]; ok {
		m[ch]++
	} else {
		m[ch] = 1
	}

	if m[ch] == 1 {
		q = append(q, ch)
	}
}

func FirstAppearingOnce() byte {
	for len(q) > 0 && m[q[0]] > 1 {
		q = q[1:]
	}
	if len(q) == 0 {
		return '#'
	}
	return q[0]
}