/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/10 2:18 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 基数排序

import (
	"fmt"
	"testing"
)

func TestRadixSort(t *testing.T) {
	elem := []int{5, 3, 4, 6, 2, 1, 8, 2}
	atual := RadixSort(elem)
	fmt.Println(atual)
}