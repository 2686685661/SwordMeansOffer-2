/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 9:10 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _4_单例模式_singleton_

import "testing"

func TestGoSingleton(t *testing.T) {
	ins := GetIns()
	t.Logf("%p", ins)

	in2 := GetIns()
	t.Logf("%p", in2)

}

func TestSingleton(t *testing.T) {
	instance1 := GetInsOther()
	//查看其内存地址
	t.Logf("%p", instance1)

	instance2 := GetInsOther()
	t.Logf("%p", instance2)
}