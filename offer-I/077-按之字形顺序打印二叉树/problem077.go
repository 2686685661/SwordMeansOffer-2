/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/30 12:11 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _77_按之字形顺序打印二叉树


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Print( pRoot *TreeNode ) [][]int {
	if pRoot == nil {
		return nil
	}
	var (
		stack1 []*TreeNode
		stack2 []*TreeNode
		result [][]int
	)
	stack1 = append(stack1, pRoot)

	for len(stack1) > 0 || len(stack2) > 0 {
		var (
			ret1 []int
			ret2 []int
		)

		for len(stack1) > 0 {
			node := stack1[len(stack1)-1]
			if node.Left != nil {
				stack2 = append(stack2, node.Left)
			}
			if node.Right !=nil {
				stack2 = append(stack2, node.Right)
			}
			ret1 = append(ret1, node.Val)
			stack1 = stack1[:len(stack1)-1]
		}
		if len(ret1) > 0 {
			result = append(result, ret1)
		}

		for len(stack2) > 0 {
			node := stack2[len(stack2)-1]
			if node.Right != nil {
				stack1 = append(stack1, node.Right)
			}
			if node.Left != nil {
				stack1 = append(stack1, node.Left)
			}
			ret2 = append(ret2, node.Val)
			stack2 = stack2[:len(stack2)-1]
		}
		if len(ret2) > 0 {
			result = append(result, ret2)
		}
	}
	return result
}