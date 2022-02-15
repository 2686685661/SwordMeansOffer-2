/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/11 2:00 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package add


var datas []string

func Add(str string) int {
	data := []byte(str)
	datas = append(datas, string(data))
	return len(datas)
}