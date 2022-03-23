/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/22 11:45 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _01_整数除法

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	a,b := 17, 3

	for i := 31; i >= 0; i-- {
		fmt.Println(a >> i, b<<i)
	}

}