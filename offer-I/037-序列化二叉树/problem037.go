/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/25 9:01 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _37_序列化二叉树

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func Serialize( root *TreeNode ) string {
	var str string
	if root == nil {
		str += "#"
		return str
	}
	str += strconv.Itoa(root.Val)
	str += Serialize(root.Left)
	str += Serialize(root.Right)
	return str
}

var i int = -1
func Deserialize( s string ) *TreeNode {
	fmt.Println(strings.Split(s, ","))
	i += 1
	if i >= len(s) {
		return nil
	}

	var node *TreeNode
	if s[i:i+1] != "#" {
		si, _ := strconv.Atoi(s[i:i+1])
		fmt.Println(s[i:i+1])
		node = &TreeNode{
			Val:   si,
			Left:  nil,
			Right: nil,
		}
		node.Left = Deserialize(s)
		node.Right = Deserialize(s)
	}
	return node
}