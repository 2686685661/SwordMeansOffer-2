/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/25 1:35 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _34_二叉树中和为某一值的路径

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func FindPath( root *TreeNode ,  expectNumber int ) [][]int {
	if root == nil {
		return nil
	}
	var (
		res [][]int
		stack []*TreeNode
		count int
	)
	res = find(root, stack, count, expectNumber, res)
	return res
}

func find(root *TreeNode, stack []*TreeNode, count, expectNumber int, res [][]int) [][]int {
	stack = append(stack, root)
	count += root.Val
	if root.Left == nil && root.Right == nil && count == expectNumber {
		var tmp []int
		for _, v := range stack {
			tmp = append(tmp, v.Val)
		}
		res = append(res, tmp)
		fmt.Println(res)
	}
	if root.Left != nil {
		res = find(root.Left, stack, count, expectNumber, res)
	}
	if root.Right != nil {
		res = find(root.Right, stack, count, expectNumber, res)
	}
	stack = stack[:len(stack)-1]
	return res
}