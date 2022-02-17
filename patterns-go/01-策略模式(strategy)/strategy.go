/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 12:46 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _1_策略模式_strategy_


// 实现该接口，则为一个策略
type IStrategy interface {
	do(int, int) int
}

// 加 策略
type add struct {}

func (*add) do(a, b int) int {
	return a + b
}

// 减 策略
type reduce struct {}

func (*reduce) do(a, b int) int {
	return a - b
}

// 策略执行者
type Operator struct {
	strategy IStrategy
}

// 设置具体策略
func (op *Operator) setStrategy(strategy IStrategy) {
	op.strategy = strategy
}

// 执行策略
func (op *Operator) doCall(a, b int) int {
	return op.strategy.do(a, b)
}

