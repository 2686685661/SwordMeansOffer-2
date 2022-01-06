/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/23 11:48 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _07_数组中和为_0_的三个数

import "sort"

func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}

	var result [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums) - 2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		begin, end := i + 1, len(nums) - 1

		bc := func() {
			for begin < end && nums[begin] == nums[begin - 1] {
				begin++
			}
		}
		ec := func() {
			for begin < end && nums[end] == nums[end + 1] {
				end--
			}
		}
		for begin < end {
			sum := nums[i] + nums[begin] + nums[end]
			if sum < 0 {
				begin++
				bc()
			} else if sum > 0 {
				end--
				ec()
			} else {
				result = append(result, []int{nums[i], nums[begin], nums[end]})
				begin++
				end--
				bc()
				ec()
			}

		}
	}
	return result

}