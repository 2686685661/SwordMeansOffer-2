/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/7/14 1:15 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package linklist

import (
	"fmt"
	"testing"
)

func TestLinkNode_Add(t *testing.T) {
	head := NewLinkNode(10)
	fmt.Println("length: ", head.GetLength())
	head.Add(15)
	for head != nil{
		//t.Errorf("%s", head.Paylod)
		fmt.Println(head.Paylod)
		head = head.Next
	}
}

func TestLinkNode_Delete(t *testing.T) {
	head := NewLinkNode(10)
	tmp := head
	for tmp != nil{
		fmt.Print(tmp.Paylod, ",")
		tmp = tmp.Next
	}
	head.Delete(1)
	fmt.Println("------")
	//head.Delete(3)
	for head != nil{
		fmt.Print(head.Paylod, ",")
		head = head.Next
	}
}

func TestLinkNode_Insert(t *testing.T) {
	head := NewLinkNode(10)
	head.Insert(10, 999)
	for head != nil{
		fmt.Print(head.Paylod, ",")
		head = head.Next
	}
}

func TestLinkNode_Search(t *testing.T) {
	head := NewLinkNode(10)
	fmt.Println(head.Search(8))
}

func TestLinkNode_Reverse(t *testing.T) {
	head := NewLinkNode(10)
	tmp := head.Reverse()
	for tmp != nil{
		fmt.Print(tmp.Paylod, ",")
		tmp = tmp.Next
	}
}

func TestLinkNode_ReverseRecursive(t *testing.T) {
	head := NewLinkNode(10)
	tmp := head.ReverseRecursive()
	for tmp != nil{
		fmt.Print(tmp.Paylod, ",")
		tmp = tmp.Next
	}
}
