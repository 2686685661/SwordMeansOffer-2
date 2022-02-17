/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 11:04 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package simple

import "testing"

func TestSimple(t *testing.T) {

	// 创建工厂
	girlFactory := new(girlFactory)

	// 传递你喜欢类型的姑娘即可
	girl := girlFactory.createGirl("fat")
	girl.weight()

	girl = girlFactory.createGirl("thin")
	girl.weight()
}
