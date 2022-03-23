/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/3/16 5:03 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package data_structure

import "container/heap"

/**
Top K问题
https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/solution/tu-jie-top-k-wen-ti-de-liang-chong-jie-fa-you-lie-/

1. 堆
2. 快速选择
 */
// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。


// 大顶堆
type Heap []int
func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}
func getLeastNumbers1(arr []int, k int) []int {
	if len(arr) < k {
		return arr
	}
	if  k < 1 {
		return nil
	}
	h := &Heap{}
	heap.Init(h)
	for _, n := range arr {
		heap.Push(h, n)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, 0)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}
	return res

}





func getLeastNumbers2(arr []int, k int) []int {
	if len(arr) < k {
		return arr
	}
	if  k < 1 {
		return nil
	}
	return quickSearch(arr, 0, len(arr) - 1, k)
}
func quickSearch(arr []int, low, high, k int) []int {
	mid := partition1(arr, low, high)
	if mid == k - 1 {
		return arr[:k]
	} else if mid > k - 1 {
		return quickSearch(arr, low, mid - 1, k)
	} else {
		return quickSearch(arr, mid + 1, high, k)
	}
}

func partition1(arr []int, low, high int) int {
	tmp := arr[low]

	for low < high {
		for low < high && arr[high] >= tmp {
			high--
		}
		arr[low] = arr[high] // 采用替换而不是交换，提高性能
		for low < high && arr[low] <= tmp {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = tmp
	return low
}