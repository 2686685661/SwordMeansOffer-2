/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/30 5:52 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _34_外星语言是否排序


// 字典序概念
// 将words中所有字符替换为正常字典序的字符
func isAlienSorted(words []string, order string) bool {
	dict := make(map[rune]rune)
	for i, v := range order {
		dict[v] = rune(97 + i)
	}

	h := func(s string) string {
		rs := []rune(s)
		for i, v := range rs {
			rs[i] = dict[v]
		}
		return string(rs)
	}
	for i, w := range words {
		words[i] = h(w)
		if i > 0 && words[i] < words[i - 1] {
			return false
		}
	}
	return true
}