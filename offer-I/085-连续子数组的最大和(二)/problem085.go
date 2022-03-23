/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/6 8:29 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _85_连续子数组的最大和_二_

import "math"

func FindGreatestSumOfSubArray( array []int ) []int {
	dp := make([]int, len(array))
	for i, v := range array {
		dp[i] = v
	}
	maxStart, maxEnd := 0, 1
	start, end := 0, 1
	for i := 1; i < len(array); i++ {
		dp[i] = int(math.Max(float64(array[i] + dp[i - 1]), float64(array[i])))
		if array[i] + dp[i - 1] >= array[i] {
			end += 1
		} else {
			start, end = i, i + 1
		}
		if sum(array[start:end]) > sum(array[maxStart:maxEnd]) {
			maxStart, maxEnd = start, end
		} else if sum(array[start:end]) == sum(array[maxStart:maxEnd]) {
			if len(array[start:end]) > len(array[maxStart:maxEnd]) {
				maxStart, maxEnd = start, end
			}
		} else {
			continue
		}
	}
	return array[maxStart:maxEnd]

}

func sum(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}