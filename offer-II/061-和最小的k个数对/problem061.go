/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/19 11:18 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _61_和最小的k个数对

import "container/heap"


// 大顶堆
type Heap [][]int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i][0] + h[i][1] > h[j][0] + h[j][1]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}


func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}


func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	h := &Heap{}
	heap.Init(h)

	n1, n2 := Min(len(nums1), k), Min(len(nums2), k)


	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			heap.Push(h, []int{nums1[i], nums2[j]})
			if h.Len() > k {
				heap.Pop(h)
			}
		}
	}

	var size = h.Len()
	var res = make([][]int, size)

	for i := 0; i <size; i++ {
		res[size - i - 1] =  heap.Pop(h).([]int)
	}
	return res
}