/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/3 3:51 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package tree

import "fmt"

/*
实现二叉树的基本操作
创建一棵树 done
还原一棵树 done
复制一棵树 done
计算树的高度 done
计算树的所有节点数 done
计算树的叶子节点数 done
二叉树的先序、中序、后序 递归遍历 done
二叉树的先序、中序、后序 非递归遍历
 */

const (
	PreOrderTraverse = iota // 先序遍历
	InOrderTraverse         // 中序遍历
	PostOrderTraverse       // 后序遍历
)

type BiTreeNode struct {
	data string          // 节点值
	left *BiTreeNode     // 左子树
	right *BiTreeNode    // 右子树
	iCount int           // 计算
	strTree string       // 字符串形式二叉树结构，AB##C##
	TraverseType int     // 遍历类型
}

func NewBiTree(strTree string, traverseType int) *BiTreeNode {
	return &BiTreeNode{
		iCount:       -1,
		strTree:      strTree,
		TraverseType: traverseType,
	}
}

// 创建二叉树：先序、后序
// 中序序列不能确定一颗二叉树
func (t *BiTreeNode) CreateTree() *BiTreeNode {
	var tree *BiTreeNode
	switch t.TraverseType {
	case PreOrderTraverse:
		tree = t.createTreePre()
	case InOrderTraverse:
		fmt.Println("中序序列无法确定一颗二叉树，会有不同的树对应一个中序输入")
		return nil
	case PostOrderTraverse:
		tmp := t.strTree
		r := []rune(tmp)
		for i,j := 0, len(r) - 1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		t.strTree = string(r)
		tree = t.createTreePost()
		t.strTree = tmp
	}

	return tree
}

// 先序序列创建二叉树
func (t *BiTreeNode) createTreePre() *BiTreeNode {
	t.iCount++
	if t.iCount >= len(t.strTree) {
		return nil
	}

	var node *BiTreeNode
	s := t.strTree[t.iCount]
	if s == '#' {
		return nil
	} else {
		node = new(BiTreeNode)
		node.data = string(s)
		node.left = t.createTreePre()
		node.right = t.createTreePre()
	}
	return node
}

// 后序序列创建二叉树
func (t *BiTreeNode) createTreePost() *BiTreeNode {
	t.iCount++
	if t.iCount >= len(t.strTree) {
		return nil
	}
	var node *BiTreeNode
	s := t.strTree[t.iCount]
	if s == '#' {
		return nil
	} else {
		node = new(BiTreeNode)
		node.data = string(s)
		node.right = t.createTreePost()
		node.left = t.createTreePost()
	}
	return node
}

// 二叉树的先序，中序，后序遍历 递归版本
func (t *BiTreeNode) Tree2String(tree *BiTreeNode) string {
	if tree == nil {
		return ""
	}
	var str string

	lPound := func() {
		if tree.left == nil {
			str += "#"
		}
	}
	rPound := func() {
		if tree.right == nil {
			str += "#"
		}
	}
	switch t.TraverseType {
	case PreOrderTraverse:
		str += tree.data
		lPound()
		str += t.Tree2String(tree.left)
		rPound()
		str += t.Tree2String(tree.right)
	case InOrderTraverse:
		str += t.Tree2String(tree.left)
		lPound()
		str += tree.data
		str += t.Tree2String(tree.right)
		rPound()
	case PostOrderTraverse:
		lPound()
		str += t.Tree2String(tree.left)
		str += t.Tree2String(tree.right)
		rPound()
		str += tree.data
	}
	return str
}


// 二叉树的先序遍历 非递归
// 第一次就访问节点值，然后进栈，依次访问其左孩子，直到左孩子为空弹出栈顶，访问其右孩子，循环直到栈为空，tree也为空
// 参考：https://blog.csdn.net/LYue123/article/details/88836008
func (t *BiTreeNode) Tree2StringNoRecursivePre(tree *BiTreeNode) string {

	var str string

	lPound := func() {
		if tree.left == nil {
			str += "#"
		}
	}
	rPound := func() {
		if tree.right == nil {
			str += "#"
		}
	}

	stack := []*BiTreeNode{}
	for tree != nil || len(stack) != 0 {
		if tree != nil {
			str += tree.data
			lPound()
			stack = append(stack, tree)
			tree = tree.left
		} else {
			tree = stack[len(stack) - 1]
			stack = stack[:len(stack)-1]
			rPound()
			tree = tree.right
		}
	}
	return str
}

