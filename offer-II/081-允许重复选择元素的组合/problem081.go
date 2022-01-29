/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/28 5:15 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _81_允许重复选择元素的组合


/**
回溯法
 */
import "sort"

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) < 1 || target < 1 {
		return nil
	}

	var res [][]int
	var cur []int
	var dfs func(idx, sum int)
	sort.Ints(candidates) // 排序非常关键，不仅仅是起剪枝作用，还有去重
	dfs = func(idx, sum int) {
		if idx >= len(candidates) || sum > target {
			return
		}
		if sum == target {
			res = append(res, append([]int{}, cur...))
			return
		}
		for i := 0; i < len(candidates); i++ {
			// 剪枝 + 去重
			//去重：例如target=7，candidates={2，3，7}中我们有2-2-3，2-3-2，3-2-2需要去重，可以令取的下一个数不小于当前所取的数，仅保留2-2-3，仅适用于candidates无重复项
			if sum + candidates[i] <= target && idx <= i {
				cur = append(cur, candidates[i])
				sum += candidates[i]
				dfs(i, sum)
				cur = cur[:len(cur)-1]
				sum -= candidates[i]
			}
		}
	}
	dfs(0, 0)
	return res
}