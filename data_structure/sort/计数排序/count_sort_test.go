/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/10 1:30 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 计数排序

import (
	"fmt"
	"testing"
)

// 计数排序
// 计数排序对每一个输入的元素elem[i]，确定小于 elem[i] 的元素个数。也就是说需要有一个大于待排序数组最大值长度的数组，然后直接把 elem[i] 放到它输出数组中的位置上对于每个待排序的元素数据都有自己用来记录自己的个数。
func TestCountSort(t *testing.T) {
	elem := []int{5, 3, 4, 6, 2, 1, 8, 2}
	atual := CountSort(elem)
	fmt.Println(atual)
}