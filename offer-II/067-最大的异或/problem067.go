/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/24 3:33 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _67_最大的异或


// 暴力-超出限制
func findMaximumXOR(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] ^ nums[j] > res {
				res = nums[i] ^ nums[j]
			}
		}
	}
	return res
}




// 字典树

type Trie struct {
	children [2]*Trie
	isEnd bool
}

const (
	numLen = 30
)
func (t *Trie) Insert(num int) {
	node := t
	for i := numLen; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[bit] == nil {
			node.children[bit] = &Trie{}
		}
		node = node.children[bit]
	}
	node.isEnd = true
}

func (t *Trie) Check(num int) int {
	node := t
	var res int
	for i := numLen; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[bit ^ 1] != nil {
			res |= 1 << i
			node = node.children[bit ^ 1]
		} else {
			node = node.children[bit]
		}
	}
	return res
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func findMaximumXOR2(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	t := Trie{}
	t.Insert(nums[0])
	var res int
	for i := 1; i < len(nums); i++ {
		res = Max(res, t.Check(nums[i]))
		t.Insert(nums[i])
	}
	return res
}