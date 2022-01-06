/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 1:21 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 折半插入排序


// 折半插入排序
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
// 稳定排序
func BInsertSort(elem []int) []int {
	for i := 2; i < len(elem); i++ {
		elem[0] = elem[i] // 设置哨兵
		low, high := 1, i-1
		for low <= high {
			mid := (low + high) >> 1
			if elem[0] < elem[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

		for j := i - 1;j >= high + 1; j-- {
			elem[j + 1] = elem[j]
		}
		elem[high + 1] = elem[0]
	}

	return elem
}