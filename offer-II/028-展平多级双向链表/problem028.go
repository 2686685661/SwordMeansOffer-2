/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/30 11:44 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _28_展平多级双向链表

type Node struct {
	Val int
	Prev *Node
	Next *Node
	Child *Node
}


/**
深度优先搜索
https://leetcode-cn.com/problems/Qv1Da2/solution/zhan-ping-duo-ji-shuang-xiang-lian-biao-x5ugr/
 */
func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	dfs(root)
	return root

}

func dfs(node *Node) *Node {
	cur := node
	var last *Node
	for cur != nil {
		next := cur.Next
		if cur.Child != nil {
			chls := dfs(cur.Child)
			cur.Next = cur.Child
			cur.Child.Prev = cur

			if next != nil {
				chls.Next = next
				next.Prev = chls
			}
			cur.Child = nil
			last = chls
		} else {
			last = cur
		}
		cur = next
	}
	return last
}