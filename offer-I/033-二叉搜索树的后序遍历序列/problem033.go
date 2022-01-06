/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/25 12:09 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _33_二叉搜索树的后序遍历序列

func VerifySquenceOfBST( sequence []int ) bool {
	if sequence == nil || len(sequence) <= 0 {
		return false
	}
	r := sequence[len(sequence) - 1]
	var i, j int
	for i = 0; i < len(sequence) - 1; i++ {
		if sequence[i] > r {
			break
		}
	}

	for j = i; j < len(sequence) - 1; j++ {
		if sequence[j] < r {
			return false
		}
	}
	left := true
	if i > 0 {
		left = VerifySquenceOfBST(sequence[:i])
	}
	right := true
	if i < len(sequence) - 1 {
		right = VerifySquenceOfBST(sequence[i:len(sequence)-1])
	}
	return left && right
}