/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/22 8:17 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strconv"
)

// 异或：相同为0，不同为1
func main() {
	//fmt.Println(4 ^ 0)

	//maxProduct([]string{"ab", "cd"})

	//nums := []int{-1,0,1,2,-1,-4}
	//sort.Ints(nums)
	//fmt.Println(nums)

	//m := make(map[int]int)
	//m[1] = 1
	//m[2] = 2
	//m[3] = 3
	//if _,ok := m[3]; ok {
	//	fmt.Println("fuck")
	//}
	//for k, _ := range m {
	//	fmt.Println(k)
	//	//break
	//}

	//fmt.Println(1 ^ 1)
	//set := &treap{}
	//for i := 0; i < 10; i++ {
	//	set.Put(i)
	//}
	//
	//
	//inorder(set.root)
	//fmt.Printf("\n")
	//
	//set.Del(3)
	//inorder(set.root)
	arr := []string{"me", "time"}
	sort.Strings(arr)
	fmt.Println(arr)


}





func inorder(n *node) {
	if n == nil {
		return
	}
	inorder(n.ch[0])
	fmt.Printf("%d\t", n.val)
	inorder(n.ch[1])
}

// 树节点
type node struct {
	priority int
	val int
	ch [2]*node
}

// 二分比较
func (n *node) contrast(val int) (res int) {
	res = -1
	if val < n.val {
		res = 0
	} else if val > n.val {
		res = 1
	}
	return
}

// 树旋转
// d=0 对o进行左旋
// d=1 对o进行右旋
func (n *node) rotate(d int) *node {
	pivot := n.ch[d ^ 1] // 左旋，则转轴是o的右子节点；右旋，则转轴是o的左子节点
	n.ch[d ^ 1] = pivot.ch[d]
	pivot.ch[d] = n
	return pivot
}


// 树堆结构
type treap struct {
	root *node
}


// 插入树堆
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
	root.ch[d] = t.put(root.ch[d], val)
	if root.ch[d].priority > root.priority {
		root = root.rotate(d ^ 1) //插入节点是o的左子节点，进行右旋；插入节点是o的右子节点，进行左旋
	}
	return root
}


// 从树堆中删除
func (t *treap) Del(val int) {
	t.root = t.del(t.root, val)
}

func (t *treap) del(root *node, val int) *node {
	if d := root.contrast(val); d >= 0 {
		root.ch[d] = t.del(root.ch[d], val)
		return root
	}

	if root.ch[0] == nil {
		return root.ch[1]
	}
	if root.ch[1] == nil {
		return root.ch[0]
	}

	d := 0
	if root.ch[0].priority > root.ch[1].priority {
		d = 1
	}
	root = root.rotate(d)
	root.ch[d] = t.del(root.ch[d], val)
	return root
}






















const highBit = 30

type trie struct {
	left, right *trie
}

func (t *trie) add(num int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			if cur.left == nil {
				cur.left = &trie{}
			}
			cur = cur.left
		} else {
			if cur.right == nil {
				cur.right = &trie{}
			}
			cur = cur.right
		}
	}
}

func (t *trie) check(num int) (x int) {
	cur := t
	for i := highBit; i >= 0; i-- {
		bit := num >> i & 1
		if bit == 0 {
			// a_i 的第 k 个二进制位为 0，应当往表示 1 的子节点 right 走
			if cur.right != nil {
				cur = cur.right
				x = x*2 + 1
			} else {
				cur = cur.left
				x = x * 2
			}
		} else {
			// a_i 的第 k 个二进制位为 1，应当往表示 0 的子节点 left 走
			if cur.left != nil {
				cur = cur.left
				x = x*2 + 1
			} else {
				cur = cur.right
				x = x * 2
			}
		}
	}
	return
}

