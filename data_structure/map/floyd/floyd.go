/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/15 11:02 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package floyd

const (
	MAXVEX = 9
	MAXWEIGHT = 65535
)
var (
	Patharc [][]int = make([][]int, MAXVEX)
	ShortPathTable [MAXVEX][MAXVEX]int
)


func ShortestPath_Floyd() {
		for i := 0; i < MAXVEX; i++ {
			Patharc[i] = make([]int, MAXVEX)
		}

		graph := NewGraph()
		ShortPathTable = graph

		var (
			k int
			v int
			w int
		)

		for v = 0; v < MAXVEX; v++ {
			for w = 0; w < MAXVEX; w++ {
				Patharc[v][w] = w
			}
		}

		for k = 0; k < MAXVEX; k++ {
			for v = 0; v < MAXVEX; v++ {
				for w = 0; w < MAXVEX; w++ {
					if ShortPathTable[v][w] > ShortPathTable[v][k] + ShortPathTable[k][w] {
						ShortPathTable[v][w] = ShortPathTable[v][k] + ShortPathTable[k][w]
						Patharc[w][w] = k
					}
				}
			}
		}


		// 结果打印

}

func NewGraph() [MAXVEX][MAXVEX]int {
	var graph [MAXVEX][MAXVEX]int
	var v0 = [MAXVEX]int{0, 1, 5, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
	var v1 = [MAXVEX]int{1, 0, 3, 7, 5, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
	var v2 = [MAXVEX]int{5, 3, 0, MAXWEIGHT, 1, 7, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT}
	var v3 = [MAXVEX]int{MAXWEIGHT, 7, MAXWEIGHT, 0, 2, MAXWEIGHT, 3, MAXWEIGHT, MAXWEIGHT}
	var v4 = [MAXVEX]int{MAXWEIGHT, 5, 1, 2, 0, 3, 6, 9, MAXWEIGHT}
	var v5 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, 7, MAXWEIGHT, 3, 0, MAXWEIGHT, 5, MAXWEIGHT}
	var v6 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 3, 6, MAXWEIGHT, 0, 2, 7}
	var v7 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 9, 5, 2, 0, 4}
	var v8 = [MAXVEX]int{MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, MAXWEIGHT, 7, 4, 0}
	graph[0] = v0
	graph[1] = v1
	graph[2] = v2
	graph[3] = v3
	graph[4] = v4
	graph[5] = v5
	graph[6] = v6
	graph[7] = v7
	graph[8] = v8
	return graph
}