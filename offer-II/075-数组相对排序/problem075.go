/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/27 11:00 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _75_数组相对排序


// 使用计数排序
func relativeSortArray(arr1 []int, arr2 []int) []int {
	maxVal := arr1[0]
	for _, n := range arr1 {
		if maxVal < n {
			maxVal = n
		}
	}

	counts := make([]int, maxVal + 1)

	for _, n := range arr1 {
		counts[n]++
	}


	var i int
	for _, n := range arr2 {
		for {
			if counts[n] <= 0 {
				break
			}
			arr1[i] = n
			i++
			counts[n]--
		}
	}

	for n, _ := range counts {
		for {
			if counts[n] <= 0 {
				break
			}
			arr1[i] = n
			i++
			counts[n]--
		}
	}
	return arr1

}