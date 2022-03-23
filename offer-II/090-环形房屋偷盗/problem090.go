/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/8 6:04 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _90_环形房屋偷盗

/**
动态规划
https://leetcode-cn.com/problems/PzWKhm/solution/jian-zhi-offer-2-mian-shi-ti-90-shu-zhon-c9yi/
 */
func rob(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	helper := func(nums []int, start, end int) int {
		if start == end {
			return nums[start]
		}
		dp := make([]int, len(nums))
		dp[0], dp[1] = nums[start], max(nums[start], nums[start + 1])
		for i := start + 2; i <= end; i++ {
			dp[i] = max(dp[i - 2] + nums[i], dp[i - 1])
		}
		return dp[end]
	}

	n := len(nums)
	return max(helper(nums, 0, n - 2), helper(nums, 1, n - 1))
}