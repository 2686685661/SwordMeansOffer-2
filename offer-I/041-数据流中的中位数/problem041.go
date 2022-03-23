/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/29 6:03 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _41_数据流中的中位数

import (
	"container/heap"
)


var (
	max = make(MaxHeap, 0)
	min = make(MinHeap, 0)
)
func heapAdjust() {
	heap.Init(&max)
	heap.Init(&min)
}


func Insert(num int){
	if ((max.Len() + min.Len()) & 1) == 0 {
		if max.Len() > 0 && num < max.Top() {
			max.Push(num)
			num = (max.Pop()).(int)
		}
		min.Push(num)
	} else {
		if min.Len() > 0 && num > min.Top() {
			min.Push(num)
			num = (min.Pop()).(int)
		}
		max.Push(num)
	}
}

func GetMedian() float64{
	count := min.Len() + max.Len()
	if count == 0 {
		return 0.0
	}
	if count & 1 == 1 {
		return float64(min.Top())
	} else {
		return float64((min.Top() + max.Top()) / 2)
	}
}



// 大顶堆
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	// 由于是最大堆，所以使用大于号
	return h[i] > h[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{}{
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
func (h *MaxHeap) Top() int {
	res := (*h)[len(*h)-1]
	return res
}


// 小顶堆
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	// 由于是最大堆，所以使用小于号
	return h[i] < h[j]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{}{
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
func (h *MinHeap) Top() int {
	res := (*h)[len(*h)-1]
	return res
}