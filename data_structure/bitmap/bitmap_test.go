/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/3/18 12:07 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package bitmap

import (
	"fmt"
	"testing"
)

func TestBitMap(t *testing.T) {
	bm := NewBitMap(10)
	bm.Add(3)
	bm.Add(2)
	bm.Add(9)
	bm.Add(36)

	bm.Clear(9)
	fmt.Println(bm.Contain(9))

	fmt.Printf("%b, %b\n", bm.bits[0],bm.bits[1])
	fmt.Println(bm.Contain(3), bm.Contain(4))

}