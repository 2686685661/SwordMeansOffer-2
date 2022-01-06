/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/6 1:56 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _map
/**
	无向图的邻接矩阵表示法
	Adjacency Matrix Graph
 */

type VerTextType byte // 顶点类型
type ArcType int // 边的权值类型

type AMGraph struct {
	VerTextType []VerTextType // 顶点表
	ArcType [][]ArcType // 邻接矩阵
	verNum int // 图的顶点数
	arcNum int // 图的边树
}

func CreateUDN(verList []VerTextType) *AMGraph {
	max := len(verList)
	arcList := make([][]ArcType, max)
	for i := 0; i < max; i++ {
		arcCol := make([]ArcType, max)
		for j := 0; j < max; j++ {
			arcCol[j] = 1<<32 - 1
		}
		arcList[i] = arcCol
	}
	
	return &AMGraph{
		VerTextType: verList,
		ArcType:     arcList,
		verNum:      max,
	}
}