/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/29 12:09 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _25_链表中的两数相加

type ListNode struct {
	Val int
	Next *ListNode
}


// 反转链表实现，会更改链表结构
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l1, l2 = reverse(l1), reverse(l2)
	carry := 0

	var node, next *ListNode
	node = &ListNode{0, nil}

	for l1 != nil || l2 != nil || carry != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		sum := v1 + v2 + carry
		carry = sum / 10
		node = &ListNode{
			Val:  sum % 10,
			Next: next,
		}
		next = node

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return node
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}


// 栈实现
type Stack []int
func (s Stack) Len() int {
	return len(s)
}
func (s Stack) Cap() int {
	return cap(s)
}

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

func (s Stack) Top() int {
	if len(s) == 0 {
		return 0
	}
	return s[len(s) - 1]
}

func (s *Stack) Pop() int {
	ts := *s
	if len(ts) == 0 {
		return 0
	}
	v := ts[len(ts) - 1]
	*s = ts[:len(ts) - 1]
	return v
}

func (s Stack) IsEmpty() bool  {
	return len(s) == 0
}
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := Stack{}, Stack{}

	for l1 != nil {
		s1.Push(l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		s2.Push(l2.Val)
		l2 = l2.Next
	}

	carry := 0
	var node, next *ListNode
	node = &ListNode{0, nil}

	for !s1.IsEmpty() || !s2.IsEmpty() || carry != 0 {
		v1, v2 := 0, 0
		if !s1.IsEmpty() {
			v1 = s1.Pop()
		}
		if !s2.IsEmpty() {
			v2 = s2.Pop()
		}
		sum := v1 + v2 + carry
		carry = sum / 10
		node = &ListNode{
			Val:  sum % 10,
			Next: next,
		}
		next = node
	}
	return node

}

