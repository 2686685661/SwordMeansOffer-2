/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 1:12 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 直接插入排序

import (
	"fmt"
	"testing"
)

func TestInsetSort(t *testing.T) {
	elem := []int{0, 5, 3, 4, 6, 2}
	atual := InsetSort(elem)
	atual = atual[1:]
	fmt.Println(atual)

}