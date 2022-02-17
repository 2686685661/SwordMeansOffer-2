/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 9:04 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _4_单例模式_singleton_

import "sync"

type singleton struct {}

var (
	ins *singleton
	once sync.Once
)

func GetIns() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}