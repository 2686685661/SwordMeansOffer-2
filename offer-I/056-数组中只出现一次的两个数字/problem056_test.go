/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/16 3:46 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _56_数组中只出现一次的两个数字

import (
	"fmt"
	"testing"
)

func TestFindNumsAppearOnce(t *testing.T) {
	arr := []int{1,4,1,6}
	fmt.Println(FindNumsAppearOnce(arr))
}