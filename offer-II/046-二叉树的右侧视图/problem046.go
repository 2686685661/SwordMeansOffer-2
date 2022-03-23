/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/7 1:43 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _46_二叉树的右侧视图

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	//return bfs(root)
	dfs(root, 0)
	return list

}


//bfs
func bfs(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var queue []*TreeNode
	var list []int
	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == 0 {
				list = append(list, node.Val)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
	}
	return list
}

// dfs
var (
	maxDepth = -1
	list []int
)
func dfs(root *TreeNode, depth int) {
	if root == nil {
		return
	}

	if depth > maxDepth {
		list = append(list, root.Val)
		maxDepth = depth
	}

	dfs(root.Right, depth+1)
	dfs(root.Left, depth+1)

}