func findMaximumXOR(nums []int) (x int) {
	root := &trie{}
	for i := 1; i < len(nums); i++ {
		// 将 nums[i-1] 放入字典树，此时 nums[0 .. i-1] 都在字典树中
		root.add(nums[i-1])
		// 将 nums[i] 看作 ai，找出最大的 x 更新答案
		x = max(x, root.check(nums[i]))
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}



func divide(a int, b int) int {
	if b == 0 {
		return 0
	}

	if a == math.MinInt32 {
		if b == -1 {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}

	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	a, b = abs(a), abs(b)
	res := 0

	for i := 31; i >= 0; i-- {
		if a >> i >= b {
			a -= b << i
			res += 1 << i
		}
	}
	return res * sign
}


func addBinary(a string, b string) string {

	al, bl := len(a), len(b)
	if al < 1 || bl < 1 {
		return ""
	}

	carry, res := 0, ""

	max := func(x, y int) int{
		if x > y {
			return x
		}
		return y
	}

	n := max(al, bl)
	for i := 0; i < n; i++ {
		if i < al {
			carry += int(a[al - i - 1] - '0')
		}
		if i < bl {
			carry += int(b[bl - i - 1] - '0')
		}
		res = strconv.Itoa(carry % 2) + res
		carry /= 2
	}
	if carry > 0 {
		res = "1" + res
	}
	return res
}


func countBits(n int) []int {
	bits := make([]int, n + 1)
	onceCount := func(x int) (ones int) {
		for ; x > 0; x &= x - 1 {
			ones++
		}
		return
	}
	for i := range bits {
		bits[i] = onceCount(i)
	}
	return bits

}


func singleNumber(nums []int) int {
	res := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total % 3 > 0 {
			res |= 1 << i
		}
	}
	return int(res)
}


func maxProduct(words []string) int {
	if words == nil || len(words) < 2 {
		return 0
	}

	bitRes := make([]int, len(words))
	for i := 0; i < len(bitRes); i++ {
		for _, c := range words[i] {
			bitRes[i] |= 1 << (c - 'a')
		}
	}

	res := 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 0; i < len(bitRes); i++ {
		for j := i + 1; j < len(bitRes); j++ {
			if bitRes[i] & bitRes[j] == 0 {
				res = max(res, len(words[i]) * len(words[j]))
			}
		}
	}
	return res
}

func twoSum(numbers []int, target int) []int {

	if numbers == nil || len(numbers) < 2 {
		return nil
	}

	start, end := 0, len(numbers) - 1
	for start < end {
		if numbers[start] + numbers[end] < target {
			start++
		} else if numbers[start] + numbers[end] > target {
			end--
		} else {
			return []int{start, end}
		}
	}
	return nil
}



func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)

	var res [][]int
	for i := 0; i < len(nums) - 2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		start, end := i + 1, len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum < 0 {
				start++
				for start < end && nums[start] == nums[start - 1] {
					start++
				}
			} else if sum > 0 {
				end--
				for start < end && nums[end] == nums[end + 1] {
					end--
				}
			} else {
				res = append(res, []int{nums[i], nums[start], nums[end]})
				start++
				end--
				for start < end && nums[start] == nums[start - 1] {
					start++
				}
				for start < end && nums[end] == nums[end + 1] {
					end--
				}
			}
		}
	}
	return res
}


func minSubArrayLen(target int, nums []int) int {
	if nums == nil || len(nums) < 1 || target < 1 {
		return 0
	}
	res := len(nums)
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	sum := 0
	for i, j := 0, 0; j < len(nums); j++ {
		sum += nums[j]
		for i <= j && sum >= target {
			res = min(res, j - i + 1)
			sum -= nums[i]
			i++
		}
	}
	return res
}


func numSubarrayProductLessThanK(nums []int, k int) int {
	if nums == nil || len(nums) < 1 || k < 0 {
		return 0
	}

	count, sum := 0, 1
	for i, j := 0, 0; j < len(nums); j++ {
		sum *= nums[j]
		for i <= j && sum >= k {
			sum /= nums[i]
			i++
		}
		if i <= j {
			count += j - i +1
		}
	}
	return count
}

func subarraySum(nums []int, k int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}

	preSum, count  := 0, 0
	dict := make(map[int]int)
	dict[0] = 1
	for _, n := range nums {
		preSum += n
		count += dict[preSum - k]
		dict[preSum] += 1
	}
	return count
}


func findMaxLength(nums []int) int {

	if nums == nil || len(nums) < 1 {
		return 0
	}
	for i, v := range nums {
		if v == 0 {
			nums[i] = -1
		}
	}
	preSum, ans := 0, 0
	dict := make(map[int]int)
	dict[0] = -1

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i, n := range nums {
		preSum += n
		if index, ok := dict[preSum]; ok {
			ans = max(ans, i - index)
		} else {
			dict[preSum] = i
		}
	}

	return ans
}



func pivotIndex(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return -1
	}

	total, preSum := 0, 0
	for _, n := range nums {
		total += n
	}

	for i, n := range nums {
		if total - n == 2 * preSum {
			return i
		}
		preSum += n
	}
	return -1
}


// 013-begin
type NumMatrix struct {
	sums [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	sums := make([][]int, m + 1)
	sums[0] = make([]int, n + 1)
	for i, row := range matrix {
		sums[i + 1] = make([]int, n + 1)
		for j, v := range row {
			sums[i + 1][j + 1] = sums[i + 1][j] + sums[i][j + 1] - sums[i][j] + v
		}
	}
	return NumMatrix{sums:sums}
}


func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	sum = nm.sums[row2 + 1][col2 + 1] - nm.sums[row1][col2 + 1] - nm.sums[row2 + 1][col1] + nm.sums[row1][col1]
	return sum
}
// 013-end
