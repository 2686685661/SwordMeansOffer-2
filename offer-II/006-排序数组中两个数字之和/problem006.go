/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/23 11:32 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _06_排序数组中两个数字之和


func twoSum(numbers []int, target int) []int {
	if numbers == nil || len(numbers) < 2 {
		return nil
	}

	begin, end := 0, len(numbers)-1
	res := make([]int, 2)

	for begin < end {
		if numbers[begin] + numbers[end] < target {
			begin++
		} else if numbers[begin] + numbers[end] > target {
			end--
		} else {
			res[0] = begin
			res[1] = end
			break
		}
	}
	return res
}