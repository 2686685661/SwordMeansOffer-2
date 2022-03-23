/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/26 11:06 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package data_structure
import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSort(t *testing.T) {
	elem := []int{5, 3, 4, 6, 2}
	fmt.Println("InsertSort: ", InsertSort(elem))
	fmt.Println("BInsertSort: ", BInsertSort(elem))
	fmt.Println("ShellSort: ", ShellSort(elem))
	fmt.Println("BubbleSort: ", BubbleSort(elem))
	fmt.Println("QuickSort: ", QuickSort(elem, 0, len(elem) - 1))
	fmt.Println("SelectSort: ", SelectSort(elem))
	fmt.Println("HeapSort: ", HeapSort(elem))
	fmt.Println("MergeSort up2down: ", MergeSort(elem, 0))
	fmt.Println("MergeSort down2up: ", MergeSort(elem, 1))
	fmt.Println("BucketSort: ", BucketSort(elem))
	fmt.Println("CountSort: ", CountSort(elem))
	fmt.Println("RadixSort: ", RadixSort(elem))

fmt.Sprintln()

}


/**    --插入排序--      **/


// 直接插入排序
func InsertSort(elem []int) []int {
	if elem == nil || len(elem) < 2 {
		return elem
	}

	for i := 1; i < len(elem); i++ {
		if elem[i] < elem[i - 1] {
			tmp := elem[i]
			var j int
			for j = i - 1; j >=0 && elem[j] > tmp; j-- {
				elem[j + 1] = elem[j]
			}
			elem[j + 1] = tmp
		}
	}
	return elem
}

