/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/15 2:03 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _53_数字在升序数组中出现的次数

func GetNumberOfK( data []int ,  k int ) int {
	if data == nil || len(data) == 0 {
		return 0
	}
	first := getFirstK(data, k, 0, len(data) - 1)
	last := getLastK(data, k, 0, len(data)  - 1)
	if first != -1 && last != -1 {
		return last - first + 1
	}
	return 0
}

func getFirstK(data []int, k, start, end int) int {

	if start > end {
		return -1
	}
	mid := (start + end) >> 1
	if data[mid] > k {
		end = mid - 1
	} else if data[mid] < k {
		start = mid + 1
	} else {
		if (mid > 0 && data[mid - 1] != k) || mid == 0 {
			return mid
		} else {
			end = mid - 1
		}
	}

	return getFirstK(data, k, start, end)
}

func getLastK(data []int ,k, start, end int) int {
	mid := (start + end) >> 1
	for start <= end {
		if data[mid] > k {
			end = mid - 1
		} else if data[mid] < k {
			start = mid + 1
		} else {
			if (mid < len(data) - 1 && data[mid + 1] != k) || mid == len(data) - 1 {
				return mid
			} else {
				start = mid + 1
			}
		}
		mid = (start + end) >> 1
	}
	return -1
}

