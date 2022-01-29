/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/28 11:19 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _79_所有子集

/**
回溯法
 */

/**
给定一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 */

/**
解题思路：https://leetcode-cn.com/problems/subsets/solution/shou-hua-tu-jie-zi-ji-hui-su-fa-xiang-jie-wei-yun-/
 */


func subsets(nums []int) [][]int {

	var res [][]int

	if nums == nil || len(nums) == 0 {
		return res
	}

	var backtrack func(idx int)
	var cur []int

	backtrack = func(idx int) {
		if idx == len(nums) {
			res = append(res, append([]int{}, cur...))
			return
		}
		// 不选nums[idx],往下走
		backtrack(idx + 1)

		// 选择nums[idx]，往下走
		cur = append(cur, nums[idx])
		backtrack(idx + 1)
		cur = cur[:len(cur)-1]
	}

	backtrack(0)

	return res
}


func subsets2(nums []int) [][]int {
	var res [][]int
	if nums == nil || len(nums) == 0 {
		return res
	}

	var backtrack func(idx int)
	var cur []int

	backtrack = func(idx int) {
		res = append(res, append([]int{}, cur...))
		for i := idx; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtrack(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0)
	return res
}