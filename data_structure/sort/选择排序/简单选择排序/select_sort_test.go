/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 3:55 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 简单选择排序

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	elem := []int{0, 5, 3, 4, 6, 2}
	atual := SelectSort(elem)
	atual = atual[1:]
	fmt.Println(atual)
}
