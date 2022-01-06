/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/10 1:21 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 计数排序

func CountSort(elem []int) []int {

	if elem == nil || len(elem) <1 {
		return elem
	}

	max := 0
	for _, v := range elem {
		if v > max {
			max = v
		}
	}

	counts := make([]int, max + 1)

	for _,v := range elem {
		counts[v]++
	}


	for i, j := 0, 0; i <= max; i++ {
		for {
			if counts[i] <= 0 {
				break
			}
			elem[j] = i
			j++
			counts[i]--
		}
	}

	return elem
}