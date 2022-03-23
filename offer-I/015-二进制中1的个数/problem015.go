/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/19 5:40 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _15_二进制中1的个数


func NumberOf1( n int ) int {

	cnt := 0
	if n < 0 {
		n &= 0x7FFFFFFF
		cnt++
	}
	for n != 0 {
		cnt += n & 1
		n >>= 1
	}
	return cnt
}