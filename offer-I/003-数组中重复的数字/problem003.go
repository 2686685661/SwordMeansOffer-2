/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/7/13 1:54 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package main

import "fmt"

// 通过hash table
func duplicate003_1( numbers []int ) int {
	// write code here
	appear := make(map[int]bool)
	for _, v := range numbers {
		if _, ok := appear[v]; ok {
			return v
		}
		appear[v] = true
	}

	return len(numbers) + 1
}

// 交换法
func duplicate003_2( numbers []int ) int {
	// write code here

	for i := 0; i < len(numbers); i++ {
		if i == numbers[i] {
			continue
		}
		if numbers[i] == numbers[numbers[i]] {
			return numbers[i]
		} else {
			numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
			// 遍历完0位元素以及交换完数字后，如果0位元素仍不符合数组存放原则：numbers[i] = i，那么还要重新遍历0位元素
			i--
		}
	}

	return len(numbers) + 1
}

func main() {
	numbers := []int{2,3,1,0,2,5,3}
	fmt.Println(duplicate003_2(numbers))

}


