/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/8 5:42 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _89_房屋偷盗

/**
动态规划
https://leetcode-cn.com/problems/Gu0c2T/solution/jian-zhi-offer-2-mian-shi-ti-89-shu-zhon-86xo/
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

	dp := make([]int, len(nums))

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i - 2] + nums[i], dp[i - 1])
	}
	return dp[len(nums) - 1]

}