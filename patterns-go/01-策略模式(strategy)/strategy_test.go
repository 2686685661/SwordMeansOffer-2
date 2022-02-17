/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 12:54 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _1_策略模式_strategy_

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	op := Operator{}
	op.setStrategy(&add{})
	fmt.Println("add:", op.doCall(1, 2))

	op.setStrategy(&reduce{})
	fmt.Println("reduce:", op.doCall(2, 1))
}