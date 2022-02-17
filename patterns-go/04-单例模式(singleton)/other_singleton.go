/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 9:08 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _4_单例模式_singleton_

import "sync"

var (
	ins2 *singleton
	lock sync.Mutex
)

func GetInsOther() *singleton {
	if ins2 == nil {
		lock.Lock()
		defer lock.Unlock()
		if ins2 == nil {
			ins2 = &singleton{}
		}
	}
	return ins2
}