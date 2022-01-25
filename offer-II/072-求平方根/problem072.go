/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/25 1:31 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _72_求平方根

/**
二分查找
步骤：
确定左右边界；
根据中间值 mid 的平方与 x 的大小关系，确定查找的区间；
确定最后返回的是左边界还是右边界（跳出循环后）。
 */
func mySqrt(x int) int {
	low, high := 1, (x >> 1) + 1
	for low <= high {
		mid := (low + high) >> 1
		if mid * mid > x {
			high = mid - 1
		} else if mid * mid < x {
			low = mid + 1
		} else {
			return mid
		}
	}
	return high // 最后返回右边界，因为当查找区间左闭右闭时，循环退出条件时left = right + 1，而本题有向下取整的意思（8的平方根是2）
}