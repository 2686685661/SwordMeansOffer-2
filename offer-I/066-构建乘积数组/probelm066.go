/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/17 5:39 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _66_构建乘积数组


func multiply( A []int ) []int {

	if A == nil || len(A) == 0 {
		return nil
	}

	b := make([]int, len(A))

	b[0] = 1

	for i := 1; i < len(A); i++ {
		b[i] = b[i - 1] * A[i - 1]
	}

	tmp := 1

	for i := len(A) - 2; i >= 0; i-- {
		tmp *= A[i + 1]
		b[i] *= tmp
	}
	return b
}