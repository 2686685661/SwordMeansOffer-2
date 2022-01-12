/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/10 11:31 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _48_序列化与反序列化二叉树

import (
	"bytes"
	"strconv"
	"strings"
)

//dfs

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	if root == nil {
		return ""
	}
	return serialize(root, bytes.NewBufferString("")).String()

}

func serialize(root *TreeNode, bf *bytes.Buffer) *bytes.Buffer {
	if root == nil {
		bf.WriteString("null,")
		return bf
	}
	bf.WriteString(strconv.Itoa(root.Val) + ",")
	serialize(root.Left, bf)
	serialize(root.Right, bf)
	return bf
}


// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	strs := strings.Split(data, ",")
	return deserialize(strs)

}

var idx = 0
func deserialize(strs []string) *TreeNode {
	if idx >= len(strs) || strs[idx] == "null" || strs[idx] == "" {
		idx++
		return nil
	}
	n, _ := strconv.Atoi(strs[idx])
	node := new(TreeNode)
	node.Val = n
	idx++
	node.Left = deserialize(strs)
	node.Right = deserialize(strs)
	return node
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */