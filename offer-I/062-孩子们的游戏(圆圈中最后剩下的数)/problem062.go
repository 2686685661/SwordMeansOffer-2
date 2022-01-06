/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/17 12:08 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _62_孩子们的游戏_圆圈中最后剩下的数_

func LastRemaining_Solution( n int ,  m int ) int {
	if n < 1 || m == 0 {
		return n
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}

	var (
		pos int = -1
		step int
		count int = n
	)

	for count > 0 {
		pos++
		if pos >= n {
			pos = 0
		}
		if nums[pos] == -1 {
			continue
		}
		step++
		if step == m {
			nums[pos] = -1
			step = 0
			count--
		}
	}

	return pos
}