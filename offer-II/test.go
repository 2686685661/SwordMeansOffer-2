package test

import (
	"container/heap"
	"sort"
)

type IHeap [][]int
func (h IHeap) Len() int {
	return len(h)
}
func (h IHeap) Less(i, j int) bool {
	return h[i][0] + h[i][1] > h[j][0] + h[j][1]
}
func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) < 1 || len(nums2) < 1 || k < 1 {
		return nil
	}
	Min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	h := &IHeap{}
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
	res := make([][]int, 0)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).([]int))
	}
	return res
}





//
//
//
//
//type Trie struct {
//	child [26]*Trie
//	isEnd bool
//
//}
///** Initialize your data structure here. */
//func Constructor() Trie {
//	return Trie{}
//}
///** Inserts a word into the trie. */
//func (this *Trie) Insert(word string)  {
//	node := this
//	for _, ch := range word {
//		ch -= 'a'
//		if node.child[ch] == nil {
//			node.child[ch] = &Trie{}
//		}
//		node = node.child[ch]
//	}
//	node.isEnd = true
//}
///** Returns if the word is in the trie. */
//func (this *Trie) Search(word string) bool {
//	node := this.SearchPrefix(word)
//	return node != nil && node.isEnd
//}
///** Returns if there is any word in the trie that starts with the given prefix. */
//func (this *Trie) StartsWith(prefix string) bool {
//	node := this.SearchPrefix(prefix)
//	return node != nil
//}
//func (this *Trie) SearchPrefix(prefix string) *Trie {
//	node := this
//	for _, ch := range prefix {
//		ch -= 'a'
//		if node.child[ch] == nil {
//			return nil
//		}
//		node = node.child[ch]
//	}
//	return node
//}




type trie struct {
	child [2]*trie
	isEnd bool
}
const numSize = 30
func (t *trie) insert(num int) {
	node := t
	for i := numSize; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.child[bit] == nil {
			node.child[bit] = &trie{}
		}
		node = node.child[bit]
	}
	node.isEnd = true
}
func (t *trie) check(num int) int {
	node := t
	res := 0
	for i := numSize; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.child[bit ^ 1] != nil {
			res |= 1 << i
			node = node.child[bit ^ 1]
		} else {
			node = node.child[bit]
		}
	}
	return res
}
func findMaximumXOR(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	res := 0
	t := &trie{}
	t.insert(nums[0])
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1; i < len(nums); i++ {
		res = max(res, t.check(nums[i]))
		t.insert(nums[i])
	}
	return res
}





func searchInsert(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}

	low, mid, high := 0, 0, len(nums) - 1
	for low <= high {
		mid = (low + high) >> 1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	if nums[mid] < target {
		return mid + 1
	} else {
		return mid
	}
}



func singleNonDuplicate(nums []int) int {
	f1 := func(nums []int) int {
		res := 0
		for _, n := range nums {
			res ^= n
		}
		return res
	}

	f2 := func(nums []int) int {
		if len(nums) < 1 {
			return -1
		}
		low, mid, high := 0, 0, len(nums) - 1
		for low <= high {
			mid = (low + high) >> 1
			if mid < len(nums) - 1 && nums[mid] == nums[mid + 1] {
				if mid % 2 == 0 { // 偶奇
					low = mid + 2
				} else {
					high = mid -1
				}
			} else if mid > 0 && nums[mid] == nums[mid - 1] {
				if mid % 2 == 0 {
					high = mid - 2
				} else {
					low = mid + 1
				}
			} else {
				return nums[mid]
			}
		}
		return 0
	}
	f1(nums)
	return f2(nums)
}


func mySqrt(x int) int {
	if x < 0 {
		return 0
	}
	low, high := 1, x >> 1 + 1
	for low <= high {
		mid := (low + high) >> 1
		if mid * mid > x {
			high = mid - 1
		} else if mid * mid < x {
			low = mid + 1
		} else {
			return mid
		}
	}
	return high
}


func minEatingSpeed(piles []int, h int) int {
	if len(piles) <  1 || h < len(piles) {
		return 0
	}

	low, high := 1, 0
	for _, p := range piles {
		if p > high {
			high = p
		}
	}
	countTime := func(k int) (t int) {
		for _, p := range piles {
			t = p / k
			if p % k > 0 {
				t += 1
			}
		}
		return
	}
	for low <= high {
		mid := (low + high) >> 1
		if countTime(mid) < h {
			if mid == 1 || countTime(mid - 1) > h {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := make([][]int, 0)
	merged = append(merged, intervals[0])
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > merged[len(merged) - 1][1] {
			merged = append(merged, intervals[i])
		} else {
			merged[len(merged) - 1][1] = max(merged[len(merged) - 1][1], intervals[i][1])
		}
	}
	return merged
}


func relativeSortArray(arr1 []int, arr2 []int) []int {
	if len(arr1) < 1 || len(arr2) < 1 {
		return nil
	}

	max := 0
	for _, n := range arr1 {
		if n > max {
			max = n
		}
	}
	counts := make([]int, max + 1)

	for _, n := range arr1 {
		counts[n]++
	}
	var i int
	for _, n := range arr2 {
		for {
			if counts[n] <= 0 {
				break
			}
			arr1[i] = n
			i++
			counts[n]--
		}
	}
	for n, _ := range counts {
		for {
			if counts[n] <= 0 {
				break
			}
			arr1[i] = n
			i++
			counts[n]--
		}
	}
	return arr1
}

func findKthLargest(nums []int, k int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i - 1] {
			tmp := nums[i]
			j := 0
			for j = i - 1; j >= 0 && nums[j] > tmp; j-- {
				nums[j + 1] = nums[j]
			}
			nums[j + 1] = tmp
		}
	}
	return nums[len(nums) - k]
}



