/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/16 7:30 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _59_滑动窗口的最大值

import (
	"fmt"
	"testing"
)

func TestMaxInWindows(t *testing.T) {

	arr := []int{2,3,4,2,6,2,5,1}
	size := 3
	fmt.Println(MaxInWindows(arr, size))
}