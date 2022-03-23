/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/23 11:11 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _05_单词长度的最大乘积

func maxProduct(words []string) int {

	if words == nil || len(words) < 1 {
		return 0
	}

	bitRes := make([]int, len(words))

	for i := 0; i < len(bitRes); i++ {
		for _, c := range words[i] {
			bitRes[i] = bitRes[i] | (1 << (c - 'a'))
		}
	}

	max := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if bitRes[i] & bitRes[j] == 0 {
				max = Max(max, len(words[i]) * len(words[j]))
			}
		}
	}

	return max

}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}