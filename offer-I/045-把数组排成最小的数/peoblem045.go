/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/1 8:51 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _45_把数组排成最小的数

import "strconv"

func PrintMinNumber( numbers []int ) string {

	if numbers == nil || len(numbers) == 0 {
		return ""
	}
	var str string

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if merge2Int(numbers[i], numbers[j]) > merge2Int(numbers[j], numbers[i]) {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	for _, v := range numbers {
		str += strconv.Itoa(v)
	}
	return str
}


func merge2Int(a, b int) int64 {
	s := strconv.Itoa(a) + strconv.Itoa(b)
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}