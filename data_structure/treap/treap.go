/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/13 2:10 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package treap

import (
	"fmt"
	"math/rand"
)

/*
树堆

相关参考：
Treap介绍: https://www.cnblogs.com/guoshaoyang/p/11300886.html
树堆（Treap）图文详解与实现: https://blog.csdn.net/K346K346/article/details/50808879

树旋转: https://zh.wikipedia.org/wiki/%E6%A0%91%E6%97%8B%E8%BD%AC
 */


// 树节点
type node struct {
	priority int
	val int
	ch [2]*node
}

// 二分比较
func (n *node) contrast(val int) (res int) {
	res = -1
	if val < n.val {
		res = 0
	} else if val > n.val {
		res = 1
	}
	return
}

// 树旋转
// d=0 对o进行左旋
// d=1 对o进行右旋
func (n *node) rotate(d int) *node {
	pivot := n.ch[d ^ 1] // 左旋，则转轴是o的右子节点；右旋，则转轴是o的左子节点
	n.ch[d ^ 1] = pivot.ch[d]
	pivot.ch[d] = n
	return pivot
}

// 中序遍历打印
func inorder(n *node) {
	if n == nil {
		return
	}
	inorder(n.ch[0])
	fmt.Printf("%d\t", n.val)
	inorder(n.ch[1])
}


// 树堆结构
type treap struct {
	root *node
}


// 插入树堆
func (t *treap) Put(val int) {
	t.root = t.put(t.root, val)
}

func (t *treap) put(root *node, val int) *node {
	if root == nil {
		return &node{
			priority: rand.Int(),
			val:      val,
		}
	}
	d := root.contrast(val)
	root.ch[d] = t.put(root.ch[d], val)
	if root.ch[d].priority > root.priority {
		root = root.rotate(d ^ 1) //插入节点是o的左子节点，进行右旋；插入节点是o的右子节点，进行左旋
	}
	return root
}


// 从树堆中删除
func (t *treap) Del(val int) {
	t.root = t.del(t.root, val)
}

func (t *treap) del(root *node, val int) *node {
	if d := root.contrast(val); d >= 0 {
		root.ch[d] = t.del(root.ch[d], val)
		return root
	}

	if root.ch[0] == nil {
		return root.ch[1]
	}
	if root.ch[1] == nil {
		return root.ch[0]
	}

	d := 0
	if root.ch[0].priority > root.ch[1].priority {
		d = 1
	}
	root = root.rotate(d)
	root.ch[d] = t.del(root.ch[d], val)
	return root
}