/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/16 3:38 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _56_数组中只出现一次的两个数字

func FindNumsAppearOnce( array []int ) []int {
	if array == nil || len(array) == 0 {
		return nil
	}
	res := 0
	for _, v := range array {
		res ^= v
	}
	compare := 1
	for (compare & res) == 0 {
		compare <<= 1
	}
	m, n := 0, 0
	for _, v := range array {
		if (v & compare) == 0 {
			m ^= v
		} else {
			n ^= v
		}
	}
	if m <= n {
		return []int{m, n}
	} else {
		return []int{n, m}
	}
}