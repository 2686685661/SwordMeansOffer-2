/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/17 9:12 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _71_跳台阶扩展问题


func jumpFloorII( number int ) int {
	if number <= 2 {
		return number
	}

	return 2 * jumpFloorII(number - 1)
}

func jumpFloorII2( number int ) int {
	if number <= 2 {
		return number
	}

	front := 1
	back := 0
	for i := 2; i <= number; i++ {
		back = 2 * front
		front = back
	}
	return back
}