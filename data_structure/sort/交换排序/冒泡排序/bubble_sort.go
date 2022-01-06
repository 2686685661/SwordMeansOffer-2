/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 2:29 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 冒泡排序

func BubbleSort(elem []int) []int {
	for i := 1; i < len(elem); i++ {
		for j := len(elem)-1; j >= i; j-- {
			if elem[j - 1] > elem[j] {
				elem[j - 1], elem[j] = elem[j], elem[j - 1]
			}
		}
	}
	return elem
}

func BubbleSort2(elem []int) []int {

	var flag = true

	for i := 1; i < len(elem) && flag; i++ {
		flag = false
		for j := len(elem)-1; j >= i; j-- {
			if elem[j - 1] > elem[j] {
				elem[j - 1], elem[j] = elem[j], elem[j - 1]
				flag = true
			}
		}
	}

	return elem
}