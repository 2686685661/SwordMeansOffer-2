/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/21 1:48 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _29_顺时针打印矩阵

import (
	"fmt"
	"testing"
)

func TestPrintMatrix(t *testing.T) {
	a := [][]int{
		{1,2},
		{3, 4},
	}
	fmt.Println(PrintMatrix(a))
}