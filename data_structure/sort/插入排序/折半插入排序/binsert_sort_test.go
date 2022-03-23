/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 1:27 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 折半插入排序

import (
	"fmt"
	"testing"
)

func TestBInsertSort(t *testing.T) {
	elem := []int{0, 5, 3, 4, 6, 2}
	atual := BInsertSort(elem)
	atual = atual[1:]
	fmt.Println(atual)
}