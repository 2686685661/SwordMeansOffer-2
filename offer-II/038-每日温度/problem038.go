/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/31 2:13 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _38_每日温度


// 暴力
func dailyTemperatures(temperatures []int) []int {
	if temperatures == nil || len(temperatures) < 1 {
		return nil
	}

	var res = make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				res[i] = j - i
				break
			}
		}
	}
	return res
}

// 单调栈
func dailyTemperatures2(temperatures []int) []int {

	stack, res := make([]int, 0), make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack) - 1]] {
			fi := stack[len(stack) - 1]
			stack = stack[:len(stack)-1]
			res[fi] = i - fi
		}
		stack = append(stack, i)
	}
	return res

}