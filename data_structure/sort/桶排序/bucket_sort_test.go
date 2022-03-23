/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 7:43 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 桶排序

import (
	"fmt"
	"testing"
)

func TestBucketSort(t *testing.T) {
	elem := []int{5, 3, 4, 6, 2, 1, 8, 2}
	atual := BucketSort(elem)
	fmt.Println(atual)
}