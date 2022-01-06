/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/3 4:14 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package tree

import (
	"testing"
)

func TestBiTreeNode_CreateTree(t *testing.T) {
	// 先序遍历算法
	preStr1 := "AB##C##"
	preT := NewBiTree(preStr1, PreOrderTraverse)
	preTree := preT.CreateTree()
	preStr2 := preT.Tree2String(preTree)
	if preStr1 != preStr2 {
		t.Errorf("pre expect:%s, actual:%s\n", preStr1, preStr2)
	}

	// 后序遍历
	postStr1 := "##B##CA"
	postT := NewBiTree(postStr1, PostOrderTraverse)
	postTree := postT.CreateTree()
	postStr2 := postT.Tree2String(postTree)
	if postStr1 != postStr2 {
		t.Errorf("post expect:%s, actual:%s\n", postStr1, postStr2)
	}
}

func TestBiTreeNode_Copy(t *testing.T) {

	preStr1 := "AB##C##"
	preT := NewBiTree(preStr1, PreOrderTraverse)
	preTree := preT.CreateTree()
	copyTree := preT.Copy(preTree)

	t1 := preT.Tree2String(preTree)
	t2 := preT.Tree2String(copyTree)
	if t1 != t2 {
		t.Errorf("copy pre: expect:%s, actual:%s\n", t1, t2)
	}
}

func TestBiTreeNode_Depth(t *testing.T) {
	preStr1 := "AB##C##"
	preT := NewBiTree(preStr1, PreOrderTraverse)
	preTree := preT.CreateTree()
	d := preT.Depth(preTree)
	if d != 2 {
		t.Errorf("depth : expect:%d, actual:%d\n", 2, d)
	}
}

func TestBiTreeNode_NodeCount(t *testing.T) {
	preStr1 := "AB##C##"
	preT := NewBiTree(preStr1, PreOrderTraverse)
	preTree := preT.CreateTree()
	c := preT.NodeCount(preTree)
	if c != 3 {
		t.Errorf("nodecount : expect:%d, actual:%d\n", 3, c)
	}
}

func TestBiTreeNode_LeadCount(t *testing.T) {
	preStr1 := "AB##C##"
	preT := NewBiTree(preStr1, PreOrderTraverse)
	preTree := preT.CreateTree()
	c := preT.LeadCount(preTree)
	if c != 2 {
		t.Errorf("leadcount : expect:%d, actual:%d\n", 2, c)
	}
}

func TestBiTreeNode_Tree2StringNoRecursivePre(t *testing.T) {
	preStr1 := "AB##C##"
	preT := NewBiTree(preStr1, PreOrderTraverse)
	preTree := preT.CreateTree()
	preStr2 := preT.Tree2StringNoRecursivePre(preTree)
	if preStr1 != preStr2 {
		t.Errorf("no recursive pre expect:%s, actual:%s\n", preStr1, preStr2)
	}

}

func TestBiTreeNode_Tree2StringNoRecursiveIn(t *testing.T) {
	preStr1 := "AB##C##"
	preT := NewBiTree(preStr1, PreOrderTraverse)
	preTree := preT.CreateTree()
	preStr2 := preT.Tree2StringNoRecursiveIn(preTree)
	if "#B#A#C#" != preStr2 {
		t.Errorf("no recursive in expect:%s, actual:%s\n", preStr1, preStr2)
	}
}


func TestBiTreeNode_Tree2StringNoRecursivePost(t *testing.T) {
	postStr1 := "##B##CA"
	postT := NewBiTree(postStr1, PostOrderTraverse)
	postTree := postT.CreateTree()
	postStr2 := postT.Tree2StringNoRecursivePost(postTree)
	if postStr1 != postStr2 {
		t.Errorf("no recursive post expect:%s, actual:%s\n", postStr1, postStr2)
	}
}


func TestBiTreeNode_TreeSequence(t *testing.T) {
	preStr1 := "AB##C##"
	preT := NewBiTree(preStr1, PreOrderTraverse)
	preTree := preT.CreateTree()
	str := preT.TreeSequence(preTree)
	if str != "ABC" {
		t.Errorf("no recursive post expect:%s, actual:%s\n", "ABC", str)
	}
}