// 二叉树的中序遍历 非递归
// 和前序遍历类似，只不过从栈中弹出的时候才访问节点值
func (t *BiTreeNode) Tree2StringNoRecursiveIn(tree *BiTreeNode) string {

	var str string

	lPound := func() {
		if tree.left == nil {
			str += "#"
		}
	}
	rPound := func() {
		if tree.right == nil {
			str += "#"
		}
	}

	stack := []*BiTreeNode{}
	for tree != nil || len(stack) != 0 {
		if tree != nil {
			stack = append(stack, tree)
			tree = tree.left
		} else {
			tree = stack[len(stack) - 1]
			lPound()
			str += tree.data
			stack = stack[:len(stack)-1]
			rPound()
			tree = tree.right
		}
	}
	return str
}

// 二叉树的后序遍历 非递归
// 这里要增加一个辅助节点来保存上一次放入结果数组的值。
// 当一个节点左右都是空的时候，就可以放入结果集。
// 当上一个放入结果集的节点是他的孩子节点的时候，说明他的孩子已经访问完成了，到这里就是第三次了就可以放入了。
// 还有一点要注意的是，当一个节点的左右不为空时，要先加入右孩子，再加入左孩子，这样才能先访问左孩子。
func (t *BiTreeNode) Tree2StringNoRecursivePost(tree *BiTreeNode) string {
	var str string
	if tree == nil {
		return str
	}

	stack := []*BiTreeNode{}
	pre := &BiTreeNode{}
	stack = append(stack, tree)
	for len(stack) !=0 {
		cur := stack[len(stack)-1]
		if (cur.left == nil && cur.right == nil) || (pre != nil && (pre == cur.left || pre == cur.right)) {
			if cur.left == nil {
				str += "#"
			}
			if cur.right == nil {
				str += "#"
			}
			str += cur.data
			pre = cur
			stack = stack[:len(stack)-1]
		} else {
			if cur.right != nil {
				stack = append(stack, cur.right)
			}
			if cur.left != nil {
				stack = append(stack, cur.left)
			}
		}
	}
	return str
}

// 二叉树的层序遍历
// 借用一个队列依次将根节点，左子节点，右子节点添加到队列，依次先进先出
func (t *BiTreeNode) TreeSequence(tree *BiTreeNode) string {
	queue := []*BiTreeNode{}
	queue = append(queue, tree)
	str := ""
	for len(queue) > 0 {
		node := queue[0]
		str += node.data
		queue = queue[1:]
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	return str

}

// 二叉树的复制
// 首先复制根节点，分别复制二叉树根节点的左子树和右子树
// 本质上于二叉树先序遍历的实现非常相似
func (t *BiTreeNode) Copy(tree *BiTreeNode) *BiTreeNode {
	if tree == nil {
		return nil
	}

	node := new(BiTreeNode)
	node.data = tree.data
	node.left = t.Copy(tree.left)
	node.right = t.Copy(tree.right)
	return node
}

// 计算二叉树的深度
// 如果是空树，递归结束，深度为0
// 递归计算左子树的深度为m;递归计算右子树的深度为n
// 如果m大于n，二叉树的深度为m+1, 否则为n+1
func (t *BiTreeNode) Depth(tree *BiTreeNode) int {
	if tree == nil {
		return 0
	}

	m := t.Depth(tree.left)
	n := t.Depth(tree.right)
	if m > n {
		return m + 1
	} else {
		return n + 1
	}
}

// 计算二叉树中节点的个数
// 节点个数为左子树的节点个数加上右子树的节点个数再加上1
func (t *BiTreeNode) NodeCount(tree *BiTreeNode) int {
	if tree == nil {
		return 0
	}
	return t.NodeCount(tree.left) + t.NodeCount(tree.right) + 1
}

// 计算二叉树的叶子节点个数
// 左子树和右子树都为空，为叶子节点，否则不是
// 将左子树和右子树计算的叶子结点相加得出最终的叶子结点数
func (t *BiTreeNode) LeadCount(tree *BiTreeNode) int {
	if tree == nil {
		return 0
	}

	if tree.left == nil && tree.right == nil {
		return 1
	}
	return t.LeadCount(tree.left) + t.LeadCount(tree.right)
}