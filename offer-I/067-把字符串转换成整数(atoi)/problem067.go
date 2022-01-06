/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/17 6:42 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _67_把字符串转换成整数_atoi_

import (
	"math"
	"strings"
)

func StrToInt( s string ) int {


	s =strings.Trim(s, " ")
	if s == "" || len(s) == 0 {
		return 0
	}

	num := 0
	flag := 1
	ca := []byte(s)
	if ca[0] == '+' {
		ca = ca[1:]
	} else if ca[0] == '-' {
		ca =ca[1:]
		flag = -1
	}

	if len(ca) == 0 {
		return num
	}

	for _, v :=  range ca {
		if v >= '0' && v <= '9' {
			num = num * 10 + flag * int(v - '0')
			if flag == 1 && num > math.MaxInt32 {
				num = math.MaxInt32
				break
			} else if flag == -1 && num < -math.MaxInt32 {
				num = -math.MaxInt32 - 1
				break
			}
		} else {
			break
		}
	}

	return num

}