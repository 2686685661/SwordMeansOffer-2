/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/27 11:23 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _76_数组中的第_k_大的数字

func findKthLargest(nums []int, k int) int {

	// 直插排序
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i - 1] {
			tmp := nums[i]
			j := 0
			for j = i - 1; j >= 0 && nums[j] > tmp; j-- {
				nums[j + 1] = nums[j]
			}
			nums[j + 1] = tmp
		}
	}

	return nums[len(nums) - k]

}