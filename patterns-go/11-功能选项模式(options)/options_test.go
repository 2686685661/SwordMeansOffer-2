/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/18 12:53 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _1_功能选项模式_options_

import (
	"fmt"
	"testing"
)

func TestOptions(t *testing.T) {
	msg := NewByOption(WithId(1), WithName("test1"), WithPhone(158))
	fmt.Printf("%+v\n", msg)

	msg2 := NewRequireIdByOption(2, WithAddress("localhost"), WithPhone(137), WithName("test2"))
	fmt.Printf("%+v\n", msg2)
}
