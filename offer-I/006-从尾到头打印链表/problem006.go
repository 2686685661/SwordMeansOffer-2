/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/6 11:13 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _06_从尾到头打印链表


type ListNode struct{
   Val int
   Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param head ListNode类
 * @return int整型一维数组
 */
var res []int
func printListFromTailToHead( head *ListNode ) []int {
	// write code here
	//newNode := Reverse(head)
	//newNode := ReverseRecursion(head)
	//var res []int
	//for newNode != nil {
	//	res = append(res, newNode.Val)
	//	newNode = newNode.Next
	//}
	//return res

	// 不改变链表结构
	//递归
	if head != nil {
		printListFromTailToHead(head.Next)
		res = append(res, head.Val)
	}
	return res
}

/** 改变链表结构 **/

// 反转链表- 递归
func ReverseRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	lnext := head.Next
	newNode := ReverseRecursion(lnext)
	lnext.Next = head
	head.Next = nil
	return newNode
}

// 反转链表- 循环
func Reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	var (
		lprev *ListNode
		lnext *ListNode
		newNode *ListNode
	)


	for current != nil {
		lnext = current.Next
		current.Next = lprev
		if lnext == nil {
			newNode = current
		}
		lprev = current
		current = lnext
	}
	return newNode
}

