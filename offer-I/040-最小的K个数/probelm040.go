/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/29 4:51 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _40_最小的K个数


func GetLeastNumbers_Solution( input []int ,  k int ) []int {
	// write code here

	if input == nil || len(input) == 0 || k == 0 {
		return nil
	}
	c := len(input) - 1

	for i := (len(input) - 1) / 2; i >= 0; i-- {
		minHeapAdjust(input, i, c)
	}
	var result []int
	for i := c; i > c - k; i-- {
		result = append(result, input[0])
		input[0], input[i] = input[i], input[0]
		minHeapAdjust(input, 0, i-1)
	}
	return result
}

func minHeapAdjust(arr [] int, i, c int) {
	var (
		tmp = arr[i]
		j = 0
	)

	for j = 2 * i; j <= c; j *= 2 {
		if j < c && arr[j] > arr[j+1] {
			j += 1
		}
		if tmp <= arr[j] {
			break
		}
		arr[i] = arr[j]
		i = j
	}
	arr[i] = tmp
}