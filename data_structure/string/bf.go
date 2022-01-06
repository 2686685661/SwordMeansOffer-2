/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/7/21 9:28 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package string

/**
https://blog.csdn.net/ns_code/article/details/19286279
匹配过程动画：https://www.bilibili.com/video/BV115411K7E8?from=search&seid=5550110250468223594
*/

func Bf(str, tag string, pos int) int {
	i, j := pos, 0
	for i < len(str) && j < len(tag) {
		if str[i] == tag[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j >= len(tag) {
		return i - len(tag)
	}
	return -1
}
