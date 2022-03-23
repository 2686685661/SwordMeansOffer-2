/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/6 2:03 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _map
/**
	邻接表表示法
	Adjacency List
 */

type arcNode struct {
	vertextID int // 顶点位置ID
	nextArc *arcNode // 另一个顶点对应的节点
	weight int // 权重
}

type verTextHeader struct {
	ID int // 顶点位置id
	data rune // 顶点域值
	firstArc *arcNode // 第一个对应的边
}

type ALGraph struct {
	adjList []*verTextHeader // 邻接表
	verNum int // 顶点数
	arcNUm int // 边数量
}