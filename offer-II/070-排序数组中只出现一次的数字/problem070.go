/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/25 12:15 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _70_排序数组中只出现一次的数字


// 位运算，异或
// 相同数字异或等于0， 任何数字与0 异或等于本身
func singleNonDuplicate(nums []int) int {
	var res int
	for _, n := range nums {
		res ^= n
	}
	return res
}

// 二分查找
// 参考：https://leetcode-cn.com/problems/skFtm2/solution/java-xiao-lu-100yi-bu-bu-dai-ni-er-fen-b-hwqa/
func singleNonDuplicate2(nums []int) int {
	low, high := 0, len(nums)
	res := -1
	for low <= high {
		mid := (low + high) >> 1
		if mid < len(nums) - 1 && nums[mid] == nums[mid + 1] {
			if mid % 2 == 0 {
				low = mid + 2
			} else {
				high = mid - 1
			}

		} else if mid > 0 && nums[mid] == nums[mid - 1] {
			if mid % 2 == 0 {
				high = mid - 2
			} else {
				low = mid + 1
			}
		} else {
			res = nums[mid]
			break
		}
	}
	return res
}