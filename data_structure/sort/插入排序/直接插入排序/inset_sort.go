/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 12:16 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 直接插入排序


// 直接插入排序
func InsetSort(elem []int) []int {
	if len(elem) < 2 {
		return elem
	}

	var j int
	for i := 2; i < len(elem); i++ {
		if elem[i] < elem[i - 1] {
			elem[0] = elem[i]
			for j = i - 1; elem[j] > elem[0]; j-- {
				elem[j + 1] = elem[j] // 设置哨兵
			}
			elem[j + 1] = elem[0]
		}
	}

	return elem
}