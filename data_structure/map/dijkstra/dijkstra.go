/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/15 10:46 下午, by lishanlei, create
*/

/*
DESCRIPTION
图的最短路径---迪杰斯特拉(Dijkstra)算法浅析
*/

package dijkstra

import "fmt"

const (
	MAXVEX = 9
	MAXWEIGHT = 65535
)
var (
	Patharc []int = make([]int, MAXVEX)
	ShortPathTable []int = make([]int, MAXVEX)
)

func ShortestPath_Dijkstra(v0 int) {
	graph := NewGraph()
	fin := make([]int, MAXVEX)

	var (
		v int = 0
		w int = 0
		k int = 0
	)
	for v = 0; v < MAXVEX; v++ {
		fin[v] = 0
		Patharc[v] = 0
		ShortPathTable[v] = graph[v0][v]
	}
	ShortPathTable[v0] = 0
	fin[v0] = 1

	for v = 1; v < MAXVEX; v++ {
		min := MAXWEIGHT

		for w = 0; w < MAXVEX; w++ {
			if fin[w] == 0 && ShortPathTable[w] < min {
				min = ShortPathTable[w]
				k = w
			}
		}
		fin[k] = 1

		for w = 0; w < MAXVEX; w++ {
			if fin[w] == 0 && min + graph[k][w] < ShortPathTable[w] {
				ShortPathTable[w] = min + graph[k][w]
				Patharc[w] = k
			}
		}
		fmt.Println("遍历完V", v, "后:", ShortPathTable)
	}

	//输出
	for i := 0; i < len(ShortPathTable); i++ {
		fmt.Println("V0到V", i, "最小路径:", ShortPathTable[i])
	}

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