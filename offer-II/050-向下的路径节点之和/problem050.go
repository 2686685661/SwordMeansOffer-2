/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/10 12:17 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _50_向下的路径节点之和

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

/*
 dfs 时间复杂度O(N^2)
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := dfs(root, targetSum)

	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func dfs(node *TreeNode, target int) (ans int) {
	if node == nil {
		return 0
	}
	target -= node.Val
	if target == 0 {
		ans++
	}
	ans += dfs(node.Left, target)
	ans += dfs(node.Right, target)
	return ans
}

/**
前缀和
时间复杂度O(N)
https://leetcode-cn.com/problems/6eUYwP/solution/xiang-xia-de-lu-jing-jie-dian-zhi-he-by-a1iyy/
 */
func pathSum2(root *TreeNode, targetSum int) int {

	ans := 0
	preSum := make(map[int64]int)
	preSum[0] = 1
	var dfs func(node *TreeNode, curr int64)
	dfs = func(node *TreeNode, curr int64) {
		if node == nil {
			return
		}
		curr += int64(node.Val)
		ans += preSum[curr - int64(targetSum)]
		preSum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]--
	}
	dfs(root, 0)
	return ans

}


