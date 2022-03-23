/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 3:03 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 快速排序

import (
	"fmt"
	"testing"
)

func TestQSort(t *testing.T) {
	elem := []int{0, 5, 3, 4, 6, 2}
	atual := QSort(elem, 1, len(elem) - 1)
	atual = atual[1:]
	fmt.Println(atual)
}