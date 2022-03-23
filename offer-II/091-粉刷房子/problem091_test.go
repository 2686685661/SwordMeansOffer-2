/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/9 11:29 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _91_粉刷房子

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	var c [][]int
	_ = json.Unmarshal([]byte("[[3,5,3],[6,17,6],[7,13,18],[9,10,18]]"), &c)
	fmt.Println(minCost(c))
}