type ListNode struct {
	Val int
	Next *ListNode
}
func sortList(head *ListNode) *ListNode {

	var (
		splitList func(head *ListNode) *ListNode
		mergeList func(head1, head2 *ListNode) *ListNode
		mergeUp2DownSort func(head *ListNode) *ListNode
	)
	splitList = func(head *ListNode) *ListNode {
		slow, fast := head, head.Next
		for fast != nil &&& fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		midNode := slow.Next
		slow.Next = nil
		return midNode
	}
	mergeList = func(head1, head2 *ListNode) *ListNode {
		front := &ListNode{}
		cur := front
		for head1 != nil && head2 != nil {
			if head1.Val < head2.Val {
				cur.Next = head1
				head1 = head1.Next
			} else {
				cur.Next = head2
				head2 = head2.Next
			}
			cur = cur.Next
		}
		for head1 != nil {
			cur.Next = head1
			cur = cur.Next
			head1 = head1.Next
		}
		for head2 != nil {
			cur.Next = head2
			cur = cur.Next
			head2 = head2.Next
		}
		return front.Next
	}
	mergeUp2DownSort = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		head1, head2 := head, splitList(head)
		head1 = mergeUp2DownSort(head1)
		head2 = mergeUp2DownSort(head2)
		return mergeList(head1, head2)
	}
	return mergeUp2DownSort(head)

}


func subsets(nums []int) [][]int {

	var (
		res [][]int
		cur []int
		backtrack1 func(idx int)
		backtrack2 func(idx int)
	)

	backtrack1 = func(idx int) {
		if idx == len(nums) {
			res = append(res, append([]int{}, cur...))
			return
		}
		backtrack1(idx + 1)

		cur = append(cur, nums[idx])
		backtrack1(idx + 1)
		cur = cur[:len(cur)-1]
	}
	backtrack2 = func(idx int) {
		res = append(res, append([]int{}, cur...))
		for i := idx; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtrack2(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	//backtrack1(0)
	backtrack2(0)
	return res
}


func combine(n int, k int) [][]int {
	if n < 1 || n > 20 || k < 1 || k > n {
		return nil
	}
	var (
		cur []int
		res [][]int
		backtrack func(idx int)
	)
	backtrack = func(idx int) {
		if idx > n || len(cur) == k {
			if len(cur) == k {
				res = append(res, append([]int{}, cur...))
			}
			return
		}
		for i := idx; i <= n; i++ {
			cur = append(cur, i)
			backtrack(i + 1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(1)
	return res
}

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) < 1 || target < 1 {
		return nil
	}
	var (
		cur []int
		res [][]int
		backtrack func(idx, sum int)
	)
	sort.Ints(candidates)
	backtrack = func(idx, sum int) {
		if idx >= len(candidates) || sum > target {
			return
		}
		if sum == target {
			res = append(res, append([]int{}, cur...))
			return
		}
		for i := idx; i < len(candidates); i++ {
			if sum + candidates[i] <= target && idx <= i {
				sum += candidates[i]
				cur = append(cur, candidates[i])
				backtrack(i, sum)
				sum -= candidates[i]
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtrack(0, 0)
	return res
}


func permute(nums []int) [][]int {
	if len(nums) < 1 || len(nums) > 6 {
		return nil
	}
	var (
		res [][]int
		backtrack func(idx int)
	)
	backtrack = func(idx int) {
		if idx == len(nums) {
			res = append(res, append([]int{}, nums...))
			return
		}
		for i := idx; i < len(nums); i++ {
			nums[idx], nums[i] = nums[i], nums[idx]
			backtrack(idx + 1)
			nums[idx], nums[i] = nums[i], nums[idx]
		}
	}
	backtrack(0)
	return res
}

func permuteUnique(nums []int) [][]int {
	if len(nums) < 1 || len(nums) > 8 {
		return nil
	}
	var (
		res [][]int
		backtrack func(idx int)
	)
	backtrack = func(idx int) {
		if idx == len(nums) {
			res = append(res, append([]int{}, nums...))
			return
		}
		m := make(map[int]int)
		for i := idx; i < len(nums); i++ {
			if _, ok := m[nums[i]]; !ok || i == idx {
				m[nums[i]] = 1
				nums[idx], nums[i] = nums[i], nums[idx]
				backtrack(idx + 1)
				nums[idx], nums[i] = nums[i], nums[idx]
			}
		}
	}
	backtrack(0)
	return res
}

func generateParenthesis(n int) []string {
	if n < 1 || n > 8 {
		return nil
	}
	var (
		res []string
		backstrack func(cur string, left, right int)
	)
	backstrack = func(cur string, left, right int) {
		if left > n || right > n || left < right {
			return
		}
		if len(cur) == 2 * n && left == right {
			res = append(res, cur)
			return
		}
		backstrack(cur + "(", left + 1, right)
		backstrack(cur + ")", left, right + 1)
	}
	backstrack("", 0, 0)
	return res
}

func partition(s string) [][]string {

	if len(s) < 1 || len(s) > 16 {
		return nil
	}
	var (
		cur []string
		res [][]string
		backtrack func(idx int)
		isPalindrome func(s string) bool
	)
	isPalindrome = func(s string) bool {
		l, r := 0, len(s) - 1
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	backtrack = func(idx int) {
		if idx == len(s) {
			res = append(res, append([]string{}, cur...))
			return
		}
		for i := idx + 1; i <= len(s); i++ {
			if isPalindrome(s[idx:i]) {
				cur = append(cur, s[idx:i])
				backtrack(i)
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtrack(0)
	return res
}