/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/29 11:07 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _82_含有重复元素集合的组合


/**
回溯法
 */
import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) < 1 || target < 1 {
		return nil
	}

	var res [][]int
	var cur []int
	var backtrack func(idx, sum int)
	sort.Ints(candidates)
	backtrack = func(idx, sum int) {
		if sum == target {
			res = append(res, append([]int{}, cur...))
			return
		}
		if idx >= len(candidates) || sum > target {
			return
		}

		m := make(map[int]int)  // 记录之前是否存在重复，如果之前有重复元素，跳过
		for i := idx; i < len(candidates); i++ {
			if _, ok := m[candidates[i]]; !ok && sum + candidates[i] <= target {
				m[candidates[i]] = 1
				cur = append(cur, candidates[i])
				sum += candidates[i]
				backtrack(i + 1, sum)
				cur = cur[:len(cur)-1]
				sum -= candidates[i]
			}
		}
	}
	backtrack(0, 0)
	return res
}