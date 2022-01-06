/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/10 2:11 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 基数排序

func RadixSort(elem []int) []int {
	max, exp := 0, 0
	for _, v := range elem {
		if v > max {
			max = v
		}
	}

	for exp = 1; max / exp > 0; exp *= 10 {
		elem = countSort(elem, exp)
	}
	return elem
}

func countSort(elem []int, exp int) []int {
	output, buckets := make([]int, len(elem)), make([]int, 10)

	for _, v := range elem {
		buckets[(v / exp) % 10]++
	}
	for i := 1; i < 10; i++ {
		buckets[i] += buckets[i - 1]
	}

	for i := len(elem) -1; i >= 0; i-- {
		output[ buckets[(elem[i] / exp) % 10] - 1 ] = elem[i]
		buckets[(elem[i] / exp) % 10]--
	}
	return output
}