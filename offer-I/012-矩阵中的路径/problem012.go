/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/18 5:47 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _12_矩阵中的路径

// dfs
func hasPath( matrix [][]byte ,  word string ) bool {
	visited := make([][]int, len(matrix))
	for i := range visited {
		visited[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix);i ++ {
		for j := 0; j < len(matrix[0]); j++ {
			if dfs(matrix, word, i, j, 0, visited) {
				return true
			}
		}
	}
	return false
}

func dfs(matrix [][]byte, word string, i, j, k int, visited [][]int) bool {
	if i < 0 || j < 0 || i >= len(matrix) || j >= len(matrix[0]) || visited[i][j] == 1 || k > len(word) || matrix[i][j] != word[k] {
		return false
	}

	if k == len(word) - 1 {
		return true
	}

	visited[i][j] = 1

	if dfs(matrix, word, i-1, j, k+1, visited) || dfs(matrix, word, i+1, j, k+1, visited) || dfs(matrix, word, i, j-1, k+1, visited) || dfs(matrix, word, i, j+1, k+1, visited) {
		return true
	}
	visited[i][j] = 0
	return false
}