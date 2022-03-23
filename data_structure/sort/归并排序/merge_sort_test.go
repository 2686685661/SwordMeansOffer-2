/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 5:37 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 归并排序

import (
	"fmt"
	"testing"
)

func TestMergeSortUp2Down(t *testing.T) {
	elem := []int{5, 3, 4, 6, 2, 1, 8, 2}
	atual := MergeSortUp2Down(elem, 0, len(elem) - 1)
	fmt.Println(atual)
}

func TestMergeSortDown2Up(t *testing.T) {
	elem := []int{5, 3, 4, 6, 2, 1, 8, 2}
	atual := MergeSortDown2Up(elem)
	fmt.Println(atual)
}