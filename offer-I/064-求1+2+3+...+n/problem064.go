/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/17 1:59 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _64_求1_2_3_____n

import "math"

func Sum_Solution( n int ) int {
	return (int(math.Pow(float64(n), 2)) + n) >> 2
}