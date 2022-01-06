/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/21 1:25 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _29_顺时针打印矩阵

func PrintMatrix( matrix [][]int ) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	var res []int
	s := 0
	for s * 2 < len(matrix) && s * 2 < len(matrix[0]) {
		res = printMatrixCore(matrix, s, res)
		s += 1
	}
	return res
}

func printMatrixCore(matrix [][]int, s int, res []int) []int {
	ex, ey := len(matrix[0]) - 1 - s,  len(matrix) - 1 - s

	// 从左到右 打印一行
	for i := s; i <= ex; i++ {
		res = append(res, matrix[s][i])
	}

	// 从上到下 打印一列
	if s < ey {
		for i := s + 1; i <= ey; i++ {
			res = append(res, matrix[i][ex])
		}
	}

	// 从右到左 打印一行
	if s < ey && s < ex {
		for i := ex - 1; i >= s; i-- {
			res = append(res, matrix[ey][i])
		}
	}

	// 从下到上 打印一列
	if s < ex && s < ey - 1 {
		for i := ey - 1; i >= s + 1; i-- {
			res = append(res, matrix[i][s])
		}
	}

	return res
}