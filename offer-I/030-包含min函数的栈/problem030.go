/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/21 2:11 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _30_包含min函数的栈

import "math"

var (
	stack []int
	minStack []int
	min int = math.MaxInt64
)
func Push(node int) {
	// write code here
	if node < min {
		minStack = append(minStack, node)
		min = minStack[len(minStack) - 1]
	} else {
		minStack = append(minStack, min)
	}
	stack = append(stack, node)
}

func Pop() {
	// write code here
	if len(stack) != 0 {
		stack = stack[:len(stack)-1]
	}
	if len(minStack) != 0 {
		minStack = minStack[:len(minStack)-1]
		min = minStack[len(minStack) - 1]
	}
}

func Top() int {
	// write code here
	if len(stack) == 0 {
		return 0
	}
	return stack[len(stack) - 1]
}

func Min() int {
	// write code here
	return min
}

