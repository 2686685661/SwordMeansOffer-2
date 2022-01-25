/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/25 11:59 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _69_山峰数组的顶部

import "sort"

// 顺序查找
func peakIndexInMountainArray(arr []int) int {
	for i := 1; ; i++ {
		if arr[i] > arr[i+1] {
			return i
		}
	}
}


// 二分查找
func peakIndexInMountainArray2(arr []int) int {
	if arr == nil || len(arr) < 3 {
		return -1
	}

	low, mid, high := 1, 0, len(arr) - 2
	var ans int
	for low <= high {
		mid = (low + high) >> 1
		if arr[mid] > arr[mid + 1] {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}

// 使用库函数 sort.Search() 使用二分查找
func peakIndexInMountainArray3(arr []int) int {
	return sort.Search(len(arr) - 1, func(i int) bool {
		return arr[i] > arr[i + 1]
	})
}

