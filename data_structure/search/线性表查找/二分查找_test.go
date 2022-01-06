/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/6 4:20 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 线性表查找

import "testing"

func TestBinarySearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	k, e := 6, 5
	if aa := BinarySearch(a, k); aa != e {
		t.Errorf("expect:%d, actual:%d\n", e, aa)
	}
}

func TestBinarySearchRec(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	k, e := 6, 5
	if aa := BinarySearchRec(a, k,0, len(a)-1); aa != e {
		t.Errorf("expect:%d, actual:%d\n", e, aa)
	}
}