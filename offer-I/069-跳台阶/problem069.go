/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/17 8:48 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _69_跳台阶



// 递归
func jumpFloor( number int ) int {
	if number <= 2 {
		return number
	}
	return jumpFloor(number - 1) + jumpFloor(number - 2)
}

// 循环
func jumpFloor2( number int ) int {
	if number <= 2 {
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