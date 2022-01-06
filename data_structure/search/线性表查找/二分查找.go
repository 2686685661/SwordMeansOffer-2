/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/6 4:09 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 线性表查找

// 要求线性表必须是顺序存储结构，且表中元素按照关键字有序排列

// 二分查找-非递归
func BinarySearch(a []int, key int) int {

	low, high := 0, len(a)-1
	for low <= high {
		mid := (low + high) >> 1
		if a[mid] > key {
			high = mid - 1
		} else if a[mid] < key {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}


// 二分查找-递归
func BinarySearchRec(a []int, key, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) >> 1
	if a[mid] > key {
		return BinarySearchRec(a, key, low, mid-1)
	} else if a[mid] < key {
		return BinarySearchRec(a, key, mid+1, high)
	} else {
		return mid
	}
}