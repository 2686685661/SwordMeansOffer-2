/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/16 4:07 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _57_和为S的两个数字


func FindNumbersWithSum( array []int ,  sum int ) []int {
	if array == nil || len(array) < 2 {
		return nil
	}

	start, end := 0, len(array)  -1

	for start < end {
		if array[start] + array[end] < sum {
			start+=1
		} else if array[start] + array[end] > sum {
			end -= 1
		} else {
			break
		}
	}

	if start >= end {
		return nil
	}
	return []int{array[start], array[end]}
}