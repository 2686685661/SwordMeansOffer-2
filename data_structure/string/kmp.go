/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/7/21 9:28 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package string

// https://www.cnblogs.com/yjiyjige/p/3263858.html
// https://www.ruanyifeng.com/blog/2013/05/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm.html
// https://blog.csdn.net/v_july_v/article/details/7041827

func kmp(str, ps string) int {
	i, j := 0, 0
	next := getNext(ps)
	for i < len(str) && j < len(ps) {
		if j == -1 || str[i] == ps[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(ps) {
		return i - j
	}
	return -1

}

func getNext(ps string) []int {
	chPs := []byte(ps)
	next := make([]int, len(chPs))
	next[0] = -1
	j, k := 0, -1
	for k < len(chPs) - 1 {
		if k == -1|| chPs[j] == chPs[k] {
			j += 1
			k += 1
			next[j] = k
		} else {
			k = next[k]
		}
	}
	return next
}