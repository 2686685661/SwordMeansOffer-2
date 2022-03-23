/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/18 3:42 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _10_斐波那契数列

// 递归
func Fibonacci( n int ) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fibonacci(n - 1) + Fibonacci(n - 2)
}

// 循环 时间复杂度O(n)，空间复杂度O(n)
func Fibonacci_2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	arr := make([]int, n+1)
	arr[1],arr[2] = 1, 1

	for i := 3; i <= n; i++ {
		arr[i] = arr[i - 1] + arr[i - 2]
	}
	return arr[n]
}

// 循环 时间复杂度O(n)，空间复杂度O(1)
func Fibonacci_3(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	o, t := 1, 1

	var fn int

	for i := 3; i <= n; i++ {
		fn = o + t
		o, t = t, fn
	}
	return fn
}

