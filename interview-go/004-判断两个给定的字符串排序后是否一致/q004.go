/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/9 4:09 下午, by lishanlei, create
*/

/*
问题描述
给定两个字符串，请编写程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。 这里规定【大小写为不同字符】，且考虑字符串重点空格。给定一个string s1和一个string s2，请返回一个bool，代表两串是否重新排列后可相同。 保证两串的长度都小于等于5000。

解题思路
首先要保证字符串长度小于5000。之后只需要一次循环遍历s1中的字符在s2是否都存在即可。

源码解析
这里还是使用golang内置方法 strings.Count 来判断字符是否一致。
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	isRegroup := func(s1, s2 string) bool {
		l1, l2 := len(s1), len(s2)
		if l1 > 5000 || l2 > 5000 || l1 != l2 {
			return false
		}

		for _, v := range s1 {
			if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
				return false
			}
		}
		return true
	}

	fmt.Println("isRegroup: ", isRegroup("aabc", "bcaa"))
}
