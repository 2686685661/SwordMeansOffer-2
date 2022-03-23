/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/17 11:43 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _61_扑克牌顺子

import "sort"

// 排序
// 统计王的数量
// 计算间隔数
// 间隔数<=王数量，是顺子，否则不是。如果发现有对子，则不是顺子
func IsContinuous( numbers []int ) bool {
	if numbers == nil || len(numbers) != 5 {
		return false
	}

	sort.Sort(sort.IntSlice(numbers))

	numZero := 0
	numSpace := 0
	for _, v := range numbers {
		if v == 0 {
			numZero += 1
		} else {
			break
		}
	}

	front, after := numZero, numZero + 1

	for after < len(numbers) {
		if numbers[front] == numbers[after] {
			return false
		}
		numSpace += numbers[after] - numbers[front] - 1
		front = after
		after += 1
	}

	if numSpace <= numZero {
		return true
	} else {
		return false
	}
}