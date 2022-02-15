/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/9 3:54 下午, by lishanlei, create
*/

/*
问题描述
请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。 给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。

解题思路
这里有几个重点，第一个是ASCII字符，ASCII字符字符一共有256个，其中128个是常用字符，可以在键盘上输入。128之后的是键盘上无法找到的。
然后是全部不同，也就是字符串中的字符没有重复的，再次，不准使用额外的储存结构，且字符串小于等于3000。
如果允许其他额外储存结构，这个题目很好做。如果不允许的话，可以使用golang内置的方式实现。

源码解析

两种方法都可以实现这个算法。

第一个方法使用的是golang内置方法strings.Count,可以用来判断在一个字符串中包含的另外一个字符串的数量。

第二个方法使用的是golang内置方法strings.Index和strings.LastIndex，用来判断指定字符串在另外一个字符串的索引未知，分别是第一次发现位置和最后发现位置。
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	isUniqueString := func(s string) bool {
		if strings.Count(s, "") > 3000 {
			return false
		}
		for _, v := range s {
			if v > 127 {
				return false
			}
			if strings.Count(s, string(s)) > 1 {
				return false
			}
		}
		return true
	}
	fmt.Println("isUniqueString: " , isUniqueString("abcd"))


	isUniqueString2 := func(s string) bool {
		if strings.Count(s, "") > 3000 {
			return false
		}
		for k, v := range s {
			if v > 127 {
				return false
			}
			if strings.Index(s, string(v)) != k {
				return false
			}
		}
		return true
	}
	fmt.Println("isUniqueString2: " , isUniqueString2("abcdd"))


}