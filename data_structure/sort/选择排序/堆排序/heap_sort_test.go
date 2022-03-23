/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 4:30 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 堆排序

import (
	"fmt"
	"testing"
)

func TestMaxHeapSort(t *testing.T) {
	expect := []int{50, 10, 90, 30, 70, 40, 80, 60, 20}
	actual := MaxHeapSort(expect)
	fmt.Println(actual)
}