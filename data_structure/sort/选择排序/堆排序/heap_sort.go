/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 4:18 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 堆排序

// 堆排序
// 大顶堆
func MaxHeapSort(elem []int) []int {

	// 初建大顶堆
	for i := len(elem) / 2; i >= 0; i-- {
		elem = maxHeapAdjust(elem, i, len(elem) - 1)
	}

	for i := len(elem) - 1; i >= 0; i-- {
		elem[0], elem[i] = elem[i], elem[0]
		elem = maxHeapAdjust(elem, 0, i - 1)
	}
	return elem
}

func maxHeapAdjust(elem []int, s, m int) []int {
	j, tmp := 0, elem[s]
	for j = 2 * s; j <= m; j *= 2 {
		if j < m && elem[j] < elem[j + 1] {
			j += 1
		}
		if tmp >= elem[j] {
			break
		}
		elem[s] = elem[j]
		s = j
	}
	elem[s] =tmp
	return elem
}




