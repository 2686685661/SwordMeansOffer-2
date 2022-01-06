/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 3:50 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 简单选择排序

func SelectSort(elem []int) []int {

	for i := 1; i < len(elem); i++ {
		k := i
		for j := i + 1; j < len(elem); j++ {
			if elem[k] > elem[j] {
				k = j
			}
		}
		if k != i {
			elem[i], elem[k] = elem[k], elem[i]
		}
	}
	return elem
}