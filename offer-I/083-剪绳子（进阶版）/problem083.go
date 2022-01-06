/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/2 5:55 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _83_剪绳子_进阶版_

// https://www.nowcoder.com/practice/106f666170554379ab1974e5a601e741?tpId=13&tags=&title=&difficulty=0&judgeStatus=0&rp=0
func cutRope( number int64 ) int64 {
	if number > 0 && number <= 3 {
		return number - 1
	}
	products := make([]int64, number + 1)
	products[1] = 1

	var i int64
	for i = 2; i <= number; i++ {
		max := i
		var j int64
		for j = 1; j <= i / 2; j++ {
			if products[j] * products[i - j] > max {
				max = products[j] * products[i - j]
			}
		}
		products[i] = max
	}

	var res int64 = -1
	for i = 1; i <= number / 2; i++ {
		if products[i] * products[number - i] > res {
			res = products[i] * products[number - i]
		}
	}

	return res
}