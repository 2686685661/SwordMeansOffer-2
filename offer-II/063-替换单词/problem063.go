/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/19 2:24 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _63_替换单词

import "strings"

// 字典树(前缀树)
type Trie struct {
	children [26]*Trie
	isEnd bool
}

func (t *Trie) Insert(word string) {
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

// 查询是否你存在指定字符串的前缀
func (t *Trie) Search(word string) bool {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.isEnd {
			break
		}
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return true
}

func (t *Trie) Replace(word string) string {
	node := t
	var res []byte
	for _, ch := range word {
		if node.isEnd || node.children[ch - 'a'] == nil {
			break
		}
		node = node.children[ch - 'a']
		res = append(res, byte(ch))
	}
	return string(res)
}


func replaceWords(dictionary []string, sentence string) string {

	t := &Trie{}

	for _, s := range dictionary {
		t.Insert(s)
	}
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if t.Search(word) {
			words[i] = t.Replace(word)
		}
	}

	return strings.Join(words, " ")

}