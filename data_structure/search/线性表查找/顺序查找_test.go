/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/6 4:04 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 线性表查找

import (
	"testing"
)

func TestOrderSearch(t *testing.T) {
	a := []int{1,2,3, 5, 4, 7, 6}
	k, e := 3, 2
	if aa := OrderSearch(a, k); aa != e {
		t.Errorf("expect:%d, actual:%d\n", e, aa)
	}
}