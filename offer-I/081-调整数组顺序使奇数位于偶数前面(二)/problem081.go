/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/30 7:44 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _81_调整数组顺序使奇数位于偶数前面_二_

func reOrderArrayTwo( array []int ) []int {

	if array == nil || len(array) < 2 {
		return array
	}

	low, high := 0, len(array) - 1

	for low < high {
		for low < high && array[high] % 2 == 0 {
			high--
		}
		for low < high && array[low] % 2 == 1 {
			low++
		}
		array[low], array[high] = array[high], array[low]
		low++
		high--
	}
	return array
}