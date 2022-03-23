/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/27 6:28 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _14_字符串中的变位词

/**
全排列参考：https://www.cnblogs.com/sclczk/p/11024022.html
 */

/**
妈的，第一反应是全排列，我真傻b
 */
import "strings"
func checkInclusion(s1 string, s2 string) bool {

	if s1 == "" || s2 == "" || len(s1) > len(s2) {
		return false
	}
	sch := []byte(s1)
	var list []string
	list = fullPermutation(sch, 0, list)
	if len(list) == 0 {
		return false
	}

	for _, v := range list {
		if ok := strings.Contains(s2, v); ok {
			return ok
		}
	}
	return false
}

func fullPermutation(sch []byte, begin int, list []string) []string {
	if begin == len(sch) - 1 {
		list = append(list, string(sch))
	} else {
		m := make(map[byte]bool)
		for i := 0; i < len(sch); i++ {
			if _, ok := m[sch[i]]; !ok {
				m[sch[i]] = true
				sch[begin], sch[i] = sch[i], sch[begin]
				list = fullPermutation(sch, begin + 1, list)
				sch[begin], sch[i] = sch[i], sch[begin]
			}
		}
	}
	return list
}

/**
滑动窗口
 */
func checkInclusion2(s1 string, s2 string) bool {
	if s1 == "" || s2 == "" || len(s1) > len(s2) {
		return false
	}

	n, m := len(s1), len(s2)
	var cnt1, cnt2 [26]int

	for i, v := range s1 {
		cnt1[v - 'a']++
		cnt2[s2[i] - 'a']++
	}
	if cnt1 == cnt2 {
		return true
	}

	for i := n; i < m; i++ {
		cnt2[s2[i] - 'a']++
		cnt2[s2[i - n] - 'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}