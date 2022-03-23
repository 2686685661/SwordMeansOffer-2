/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/5 2:13 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _49_丑数

func GetUglyNumber_Solution( index int ) int {
	if index < 7 {
		return index
	}

	res := make([]int, index)
	res[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < index; i++ {
		res[i] = min(res[p2] * 2, res[p3] * 3, res[p5] * 5)
		if res[i] == (res[p2] * 2) {
			p2++
		}
		if res[i] == (res[p3] * 3) {
			p3++
		}
		if res[i] == (res[p5] * 5) {
			p5++
		}
	}
	return res[index - 1]
}

func min(x, y, c int) int {
	if x < y {
		if x < c {
			return x
		} else {
			return c
		}
	} else {
		if y < c {
			return y
		} else {
			return c
		}
	}
}