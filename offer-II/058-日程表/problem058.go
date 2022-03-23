/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/18 3:32 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _58_日程表

import (
	"math/rand"
)



// 节点
type Node struct {
	val []int
	priority int
	ch [2]*Node
}

// 树堆
type Treap struct {
	root *Node
}

// 比较
func (n *Node) Compare(data []int) int {
	if data[0] < n.val[0] {
		return 0
	} else if data[0] > n.val[0] {
		return 1
	}
	return -1
}


// 节点旋转
func (n *Node) Rotate(d int) *Node {
	pivot := n.ch[d ^ 1]
	n.ch[d ^ 1] = pivot.ch[d]
	pivot.ch[d] = n
	return pivot
}

// 插入
func (t *Treap) Insert(data []int) {
	t.root = t._insert(t.root, data)
}
func (t *Treap) _insert(node *Node, data []int) *Node {
	if node == nil {
		return &Node{
			val:      data,
			priority:  rand.Int(),
		}
	}
	d := node.Compare(data)
	if d < 0 {
		return node
	}
	node.ch[d] = t._insert(node.ch[d], data)
	if node.ch[d].priority > node.priority {
		node = node.Rotate(d ^ 1)
	}
	return node
}


// 查找小于等于目标值的最大值
func (t *Treap) PrevNode(data []int) (n *Node) {
	for o := t.root; o != nil; {
		if d := o.Compare(data); d == 0 {
			o =o.ch[0]
		} else if d > 0 {
			n = o
			o = o.ch[1]
		} else {
			return o
		}
	}
	return
}

// 查找大于等于目标值的最小值
func (t *Treap) NextNode(data []int) (n *Node) {
	for o := t.root; o != nil; {
		if d := o.Compare(data); d == 0 {
			n = o
			o = o.ch[0]
		} else if d > 0 {
			o = o.ch[1]
		} else {
			return o
		}
	}
	return
}

type MyCalendar struct {
	treap *Treap
}


func Constructor() MyCalendar {
	return MyCalendar{treap:&Treap{}}
}


func (this *MyCalendar) Book(start int, end int) bool {
	prevNode, nextNode := this.treap.PrevNode([]int{start, end}), this.treap.NextNode([]int{start, end})
	if this.treap.root == nil ||
		(prevNode == nil && end <= nextNode.val[0]) ||
		(nextNode == nil && start >= prevNode.val[1]) ||
		(prevNode != nil && nextNode != nil && end <= nextNode.val[0] && start >= prevNode.val[1]) {
		this.treap.Insert([]int{start, end})
		return true
	}
	return false
}


/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */