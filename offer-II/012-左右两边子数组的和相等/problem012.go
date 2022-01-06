/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/24 11:45 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _12_左右两边子数组的和相等

// 前缀和思想
func pivotIndex(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return -1
	}

	total, sum := 0, 0
	for _, n := range nums {
		total += n
	}

	for i := 0; i < len(nums); i++ {
		if total - nums[i] == 2 * sum {
			return i
		}
		sum += nums[i]
	}
	return -1

}