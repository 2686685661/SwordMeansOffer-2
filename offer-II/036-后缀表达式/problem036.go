/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/31 11:16 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _36_后缀表达式

import (
	"strconv"
)

type Stack []int
func (s Stack) Len() int {
	return len(s)
}
func (s Stack) Cap() int {
	return cap(s)
}

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

func (s Stack) Top() int {
	if len(s) == 0 {
		return 0
	}
	return s[len(s) - 1]
}

func (s *Stack) Pop() int {
	ts := *s
	if len(ts) == 0 {
		return 0
	}
	v := ts[len(ts) - 1]
	*s = ts[:len(ts) - 1]
	return v
}

func (s Stack) IsEmpty() bool  {
	return len(s) == 0
}


func evalRPN(tokens []string) int {
	if tokens == nil || len(tokens) < 1 {
		return 0
	}


	stack := new(Stack)

	for _, v := range tokens {
		if n, ok := IsNum(v); ok {
			stack.Push(n)
		} else {
			if stack.Len() < 2 {
				return 0
			}
			n1, n2 := stack.Pop(), stack.Pop()
			if v == "+" {
				stack.Push(n1 + n2)
			} else if v == "-" {
				stack.Push(n2 - n1)
			} else if v == "*" {
				stack.Push(n2 * n1)
			} else if v == "/" {
				stack.Push(n2 / n1)
			}
		}
	}
	if stack.Len() != 0 {
		return stack.Top()
	}
	return 0

}

func IsNum(s string) (int, bool) {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0, false
	}
	return int(i), true
}