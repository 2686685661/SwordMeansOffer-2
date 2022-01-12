/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/7 11:34 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _44_二叉树每层的最大值

/**
dfs & bfs 参考：https://blog.csdn.net/mingwanganyu/article/details/72033122
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	//return bfs(root)
	dfs(root, 0)
	return list
}

// 广度优先遍历
//数据结构：队列
//父节点入队，父节点出队列，先左子节点入队，后右子节点入队。递归遍历全部节点即可
func bfs(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	var list []int
	queue = append(queue, root)

	for len(queue) > 0 {
		max := -1 << 31
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			max = Max(max, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		list = append(list, max)
	}
	return list
}


// 深度优先遍历
//数据结构：栈
//父节点入栈，父节点出栈，先右子节点入栈，后左子节点入栈。递归遍历全部节点即可
var list []int
func dfs(root *TreeNode, depth int) {
	if root == nil {
		return
	}

	if len(list) < depth + 1 {
		list = append(list, root.Val)
	} else if root.Val > list[depth] {
		list[depth] = root.Val
	}
	dfs(root.Left, depth + 1)
	dfs(root.Right, depth + 1)
}


func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}