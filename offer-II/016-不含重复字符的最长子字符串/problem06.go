/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/27 9:18 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _16_不含重复字符的最长子字符串


// 双指针法
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	var max int
	m := make(map[byte]struct{})

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
		max = Max(max, r - l + 1)
	}
	return max
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}