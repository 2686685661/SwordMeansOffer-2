/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/22 8:17 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 异或：相同为0，不同为1
func main() {
	//fmt.Println(4 ^ 0)

	//maxProduct([]string{"ab", "cd"})

	//nums := []int{-1,0,1,2,-1,-4}
	//sort.Ints(nums)
	//fmt.Println(nums)

	//m := make(map[int]int)
	//m[1] = 1
	//m[2] = 2
	//m[3] = 3
	//if _,ok := m[3]; ok {
	//	fmt.Println("fuck")
	//}
	//for k, _ := range m {
	//	fmt.Println(k)
	//	//break
	//}

	//fmt.Println(1 ^ 1)
	//set := &treap{}
	//for i := 0; i < 10; i++ {
	//	set.Put(i)
	//}
	//
	//
	//inorder(set.root)
	//fmt.Printf("\n")
	//
	//set.Del(3)
	//inorder(set.root)
	arr := []string{"me", "time"}
	sort.Strings(arr)
	fmt.Println(arr)


}



func maxProduct(words []string) int {

	if words == nil || len(words) < 1 {
		return 0
	}

	bitRes := make([]int, len(words))

	for i := 0; i < len(bitRes); i++ {
		for _, c := range words[i] {
			bitRes[i] = bitRes[i] | (1 << (c - 'a'))
		}
	}

	max := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if bitRes[i] & bitRes[j] == 0 {
				max = Max(max, len(words[i]) * len(words[j]))
			}
		}
	}

	return max

}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}


// 2，2，2，3
// 010
// 010
// 010
// 011
// 0 4 1




func inorder(n *node) {
	if n == nil {
		return
	}
	inorder(n.ch[0])
	fmt.Printf("%d\t", n.val)
	inorder(n.ch[1])
}

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






















const highBit = 30

type trie struct {
	left, right *trie
}

func (t *trie) add(num int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.left == nil {
				cur.left = &trie{}
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &trie{}
			}
			cur = cur.right
		}
	}
}

func (t *trie) check(num int) (x int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			// a_i 的第 k 个二进制位为 0，应当往表示 1 的子节点 right 走
			if cur.right != nil {
				cur = cur.right
				x = x*2 + 1
			} else {
				cur = cur.left
				x = x * 2
			}
		} else {
			// a_i 的第 k 个二进制位为 1，应当往表示 0 的子节点 left 走
			if cur.left != nil {
				cur = cur.left
				x = x*2 + 1
			} else {
				cur = cur.right
				x = x * 2
			}
		}
	}
	return
}

func findMaximumXOR(nums []int) (x int) {
	root := &trie{}
	for i := 1; i < len(nums); i++ {
		// 将 nums[i-1] 放入字典树，此时 nums[0 .. i-1] 都在字典树中
		root.add(nums[i-1])
		// 将 nums[i] 看作 ai，找出最大的 x 更新答案
		x = max(x, root.check(nums[i]))
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}