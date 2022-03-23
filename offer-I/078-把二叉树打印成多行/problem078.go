/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/30 1:47 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _78_把二叉树打印成多行


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 循环
func Print( pRoot *TreeNode ) [][]int {

	if pRoot == nil {
		return nil
	}

	var (
		queue []*TreeNode
		result [][]int
	)
	queue = append(queue, pRoot)

	for len(queue) > 0 {
		b, e := 0, len(queue) - 1
		var res []int
		for i := b; i <= e; i++ {
			node := queue[0]
			queue = queue[1:]
			res = append(res, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, res)
	}
	return result
}

// 递归
func Print2( pRoot *TreeNode ) [][]int {
	var result [][]int
	Depth(pRoot, 1, &result)
	return result
}

func Depth(root *TreeNode, depth int, result *[][]int) {
	if root == nil {
		return
	}
	if depth > len(*result) {
		*result = append(*result, []int{})
	}
	(*result)[depth - 1] = append((*result)[depth - 1], root.Val)

	Depth(root.Left, depth + 1, result)
	Depth(root.Right, depth + 1, result)
}