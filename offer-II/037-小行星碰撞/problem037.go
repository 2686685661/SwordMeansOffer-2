/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/31 12:04 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _37_小行星碰撞

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

/**
任意一个条件发生，则不会碰撞，可以将当前元素入栈：
1. 栈为空
2. 栈不为空，且栈顶元素与当前行星同向
3. 栈不为空，栈顶元素向左，当前元素向右，不会碰撞

当前元素向左，栈顶元素向右，会发生碰撞
1. 当前元素绝对值小于栈顶元素，当前元素爆炸，无需入栈
2. 当前元素绝对值等于栈顶元素，双方都会爆炸，弹出栈顶元素，当前元素无需入栈
3. 当前元素绝对值大于栈顶元素，栈顶爆炸，栈顶出栈，且当前元素继续向左移动，可能引发栈中元素持续爆炸,出栈后：
	3.1 栈顶元素比当前元素小且反方向，持续爆炸
	3.2 栈空了，直接入栈
	3.3 栈顶元素和当前元素同向，直接入栈
	3.4 栈顶元素和当前元素反向且绝对值相同，同时爆炸，栈顶出栈
	3.5 栈顶元素比当前元素大且反向，当前元素爆炸，无需进栈
 */
func asteroidCollision(asteroids []int) []int {
	if asteroids == nil || len(asteroids) < 2 {
		return asteroids
	}
	stack := new(Stack)
	intAbs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for _, star := range asteroids {
		if stack.IsEmpty() || stack.Top() * star > 0 || (stack.Top() < 0 && star > 0) {
			stack.Push(star)
			continue
		}
		if stack.Top() > 0 && star < 0 {
			if intAbs(star) < stack.Top() {
				continue
			}
			if intAbs(star) == stack.Top() {
				stack.Pop()
				continue
			}

			if intAbs(star) > stack.Top() {
				for !stack.IsEmpty() && intAbs(star) > stack.Top() && star * stack.Top() <0 {
					stack.Pop()
				}
				if stack.IsEmpty() || stack.Top() * star > 0 {
					stack.Push(star)
				} else if star + stack.Top() == 0 {
					stack.Pop()
				}
			}
		}
	}

	// 栈是反着的
	var res = make([]int, stack.Len())
	for i := len(res) - 1; i >= 0; i-- {
		res[i] = stack.Pop()
	}
	return res

}
