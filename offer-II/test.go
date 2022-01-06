/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/22 8:17 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package main

// 异或：相同为0，不同为1
func main() {
	//fmt.Println(4 ^ 0)

	//maxProduct([]string{"ab", "cd"})

	//nums := []int{-1,0,1,2,-1,-4}
	//sort.Ints(nums)
	//fmt.Println(nums)

	//m := make(map[int]int)
	//m[1] = 1
	//m[2] = 2
	//m[3] = 3
	//if _,ok := m[3]; ok {
	//	fmt.Println("fuck")
	//}
	//for k, _ := range m {
	//	fmt.Println(k)
	//	//break
	//}
}


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


// 2，2，2，3
// 010
// 010
// 010
// 011
// 0 4 1