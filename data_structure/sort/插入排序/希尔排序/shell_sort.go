/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 2:06 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 希尔排序

//https://blog.csdn.net/lishanleilixin/article/details/88585863

// 希尔排序
// 增量d的取法：
// d = math.Floor(len(elem)) / 3
// d = math.Floor(d) / 3
func ShellSort(elem []int) []int {
	j, d := 0, len(elem)
	for {
		d  = d / 3 + 1
		for i := d+1; i < len(elem);i++ {
			if elem[i] < elem[i - d] {
				elem[0] = elem[i]
				for j = i - d; j > 0 && elem[0] < elem[j]; j -= d {
					elem[j + d] = elem[j]
				}
				elem[j + d] = elem[0]
			}
		}

		if d <= 1 {
			break
		}
	}
	return elem
}