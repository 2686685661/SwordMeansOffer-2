/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/5 11:41 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _47_礼物的最大价值


func maxValue( grid [][]int ) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	return maxValueCore(grid, 0, 0)
}

func maxValueCore(grid [][]int, row, col int) int {
	if row >= len(grid) || col >= len(grid[0]) {
		return 0
	}
	return max(maxValueCore(grid, row+1, col), maxValueCore(grid, row, col+1)) + grid[row][col]
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}



func maxValue2(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	dpTable := make([][]int, len(grid))
	for i, _ := range dpTable {
		dpTable[i] = make([]int, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			val := 0
			if i - 1 >= 0 {
				val = max(val, dpTable[i - 1][j])
			}
			if j - 1 >= 0 {
				val = max(val, dpTable[i][j - 1])
			}
			dpTable[i][j] = val + grid[i][j]
		}
	}
	return dpTable[len(grid) - 1][len(grid[0]) - 1]
}