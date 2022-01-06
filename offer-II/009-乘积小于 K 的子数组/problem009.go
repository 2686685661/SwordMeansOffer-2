/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/23 1:42 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _09_乘积小于_K_的子数组

/**
比如某次遍历符合题意的子数组为 ABCX，那么在该条件下符合条件的有X，CX，BCX，ABCX共四个（可以进行多个例子，发现个数符合right-left+1）
我们可能会有疑问：AB，BC也算，为什么不算进去？
记住一点我们是以最右边的X为必要条件，进行计算符合条件的子数组，否则会出现重复的！
比如在X为右侧边界时（ABCX），我们把BC算进去了，可是我们在C为最右侧时（ABC），BC已经出现过，我们重复加了BC这个子数组两次！
换言之，我们拆分子数组时，让num[right]存在，能避免重复计算。
 */
func numSubarrayProductLessThanK(nums []int, k int) int {

	if nums == nil || len(nums) < 1 || k < 0 {
		return 0
	}

	sum, count := 1, 0
	for i, j := 0, 0; j < len(nums); j++ {
		sum *= nums[j]
		for i <= j && sum >= k {
			sum /= nums[i]
			i++
		}
		if i <= j {
			count += j - i + 1
		}

	}
	return count

}