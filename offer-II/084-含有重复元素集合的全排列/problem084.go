/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/29 12:07 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _84_含有重复元素集合的全排列


func permuteUnique(nums []int) [][]int {
	if len(nums) < 1 || len(nums) > 8 {
		return nil
	}
	var res [][]int
	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx == len(nums) {
			res = append(res, append([]int{}, nums...))
			return
		}

		m := make(map[int]int)
		for i := idx; i < len(nums); i++ {
			if  _, ok := m[nums[i]]; !ok || i == idx {
				m[nums[i]] = 1
				nums[i], nums[idx] = nums[idx], nums[i]
				backtrack(idx + 1)
				nums[i], nums[idx] = nums[idx], nums[i]
			}
		}
	}
	backtrack(0)
	return res
}