/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/20 1:49 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _21_调整数组顺序使奇数位于偶数前面_一_

func ReOrderArray( array []int ) []int {

	for i,j := 0, len(array) - 1; i <= j; i++ {
		if array[i] % 2 == 1 {
			continue
		}

		array = append(array, array[i])
		array = append(array[:i], array[i+1:]...)
		i--
		j--
	}
	return array
}