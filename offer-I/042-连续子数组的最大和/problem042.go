/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/1 8:19 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _42_连续子数组的最大和


/**
https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/solution/lian-xu-zi-shu-zu-de-zui-da-he-by-leetco-tiui/
 */

/**
线段树思想
https://blog.csdn.net/zearot/article/details/48299459
https://blog.csdn.net/zearot/article/details/52280189
 */

// 动态规划
func FindGreatestSumOfSubArray( array []int ) int {

	if array == nil || len(array) == 0 {
		return 0
	}

	var (
		max = array[0]
		res = array[0]
	)

	for i := 1; i< len(array); i++ {
		max = maxResult(max + array[i], array[i])
		res = maxResult(res, max)
	}
	return res
}

func maxResult(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
