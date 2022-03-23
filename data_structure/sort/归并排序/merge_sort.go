/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 4:58 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 归并排序


// 从上往下
func MergeSortUp2Down(elem []int, low, high int) []int {
	if len(elem) < 1 || low >= high {
		return elem
	}
	mid := (low + high) >> 1
	MergeSortUp2Down(elem, low, mid)
	MergeSortUp2Down(elem, mid + 1, high)

	elem = merge(elem, low, mid, high)

	return elem
}

// 从下往上
func MergeSortDown2Up(elem []int) []int {
	if len(elem) < 1 {
		return elem
	}

	for i := 1; i < len(elem); i *= 2 {
		elem = mergeGroups(elem, len(elem), i)
	}
	return elem
}

func mergeGroups(elem []int, len, gap int) []int {
	var i int
	for i = 0; i + 2 * gap - 1 < len; i += 2 * gap {
		merge(elem, i, i + gap  -1, i + 2 * gap - 1)
	}
	if i + gap -1 < len - 1 {
		merge(elem, i, i + gap - 1, len - 1)
	}
	return elem
}

func merge(elem []int, low, mid, high int) []int {
	i, k, j := low, 0, mid + 1
	tmp := make([]int, high - low + 1)
	for i <= mid && j <= high {
		if elem[i] < elem[j] {
			tmp[k] = elem[i]
			k++
			i++
		} else {
			tmp[k] = elem[j]
			k++
			j++
		}
	}

	for i <= mid {
		tmp[k] = elem[i]
		k++
		i++
	}
	for j <= high {
		tmp[k] = elem[j]
		k++
		j++
	}
	for i = 0; i < k; i++ {
		elem[low + i] = tmp[i]
	}
	return elem
}


