/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/6 3:04 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _map


/*
	十字链表
 */

// 弧节点结构
type ArcLinkedNode struct {
	TailVexId int // 弧尾节点ID
	HeadVexId int // 弧头节点ID
	HeadLink *ArcLinkedNode // 头指针
	TailLink *ArcLinkedNode // 尾指针
}

// 顶点节点结构
type VertexNode struct {
	Data rune
	FirstInArcPrt *ArcLinkedNode // 入度边的指针，其实就是指向该顶点第一条入弧节点
	FirstOutArcPtr *ArcLinkedNode // 出度边的指针，其实就是指向该顶点第一条出弧节点
}

// 十字链表
type AdjoinTable struct {
	Table []*VertexNode
	VerNum int // 顶点的数量
	ArcNum int // 弧的数量
}

