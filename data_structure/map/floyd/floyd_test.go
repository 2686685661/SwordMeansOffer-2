/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/16 9:42 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package floyd

import (
	"fmt"
	"testing"
)

func TestShortestPath_Floyd(t *testing.T) {
	ShortestPath_Floyd()
	for v := 0; v < MAXVEX; v++ {
		for w := 0; w < MAXVEX; w++ {
			str := fmt.Sprintf("%d-->%d, weight=%d", v, w, ShortPathTable[v][w])
			k := Patharc[v][w]
			str = fmt.Sprintf("%s, path: %d", str, v)
			for k != w {
				str = fmt.Sprintf("%s->%d", str, k)
				k = Patharc[k][w]
			}
			str = fmt.Sprintf("%s->%d", str, w)
			fmt.Println(str)
		}
	}

}