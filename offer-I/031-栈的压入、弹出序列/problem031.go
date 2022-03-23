/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/25 11:26 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _31_栈的压入_弹出序列


func IsPopOrder( pushV []int ,  popV []int ) bool {

	if len(pushV) == 0 || len(popV) == 0 || len(pushV) != len(popV) {
		return false
	}

	var stack []int

	for i, j := 0, 0; i < len(pushV); i++ {
		stack = append(stack, pushV[i])
		for j < len(popV) && len(stack) > 0 && stack[len(stack) - 1] == popV[j] {
			stack = stack[:len(stack) - 1]
			j += 1
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}