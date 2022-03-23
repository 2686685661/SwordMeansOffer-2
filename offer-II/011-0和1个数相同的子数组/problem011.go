/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/23 5:41 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _11_0_和_1_个数相同的子数组

import "fmt"

// 前缀和算法
func findMaxLength(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}

	for i, n := range nums {
		if n == 0 {
			nums[i] = -1
		}
	}

	res, preSum := 0, 0
	dict := make(map[int]int)

	//假如前5个元素和为0，那么子数组的长度为 4-（-1）=5
	dict[0] = -1

	for i, n := range nums {
		preSum += n

		if preIndex, ok := dict[preSum]; ok {
			res = Max(res, i - preIndex)
		} else {
			dict[preSum] = i
		}

	}
	fmt.Println(dict)

	return res

}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}