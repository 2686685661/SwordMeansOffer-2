/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/18 3:16 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _09_用两个栈实现队列

var stack1 [] int
var stack2 [] int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int{
	if len(stack1) == 0 {
		return 0
	}
	n := stack1[0]
	stack1 = stack1[1:]
	return n
}