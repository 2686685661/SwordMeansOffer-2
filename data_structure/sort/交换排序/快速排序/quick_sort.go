/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 2:52 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 快速排序


// 快速排序
func QSort(elem []int, low, high int) []int {
	if low < high {
		pivot := partition(elem, low, high)
		QSort(elem, low, pivot - 1)
		QSort(elem, pivot + 1, high)
	}
	return elem
}

func partition(elem []int, low, high int) int {

	// 优化选取基数-三数取中法：取三个关键字先进行排序，将中间数作为基数
	mid := low + (high - low) /2
	if elem[low] > elem[high] {
		elem[low], elem[high] = elem[high], elem[low]
	}
	if elem[mid] > elem[high] {
		elem[mid], elem[high] = elem[high], elem[mid]
	}
	if elem[mid] > elem[low] {
		elem[mid], elem[low] = elem[low], elem[mid]
	}

	elem[0] = elem[low]

	for low < high {
		for low < high && elem[high] >= elem[0] {
			high--
		}
		elem[low] = elem[high] // 采用替换而不是交换，提高性能
		for low < high && elem[low] <= elem[0] {
			low++
		}
		elem[high] = elem[low]
	}
	elem[low] = elem[0]
	return low
}

