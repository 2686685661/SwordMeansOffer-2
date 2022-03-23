/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/13 3:37 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _57_值和下标之差都在给定的范围内

import "math/rand"

/**
参考：https://leetcode-cn.com/problems/7WqeDu/solution/zhi-he-xia-biao-zhi-chai-du-zai-gei-ding-94ei/
 */

/**
滑动窗口 + 有序集合
有序集合实现：树堆
 */
type Node struct {
	val int
	priority int
	ch [2]*Node
}

type Treap struct {
	root *Node
}

func (n *Node) Cmp(val int) (res int) {
	res = -1
	if val < n.val {
		res = 0
	} else if val > n.val {
		res = 1
	}
	return
}

// d = 0, 对节点n进行左旋
// d = 1, 对节点n进行右旋
func (n *Node) Rotate(d int) *Node {
	pivot := n.ch[d ^ 1]
	n.ch[d ^ 1] = pivot.ch[d]
	pivot.ch[d] = n
	return pivot
}

func (t *Treap) Put(val int) {
	t.root = t._put(t.root, val)
}
func (t *Treap) _put(node *Node, val int) *Node {
	if node == nil {
		return &Node{
			val:      val,
			priority: rand.Int(),
		}
	}
	d := node.Cmp(val)
	if d < 0 {
		return node
	}
	node.ch[d] = t._put(node.ch[d], val)
	if node.ch[d].priority > node.val {
		node = node.Rotate(d ^ 1)
	}
	return node
}

func (t *Treap) Del(val int) {
	t.root = t._del(t.root, val)
}

func (t *Treap) _del(node *Node, val int) *Node {
	if d := node.Cmp(val); d >= 0 {
		node.ch[d] = t._del(node.ch[d], val)
		return node
	}
	if node.ch[0] == nil {
		return node.ch[1]
	}
	if node.ch[1] == nil {
		return node.ch[0]
	}
	d := 0
	if node.ch[0].priority > node.ch[1].priority {
		d = 1
	}
	node.Rotate(d)
	node.ch[d] = t._del(node.ch[d], val)
	return node
}

func (t *Treap) lowerBound(val int) (lb *Node) {
	for o := t.root; o != nil; {
		if d := o.Cmp(val); d == 0 {
			lb = o
			o = o.ch[0]
		} else if d > 0 {
			o = o.ch[1]
		} else {
			return o
		}
	}
	return
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	set := &Treap{}
	for i, v := range nums {
		if lb := set.lowerBound(v - t); lb != nil && lb.val <= v + t {
			return true
		}
		set.Put(v)
		if i >= k {
			set.Del(nums[i - k])
		}
	}
	return false
}