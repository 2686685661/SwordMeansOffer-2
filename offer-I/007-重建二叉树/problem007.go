/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/18 2:20 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _07_重建二叉树

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/**
 * 根据前序、中序遍历结果，重建二叉树
 *
 * @param pre int整型一维数组
 * @param vin int整型一维数组
 * @return TreeNode类
 */
func reConstructBinaryTree( pre []int ,  vin []int ) *TreeNode {
	if len(pre) == 0 || len(vin) == 0 {
		return nil
	}
	return reConstructBinaryTreeCore(pre, vin, 0, len(pre)  -1, 0, len(vin) - 1)
}

func reConstructBinaryTreeCore(pre, vin []int, pres, pree, vins, vine int) *TreeNode {

	tree := new(TreeNode)
	tree.Val = pre[pres]
	if pres == pree && vine == vins {
		return tree
	}

	var sub int
	for i, v := range vin {
		if pre[pres] == v {
			sub = i
			break
		}
	}
	ltl := sub - vins
	rtl := vine - sub

	if ltl > 0 {
		tree.Left = reConstructBinaryTreeCore(pre, vin, pres + 1, pres + ltl, vins, sub - 1)
	}

	if rtl > 0 {
		tree.Right = reConstructBinaryTreeCore(pre, vin, pres + ltl + 1, pree, sub + 1, vine)
	}
	return tree
}