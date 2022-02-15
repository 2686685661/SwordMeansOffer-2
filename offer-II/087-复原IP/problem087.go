/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/8 4:37 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _87_复原IP

import (
	"strconv"
	"strings"
)

/**
回溯法
 */

func restoreIpAddresses(s string) []string {
	if s == "" {
		return nil
	}

	var res []string
	var cur []string

	var dfs func(idx int)
	dfs = func(idx int) {
		if len(cur) >= 4 && idx < len(s) {
			return
		}
		if len(cur) == 4 && idx == len(s) {
			res = append(res, strings.Join(append([]string{}, cur...), "."))
			return
		}
		for i := idx + 1; i <= len(s); i++ {
			if ii, _ := strconv.Atoi(s[idx:i]); ii <= 255 {
				// 剪枝
				if len(s[idx:i]) <= 1 || len(s[idx:i]) > 1 && s[idx:i][0] != '0' {
					cur = append(cur, s[idx:i])
					dfs(i)
					cur = cur[:len(cur)-1]
				}
			}
		}

	}
	dfs(0)
	return res
}