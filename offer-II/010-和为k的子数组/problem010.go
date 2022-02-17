/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/23 2:40 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _10_和为_k_的子数组


/**
前缀和算法
https://leetcode-cn.com/problems/QTMn0o/solution/shua-chuan-jian-zhi-offer-day07-shu-zu-i-jdnu/
 */
func subarraySum(nums []int, k int) int {

	if nums == nil || len(nums) < 1 {
		return 0
	}

	ret, preSum := 0, 0
	dict := make(map[int]int)
	dict[0] = 1

	for _, n := range nums {
		preSum += n
		ret += dict[preSum - k]
		dict[preSum] += 1
	}
	return ret
}