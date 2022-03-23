/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/13 3:20 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package treap

import (
	"fmt"
	"testing"
)

func TestTreap_Put(t *testing.T) {
	set := &treap{}
	for i := 0; i < 10; i++ {
		set.Put(i)
	}

	inorder(set.root)
	fmt.Printf("\n")

	set.Del(7)
	inorder(set.root)
	fmt.Printf("\n")
}