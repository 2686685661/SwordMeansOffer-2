/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/1 7:44 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _41_数据流中的中位数

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	a := []int{5,2,3,4,1,6,7,0,8}
	for _v := range a {
		Insert(_v)
		fmt.Println(GetMedian())
	}
}