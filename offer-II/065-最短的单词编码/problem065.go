/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/20 2:12 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _65_最短的单词编码

/**
字典树
https://leetcode-cn.com/problems/iSwD2y/solution/zui-duan-de-dan-ci-bian-ma-by-leetcode-s-qjxw/

去找到是否不同的单词具有相同的后缀，我们可以将其反序之后插入字典树中。例如，我们有 "time" 和 "me"，可以将 "emit" 和 "em" 插入字典树中。
然后，字典树的叶子节点（没有孩子的节点）就代表没有后缀的单词，统计叶子节点代表的单词长度加一的和即为我们要的答案。
 */
func minimumLengthEncoding(words []string) int {

	if words == nil || len(words) == 0 {
		return 0
	}

	res := 0
	t := &Trie{}
	for _, word := range words {
		if t.SearchPrefix(word) {
			continue
		}
		t.Insert(word)
	}
	var dfs func(node *Trie, depth int)
	dfs = func(node *Trie, depth int) {
		var hasChild bool
		for i := 0; i < 26; i++ {
			if node.children[i] != nil {
				hasChild = true
				dfs(node.children[i], depth + 1)
			}
		}
		if !hasChild && node.isEnd {
			res += depth
		}
	}
	dfs(t, 1)
	return res
}


type Trie struct {
	children [26]*Trie
	isEnd bool
}

func (t *Trie) Insert(word string) {
	node := t
	wb := []byte(word)
	for i := len(wb) - 1; i >= 0; i-- {
		ch := wb[i] - 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) bool {
	node := t
	pb := []byte(prefix)
	for i := len(pb) - 1; i >= 0; i-- {
		ch := pb[i] - 'a'
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}

	return node != nil
}
