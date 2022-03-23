/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/18 11:28 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _70_矩形覆盖


// 斐波那契数列
//递归
func rectCover( number int ) int {
	if number < 3 {
		return number
	}

	return rectCover(number -1) + rectCover(number -2)
}

//循环
func rectCover2( number int ) int {
	if number < 3 {
		return number
	}
	f, s, t := 1, 2, 0
	for i := 3; i <= number; i++ {
		t = f + s
		f = s
		s = t
	}

	return t
}