/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/19 5:49 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _15_二进制中1的个数

import (
	"fmt"
	"testing"
)

func TestNumberOf1(t *testing.T) {
	fmt.Println(NumberOf1(-2147483648))
	n := -2
	n &= 0x7FFFFFFF
	fmt.Println(n)
}
