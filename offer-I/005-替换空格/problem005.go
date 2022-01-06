/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/7/13 5:23 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _05_替换空格

func replaceSpace_1( s string ) string {
	// write code here

	count := 0
	// 遍历一遍字符串， 统计字符出现的数目, 计算替换后的字符串长度
	for _, v := range s {
		if v == ' ' {
			count++
		}
	}
	newLen := len(s) + count * 2
	ss := make([]rune, newLen)

	// 两个index，一个指向length-1, 另一个指向newlength-1，遍历一遍字符串，完成替换
	for i, ni := len(s) - 1, newLen -1; i >= 0 && ni >= 0; {
		if s[i] == ' ' {
			ss[ni] = '0'
			ni--
			ss[ni] = '2'
			ni--
			ss[ni] = '%'
			ni--
			i--
		} else {
			ss[ni] = rune(s[i])
			i--
			ni--
		}
	}
	return string(ss)

}