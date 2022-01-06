/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/18 4:19 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _11_旋转数组的最小数字

import (
	"fmt"
	"testing"
)

func TestMinNumberInRotateArray(t *testing.T) {
	a := []int{2,2,2,1,2}
	i := MinNumberInRotateArray(a)
	fmt.Println(i)
}