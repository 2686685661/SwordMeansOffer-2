/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/6 3:59 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 线性表查找

// 在数组[0]增加哨兵， 防止数组越界
func OrderSearch(a []int, key int) int {
	if key == a[0] {
		return 0
	}
	a[0] = key
	i := len(a) - 1
	for a[i] != key {
		i--
	}
	return i
}
