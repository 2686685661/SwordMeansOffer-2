/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/19 1:52 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _13_机器人的运动范围


// 回溯 + DFS
func movingCount( threshold int ,  rows int ,  cols int ) int {
	if threshold < 0 || rows <= 0 || cols <= 0 {
		return 0
	}

	visited := make([][]int, rows)
	for i := range visited {
		visited[i] = make([]int, cols)
	}

	count := movingCountCore(threshold, rows, cols, 0, 0, visited)
	return count
}

func movingCountCore(threshold, rows, cols, i, j int, visited [][]int) int {
	count := 0
	if check(threshold, rows, cols, i, j, visited) {
		visited[i][j] = 1
		count = 1 + movingCountCore(threshold, rows, cols, i-1, j, visited) + movingCountCore(threshold, rows, cols, i+1, j, visited) + movingCountCore(threshold, rows, cols, i, j-1, visited) + movingCountCore(threshold, rows, cols, i, j+1, visited)
	}
	return count
}

func check(threshold, rows, cols, i, j int, visited [][]int) bool {
	if i >=0 && j >= 0 && i <rows && j < cols && visited[i][j] == 0 && getDigitSum(i) + getDigitSum(j) <= threshold {
		return true
	}
	return false
}

func getDigitSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}