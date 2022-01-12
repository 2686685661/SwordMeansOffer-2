/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/12 11:48 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _55_二叉搜索树迭代器



type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 二叉树的中序遍历-迭代

type BSTIterator struct {
	stack []*TreeNode
	cur *TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		stack: make([]*TreeNode, 0),
		cur:   root,
	}

}


func (this *BSTIterator) Next() int {
	for node := this.cur; node != nil; node = node.Left {
		this.stack = append(this.stack, node)
	}
	this.cur, this.stack = this.stack[len(this.stack) - 1], this.stack[:len(this.stack)-1]
	val := this.cur.Val
	this.cur = this.cur.Right
	return val
}


func (this *BSTIterator) HasNext() bool {
	return this.cur != nil || len(this.stack) > 0

}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */