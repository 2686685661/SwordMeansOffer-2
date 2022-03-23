/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/25 11:47 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _68_查找插入位置


// 二分查找
func searchInsert(nums []int, target int) int {
	if nums == nil || len(nums) < 1 {
		return -1
	}

	var mid int
	low, high := 0, len(nums) - 1
	for low <= high {
		mid = (low + high) >> 1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}

	if nums[mid] > target {
		return mid
	} else {
		return mid + 1
	}
}