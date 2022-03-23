/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/28 5:22 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _81_允许重复选择元素的组合

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
	fmt.Println(combinationSum([]int{2,3,5}, 8))
}