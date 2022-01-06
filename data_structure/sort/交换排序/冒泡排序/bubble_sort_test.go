/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 2:32 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 冒泡排序

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	elem := []int{0, 5, 3, 4, 6, 2}
	atual := BubbleSort(elem)
	atual = atual[1:]
	fmt.Println(atual)
}

func TestBubbleSort2(t *testing.T) {
	elem := []int{0, 5, 3, 4, 6, 2}
	atual := BubbleSort(elem)
	atual = atual[1:]
	fmt.Println(atual)
}