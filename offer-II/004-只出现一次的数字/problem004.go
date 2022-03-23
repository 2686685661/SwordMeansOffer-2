/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/22 8:00 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _04_只出现一次的数字


func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
