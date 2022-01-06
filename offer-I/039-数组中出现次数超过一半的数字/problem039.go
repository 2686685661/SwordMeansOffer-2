/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/28 9:43 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _39_数组中出现次数超过一半的数字

func MoreThanHalfNum_Solution( numbers []int ) int {
	if numbers == nil || len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}

	var (
		result = numbers[0]
		count = 1
	)
	for _, v := range numbers {
		if result == v {
			count++
		} else {
			count--
		}
		if count == 0 {
			count = 1
			result = v
		}
	}

	count = 0
	for _, v := range numbers {
		if v == result {
			count += 1
		}
	}
	if count * 2 < len(numbers)-1 {
		result = 0
	}
	return result
}