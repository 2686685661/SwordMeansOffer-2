/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/27 11:14 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _75_数组相对排序

import (
	"fmt"
	"testing"
)

func TestRelativeSortArray(t *testing.T) {
	arr1 := []int{2,3,1,3,2,4,6,7,9,2,19}
	arr2 := []int{2,1,4,3,9,6}
	fmt.Println(relativeSortArray(arr1, arr2))
}
