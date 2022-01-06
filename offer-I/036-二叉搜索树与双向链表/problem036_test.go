/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/25 5:02 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _36_二叉搜索树与双向链表

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	a := &TreeNode{
		Val:   10,
		Left:  nil,
		Right: nil,
	}
	b := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	c := &TreeNode{
		Val:   14,
		Left:  nil,
		Right: nil,
	}
	d := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	e := &TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	f := &TreeNode{
		Val:   12,
		Left:  nil,
		Right: nil,
	}
	g := &TreeNode{
		Val:   16,
		Left:  nil,
		Right: nil,
	}

	a.Left = b
	a.Right = c
	b.Left = d
	b.Right = e
	c.Left = f
	c.Right = g
	res := Convert(a)



	tmp := res
	for tmp != nil {
		fmt.Println(tmp)
		tmp = res.Right
	}

}