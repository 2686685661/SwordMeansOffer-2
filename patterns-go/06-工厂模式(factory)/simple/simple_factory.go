/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 10:32 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package simple

import "fmt"


/**
简单工厂
 */


// 女孩
type girl interface {
	weight()
}


// 胖女孩
type fatGirl struct {}

func (*fatGirl) weight() {
	fmt.Println("80kg")
}

// 瘦女孩
type thinGirl struct {}

func (*thinGirl) weight() {
	fmt.Println("50kg")
}


// 女孩工厂
type girlFactory struct {}

func (*girlFactory) createGirl(like string) girl {
	if like == "fat" {
		return &fatGirl{}
	}
	if like == "thin" {
		return &thinGirl{}
	}
	return nil
}