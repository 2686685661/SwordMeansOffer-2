/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/29 9:41 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _74_和为S的连续正数序列


func FindContinuousSequence( sum int ) [][]int {

	var res [][]int
	pl, ph := 1, 2

	for ph > pl {
		cur := ((pl + ph) * (ph - pl + 1)) / 2
		if cur == sum {
			var tmp []int
			for i := pl; i <= ph; i++ {
				tmp = append(tmp, i)
			}
			res = append(res, tmp)
			pl++
		} else if cur < sum {
			ph++
		} else {
			pl++
		}
	}
	return res
}