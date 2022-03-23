/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/11 2:01 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package add

import "testing"

func TestAdd(t *testing.T) {
	_ = Add("go-programming-tour-book")
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add("go-programming-tour-book")
	}
}