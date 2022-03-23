/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/10 11:15 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _93_最长斐波那契数列


/**
动态规划
https://leetcode-cn.com/problems/Q91FMA/solution/jian-zhi-offer-2-mian-shi-ti-93-shu-zhon-2ww4/
 */


// 动态规划 + hash
func lenLongestFibSubseq(arr []int) int {
	if arr == nil || len(arr) < 3 {
		return 0
	}

	n := len(arr)

	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	mp := make(map[int]int)
	for i, v := range arr {
		mp[v] = i
	}

	res := 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			tmp := arr[i] - arr[j]
			if k, ok := mp[tmp]; ok && k < j {
				dp[i][j] = dp[j][k] + 1
			} else {
				dp[i][j] = 2
			}
			res = max(res, dp[i][j])
		}
	}
	if res > 2 {
		return res
	} else {
		return 0
	}

}


// 动态规划 + 二分查找
func lenLongestFibSubseq2(arr []int) int {
	if arr == nil || len(arr) < 3 {
		return 0
	}

	n := len(arr)

	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	res := 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	search := func(arr []int, right, target int) int {
		left := 0
		for left <= right {
			mid := (left + right) >> 1
			if arr[mid] == target {
				return mid
			}
			if arr[mid] > target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		return -1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			index := search(arr, j - 1, arr[i] - arr[j])
			if index != -1 {
				dp[i][j] = dp[j][index] + 1
			} else {
				dp[i][j] = 2
			}
			res = max(res, dp[i][j])
		}
	}
	if res > 2 {
		return res
	} else {
		return 0
	}

}


