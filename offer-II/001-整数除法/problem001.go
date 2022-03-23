/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/22 11:31 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _01_整数除法

import "math"

func divide(a int, b int) int {

	if b == 0 {
		return 0
	}
	if a == math.MinInt32 {
		if b == -1 {
			return math.MaxInt32
		}
		if b == 1 {
			return math.MinInt32
		}
	}

	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -sign
	}
	a, b = abs(a), abs(b)
	res := 0
	for a >= b {
		a -= b
		res++
	}

	return res * sign
}


func divide2(a int, b int) int {
	if b == 0 {
		return 0
	}
	if a == math.MinInt32 {
		if b == -1 {
			return math.MaxInt32
		}
		if b == 1 {
			return math.MinInt32
		}
	}

	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -sign
	}
	a, b = abs(a), abs(b)
	res := 0

	for i := 31; i >= 0; i-- {
		if a >> i >= b {
			a -= b << i
			res += 1 << i
		}
	}
	return res * sign

}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}