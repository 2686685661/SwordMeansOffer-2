/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 2:51 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _2_装饰器模式_decorator_

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	laowang := &laowang{}

	jacket := &Jacket{}
	jacket.person = laowang
	jacket.show()

	hat := &Hat{}
	hat.person = jacket
	hat.show()

	fmt.Println("cost: ", hat.cost())


}