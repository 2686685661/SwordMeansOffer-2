/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/25 12:22 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _33_二叉搜索树的后序遍历序列

import (
	"fmt"
	"testing"
)

func TestVerifySquenceOfBST(t *testing.T) {
	ta := []int{4,8,6,12,16,14,10}

	fmt.Println(VerifySquenceOfBST(ta))
}