/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/12 1:24 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _56_二叉搜索树中两个节点之和

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


// dfs + 哈希
func findTarget(root *TreeNode, k int) bool {

	m := make(map[int]int)
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if _, ok := m[k - node.Val]; ok {
			return true
		}
		m[node.Val] = 1
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}




// 单调栈实现
// ls 中序遍历，栈顶元素相当于左指针
// rs 反向中序，栈顶元素相当于右指针
func findTarget2(root *TreeNode, k int) bool {
	var (
		ls = new(Stack) // 中序遍历，单调递减栈
		rs = new(Stack) // 反向中序，单调递增栈
		tmp = root
	)

	lscall := func() {
		for tmp != nil {
			ls.Push(tmp)
			tmp = tmp.Left
		}
	}
	rscall := func() {
		for tmp != nil {
			rs.Push(tmp)
			tmp = tmp.Right
		}
	}

	lscall()
	tmp = root
	rscall()

	for !ls.IsEmpty() && !rs.IsEmpty() {
		l, r := ls.Top(), rs.Top()
		if l == r {
			return false
		}
		sum := l.Val + r.Val
		if sum > k {
			tmp = rs.Pop().Left
			rscall()
		} else if sum < k {
			tmp = ls.Pop().Right
			lscall()
		} else {
			return true
		}
	}
	return false
}

type Stack []*TreeNode
func (s Stack) Len() int {
	return len(s)
}
func (s Stack) Cap() int {
	return cap(s)
}

func (s *Stack) Push(n *TreeNode) {
	*s = append(*s, n)
}

func (s Stack) Top() *TreeNode {
	if len(s) == 0 {
		return nil
	}
	return s[len(s) - 1]
}

func (s *Stack) Pop() *TreeNode {
	ts := *s
	if len(ts) == 0 {
		return nil
	}
	v := ts[len(ts) - 1]
	*s = ts[:len(ts) - 1]
	return v
}

func (s Stack) IsEmpty() bool  {
	return len(s) == 0
}

