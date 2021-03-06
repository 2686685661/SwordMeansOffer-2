/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/19 2:49 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package trie


// 字典树(前缀树)
type Trie struct {
	children [26]*Trie
	isEnd bool
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string)  {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}


/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil

}


func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}