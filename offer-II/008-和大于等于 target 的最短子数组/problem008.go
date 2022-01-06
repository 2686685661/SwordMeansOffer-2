/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/23 1:26 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _08_和大于等于_target_的最短子数组

import "math"

func minSubArrayLen(target int, nums []int) int {
	if nums == nil || len(nums) < 1 || target < 1 {
		return 0
	}

	min := math.MaxInt32

	sum := 0
	for i, j := 0, 0; j < len(nums); j++ {
		sum += nums[j]

		for i <= j && sum >= target {
			min = Min(min, j - i + 1)
			sum -= nums[i]
			i++
		}
	}
	if min == math.MaxInt32 {
		min = 0
	}
	return min

}


func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}