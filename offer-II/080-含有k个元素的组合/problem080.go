/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/28 5:01 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _80_含有_k_个元素的组合

/**
回溯法
 */

func combine(n int, k int) [][]int {

	if n < 1 || k < 1 || k > n {
		return nil
	}

	var res [][]int
	var cur []int

	var backtrack func(idx int)
	backtrack = func(idx int) {
		if idx > n || len(cur) == k {
			if len(cur) == k {
				res = append(res, append([]int{}, cur...))
			}
			return
		}
		for i := idx; i <= n; i++ {
			cur = append(cur, i)
			backtrack(i + 1)
			cur = cur[:len(cur)-1]
		}
	}

	backtrack(1)
	return res
}