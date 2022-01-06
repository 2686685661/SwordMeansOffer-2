/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/5 12:12 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _48_最长不含重复字符的子字符串


func lengthOfLongestSubstring( s string ) int {
	if len(s) <= 1 {
		return len(s)
	}
	var max int
	for i := 0;i < len(s); i++ {
		m := make(map[byte]struct{})
		m[s[i]] = struct{}{}
		for j := i + 1; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {
				break
			}
			m[s[j]] = struct{}{}
		}
		if len(m) > max {
			max = len(m)
		}
	}
	return max
}

func lengthOfLongestSubstring2( s string ) int {
	if len(s) <= 1 {
		return len(s)
	}

	m := make(map[byte]struct{})
	res := 0
	for l, r := 0, 0; r < len(s); r++ {
		for {
			if _, ok := m[s[r]]; ok {
				delete(m, s[l])
				l++
			} else {
				break
			}
		}
		m[s[r]] = struct{}{}
		res = max(res, r - l + 1)
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}