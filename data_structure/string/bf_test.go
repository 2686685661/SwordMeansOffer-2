/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/7/21 9:33 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package string

import (
	"fmt"
	"testing"
)

func TestBf(t *testing.T) {
	s := "ababcabcacbababcac"
	ta := "abcac"
	fmt.Println(Bf(s, ta, 0))
}
