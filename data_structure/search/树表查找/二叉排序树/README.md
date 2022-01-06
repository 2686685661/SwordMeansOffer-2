# 二叉排序树

### 查找-递归
时间复杂度O(log2n)
```
// 伪代码
type BSTree struct {
    Key int
    lchild *BSTree
    rchild *BSTree
}

func SearchBST(root *BSTree, key int) *BSTree {
    if root == nil || root.Key == key {
        return root
    } else if key < root.Key {
        return SearchBST(root.lchild, key)
    } else {
        return SearchBST(root.rchild, key)
    }
}
```

### 插入
插入操作以查找为基础，从根节点想下查找，当树中不存在关键字等于key的节点时才进行插入，新插入的节点一定是一个新添加的叶子节点，并且是查找不成功时查找路径上访问的最后一个节点的左右孩子

### 创建BST
就是循环将key插入的过程

### 删除
1. 删除节点为叶子节点，直接删除
2. 删除节点只有左子树Pl或右子树Pr，直接让Pl或Pr成为双亲节点的子树
3. 删除节点P的左右子树都不为空，让P的直接前驱或直接后继替代P，然后删除直接前驱或后继