// 折半插入排序
func BInsertSort(elem []int) []int {
	if elem == nil || len(elem) < 2 {
		return elem
	}

	for i := 1; i < len(elem); i++ {
		tmp := elem[i]
		low, high := 0, i - 1
		for low <= high {
			mid := (low + high) >> 1
			if tmp > elem[mid] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		var j int
		for j = i - 1; j >= high + 1; j-- {
			elem[j + 1] = elem[j]
		}
		elem[high + 1] = tmp
	}
	return elem
}

// 希尔排序
func ShellSort(elem []int) []int {
	if elem == nil || len(elem) < 2 {
		return elem
	}
	j, d := 0, len(elem)
	for {
		d = d / 3 + 1
		for i := d + 1; i < len(elem); i++ {
			if elem[i] < elem[i - d] {
				tmp := elem[i]
				for j = i - d; j > 0 && elem[j] > tmp; j -= d {
					elem[j + d] = elem[j]
				}
				elem[j + d] = tmp
			}
		}
		if d <= 1 {
			break
		}
	}
	return elem
}



/**    --交换排序--      **/

// 冒泡排序
func BubbleSort(elem []int) []int {
	if elem == nil || len(elem) < 2 {
		return elem
	}
	flag := true
	for i := 0; i < len(elem) && flag; i++ {
		flag = false
		for j := len(elem) - 1; j > i; j-- {
			if elem[j - 1] > elem[j] {
				elem[j - 1], elem[j] = elem[j], elem[j - 1]
				flag = true
			}
		}
	}
	return elem
}

// 快速排序
func QuickSort(elem []int, low, high int) []int {
	if elem == nil || len(elem) < 2 {
		return elem
	}
	//if low < high {
	//	pivot := partition(elem, low, high)
	//	QuickSort(elem, low, pivot - 1)
	//	QuickSort(elem, pivot + 1, high)
	//}
	// 使用尾递归优化递归操作
	for low < high {
		pivot := partition(elem, low, high)
		QuickSort(elem, low, pivot - 1)
		low = pivot + 1
	}

	return elem
}
func partition(elem []int, low, high int) int {
	// 三数取中法
	mid := (low + high) >> 1
	if elem[low] > elem[high] {
		elem[low], elem[high] = elem[high], elem[low]
	}
	if elem[mid] > elem[high] {
		elem[mid], elem[high] = elem[high], elem[mid]
	}
	if elem[mid] > elem[low] {
		elem[mid], elem[low] = elem[low], elem[mid]
	}

	tmp := elem[low]

	for low < high {
		for low < high && elem[high] >= tmp {
			high--
		}
		elem[low] = elem[high] // 采用替换而不是交换，提高性能
		for low < high && elem[low] <= tmp {
			low++
		}
		elem[high] = elem[low]
	}
	elem[low] = tmp
	return low
}



/**    --选择排序--      **/


// 简单选择排序
func SelectSort(elem []int) []int {
	if elem == nil || len(elem) < 2 {
		return elem
	}

	for i := 0; i < len(elem); i++ {
		k := i
		for j := i + 1; j < len(elem); j++ {
			if elem[k] > elem[j] {
				k = j
			}
		}
		if k != i {
			elem[i], elem[k] = elem[k], elem[i]
		}
	}
	return elem
}

// 堆排序
func HeapSort(elem []int) []int {
	if elem == nil || len(elem) < 2 {
		return elem
	}

	var maxHeapAdjust func(elem []int, s, m int) []int
	maxHeapAdjust = func(elem []int, s, m int) []int {
		tmp := elem[s]
		for j := 2 * s; j <= m; j *= 2 {
			if j < m && elem[j] < elem[j + 1] {
				j += 1
			}
			if tmp >= elem[j] {
				break
			}
			elem[s] = elem[j]
			s = j
		}
		elem[s] = tmp
		return elem
	}

	for i := len(elem) / 2; i >= 0; i-- {
		elem = maxHeapAdjust(elem, i, len(elem) - 1)
	}

	for i := len(elem) - 1; i >= 0; i-- {
		elem[0], elem[i] = elem[i], elem[0]
		elem = maxHeapAdjust(elem, 0, i - 1)
	}
	return elem
}



/**    --归并排序--      **/

func MergeSort(elem []int, flag int) []int {
	if elem == nil || len(elem) < 2 {
		return elem
	}
	if flag == 0 {
		return mergeUp2DownSort(elem, 0, len(elem) - 1)
	} else {
		return mergeDown2UpSort(elem)
	}
}

// 从上往下归并排序
func mergeUp2DownSort(elem []int, low, high int) []int {
	if low >= high {
		return elem
	}
	mid := (low + high) >> 1
	mergeUp2DownSort(elem, low, mid)
	mergeUp2DownSort(elem, mid + 1, high)
	elem = merge(elem, low, mid, high)
	return elem
}

// 从下往上归并排序
func mergeDown2UpSort(elem []int) []int {
	mergeGroup := func(elem []int, gap, len int) []int {
		var i int
		for i = 0; i + 2 * gap - 1 < len; i += 2 * gap {
			merge(elem, i, i + gap - 1, i + 2 * gap - 1)
		}
		if i + gap - 1 < len {
			merge(elem, i, i + gap - 1, len - 1)
		}
		return elem
	}

	for i := 1; i < len(elem); i *= 2 {
		elem = mergeGroup(elem, i, len(elem))
	}
	return elem
}

func merge(elem []int, low, mid, high int) []int {
	i, j := low, mid + 1
	tmp := make([]int, high - low + 1)
	var k int
	for i <= mid && j <= high {
		if elem[i] < elem[j] {
			tmp[k] = elem[i]
			i++
		} else {
			tmp[k] = elem[j]
			j++
		}
		k++
	}

	for i <= mid {
		tmp[k] = elem[i]
		i++
		k++
	}
	for j <= high {
		tmp[k] = elem[j]
		j++
		k++
	}
	for i = 0; i < k; i++ {
		elem[low + i] = tmp[i]
	}
	return elem
}




/**    --非比较类排序--      **/

// 桶排序
func BucketSort(elem []int) []int {
	if elem == nil || len(elem) < 2 {
		return elem
	}
	minVal, maxVal := elem[0], elem[0]
	for _, n := range elem {
		if minVal > n {
			minVal = n
		}
		if maxVal < n {
			maxVal = n
		}
	}
	bucketSize := (maxVal - minVal) / len(elem) + 1
	buckets := make([][]int, bucketSize)

	insertSort := func(elem []int) []int {
		for i := 1; i < len(elem); i++ {
			if elem[i] < elem[i - 1] {
				tmp, j := elem[i], 0
				for j = i - 1; j >= 0 && elem[j] > tmp; j-- {
					elem[j + 1] = elem[j]
				}
				elem[j + 1] = tmp
			}
		}
		return elem
	}
	// 入桶
	for _, n := range elem {
		buckets[(n - minVal) / len(elem)] = append(buckets[(n - minVal) / len(elem)], n)
	}
	// 桶内排序
	for i, b := range buckets {
		if len(b) >= 2 {
			buckets[i] = insertSort(b)
		}
	}

	// 出桶
	i := 0
	for _, b := range buckets {
		for _,n := range b {
			elem[i] = n
			i++
		}
	}

	return elem
}


// 计数排序
func CountSort(elem []int) []int {
	if elem == nil || len(elem) < 2 {
		return elem
	}

	maxVal := elem[0]
	for _, n := range elem {
		if maxVal < n {
			maxVal = n
		}
	}

	counts := make([]int, maxVal + 1)

	for _, n := range elem {
		counts[n]++
	}

	for i, j := 0, 0; i < maxVal; i++ {
		for {
			if counts[i] <= 0 {
				break
			}
			elem[j] = i
			j++
			counts[i]--
		}
	}
	return elem
}


// 基数排序
func RadixSort(elem []int) []int {
	if elem == nil || len(elem) < 2 {
		return elem
	}
	maxVal := elem[0]
	for _, n := range elem {
		if maxVal < n {
			maxVal = n
		}
	}

	countSort := func(elem []int, exp int) []int {
		output, buckets := make([]int, len(elem)), make([]int, 10)

		// 入桶
		for _, n := range elem {
			buckets[(n / exp) % 10]++
		}
		//更改buckets[i],目的是让更改后的buckets[i]的值,是该数据在output[]中的位置。
		for i := 1; i < 10; i++ {
			buckets[i] += buckets[i - 1]
		}
		for i := len(elem) - 1; i >= 0; i-- {
			output[buckets[(elem[i] / exp) % 10] - 1] = elem[i]
			buckets[(elem[i] / exp) % 10]--
		}
		return output
	}

	for exp := 1; maxVal / exp > 0; exp *= 10 {
		elem = countSort(elem, exp)
	}
	return elem
}



type node struct {
	priority int
	val int
	child [2]*node
}
func ( n *node) contrast(val int) (res int) {
	res = -1
	if n.val < val {
		res = 0
	} else if n.val > val {
		res = 1
	}
	return
}

func (n *node) rotate(d int) *node {
	pivot := n.child[d ^ 1]
	n.child[d ^ 1] = pivot.child[d]
	pivot.child[d] = n
	return pivot
}


type treap struct {
	root *node
}

func (t *treap) Put(val int) {
	t.root = t.put(t.root, val)
}

func (t *treap) put(root *node, val int) *node {
	if root == nil {
		return &node{
			priority: rand.Int(),
			val:      val,
		}
	}
	d := root.contrast(val)
	root.child[d] = t.put(root.child[d], val)
	if root.child[d].priority > root.priority {
		root = root.rotate(d ^ 1)
	}
	return root
}

func (t *treap) Del(val int) {
	t.root = t.del(t.root, val)
}
func (t *treap) del(root *node, val int) *node {
	if d := root.contrast(val); d >= 0 {
		root.child[d] = t.del(root.child[d], val)
		return root
	}
	if root.child[0] == nil {
		return root.child[1]
	}
	if root.child[1] == nil {
		return root.child[0]
	}
	d := 0
	if root.child[0].priority > root.child[1].priority {
		d = 1
	}
	root = root.rotate(d)
	root.child[d] = t.del(root.child[d], val)
	return root
}