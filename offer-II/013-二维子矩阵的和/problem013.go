/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/24 12:20 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _13_二维子矩阵的和


// https://leetcode-cn.com/problems/O4NDxx/solution/er-wei-zi-ju-zhen-de-he-by-leetcode-solu-vtih/


/**
一维前缀和
 */
type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix))
	for i, row := range matrix {
		sums[i] = make([]int, len(row) + 1)
		for j, v := range row {
			sums[i][j + 1] = sums[i][j] + v
		}
	}
	return NumMatrix{sums}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += nm.sums[i][col2 + 1] - nm.sums[i][col1]
	}
	return sum
}

/**
二维前缀和
二维前缀和：f(i, j) = f(i-1, j) + f(i, j-1) - f(i-1, j-1) + matrix[i][j]
SumRegion(row1 int, col1 int, row2 int, col2 int) = f(row2, col2) - f(row1, col2) - f(row2, col1) + f(row1, col1)
*/
type NumMatrix2 struct {
	sums [][]int
}

func Constructor2(matrix [][]int) NumMatrix2 {
	m, n := len(matrix), len(matrix[0])
	sums := make([][]int, m + 1)
	sums[0] = make([]int, n + 1)
	for i, row := range matrix {
		sums[i + 1] = make([]int, n + 1)
		for j, v := range row {
			sums[i+1][j+1] = sums[i][j+1] + sums[i+1][j] - sums[i][j] + v
		}
	}
	return NumMatrix2{sums}

}

func (nm *NumMatrix2) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0

	sum = nm.sums[row2 + 1][col2 + 1] - nm.sums[row1][col2 + 1] - nm.sums[row2 + 1][col1] + nm.sums[row1][col1]
	return sum
}
