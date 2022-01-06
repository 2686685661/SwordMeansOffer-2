/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/7/13 10:04 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package linklist

import "fmt"

type Item interface {

}

type LinkNode struct {
	Paylod Item
	Next *LinkNode
}

type LinkNoder interface {
	Add(paylod Item)
	Delete(index int) Item
	Insert(index int, paylod Item)
	GetLength() int
	Search(paylod Item) int
	GetAll() []Item
	Reverse() *LinkNode
}

// Add 采用尾部插入的方式，给链表添加元素
func (head *LinkNode) Add(paylod Item) {
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	newNode := LinkNode{paylod, nil}
	tmp.Next = &newNode
}

// Delete 删除元素，并返回删除节点的值
func (head *LinkNode) Delete(index int) Item {
	tmp, j := head, 0
	for  ;j < index - 1 && tmp.Next != nil; j++ {
		tmp = tmp.Next
	}
	// index > n || index < 1的时候，删除位置不合理
	if tmp.Next == nil || j > index - 1 {
		fmt.Println("index out of range, please check")
		return "index out of range, please check"
	}
	q := tmp.Next
	tmp.Next = q.Next
	return q.Paylod
}

// Insert 插入元素，将新元素插入到单链表的第index个节点的位置，即插入在a(i-1)~a(i)之间
func (head *LinkNode) Insert(index int, paylod Item) {
	tmp, j := head, 0
	for ;tmp.Next != nil && j < index - 1; {
		tmp = tmp.Next
		j++
	}

	if tmp == nil || j > index - 1 {
		fmt.Println("index out of range, please check")
		return
	}
	newNode := LinkNode{
		Paylod: paylod,
		Next:   nil,
	}
	newNode.Next = tmp.Next
	tmp.Next = &newNode
}

// GetLength 获取单链表的长度
func (head *LinkNode) GetLength() int {
	tmp, length := head, 0
	for tmp.Next != nil {
		length++
		tmp =tmp.Next
	}
	return length
}

// Search 查询元素在链表中的位置
func (head *LinkNode) Search(paylod Item) int {
	tmp, index := head, 0

	for tmp.Next != nil {
		if tmp.Paylod == paylod {
			return index
		} else {
			index++
			tmp = tmp.Next
			if index > head.GetLength() - 1 {
				break
			}
		}
	}
	if tmp.Paylod == paylod {
		return index
	}
	return -1

}

func (head *LinkNode) GetAll() []Item {
	dataList := make([]Item, 0, head.GetLength())
	tmp := head
	for tmp != nil {
		dataList = append(dataList, tmp.Paylod)
		tmp = tmp.Next
	}
	return dataList
}

// Reverse 循环方式反转链表
func (head *LinkNode) Reverse() *LinkNode {
	if head == nil || head.Next == nil {
		return head
	}
	current := head
	var lnext *LinkNode
	var lprev *LinkNode
	var reverseNode *LinkNode

	for current != nil {
		lnext = current.Next
		current.Next = lprev
		if lnext == nil {
			reverseNode = current
		}
		lprev = current
		current = lnext
	}
	return reverseNode
}

// ReverseRecursive 递归方式反转链表
func (head *LinkNode) ReverseRecursive() *LinkNode {

	if head == nil || head.Next == nil {
		return head
	}

	lnext := head.Next
	newHead := lnext.ReverseRecursive()
	lnext.Next = head
	head.Next = nil
	return newHead
}

// NewLinkNode 前插法 创建链表
func NewLinkNode(length int) *LinkNode {
	if length <= 0 {
		fmt.Println("长度需要大于0")
		return nil
	}

	head := &LinkNode{}
	for i := 0; i < length; i++ {
		newNode := &LinkNode{
			Paylod: i,
			Next:   nil,
		}
		newNode.Next = head.Next
		head.Next = newNode
	}
	return head
}



