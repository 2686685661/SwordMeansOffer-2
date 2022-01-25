/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/20 11:05 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _64_神奇的字典

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

func (t * Trie) Search(word string) bool {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return node != nil && node.isEnd
}



type MagicDictionary struct {
	trie *Trie
}


/** Initialize your data structure here. */
func Constructor() MagicDictionary {
	return MagicDictionary{trie: &Trie{}}
}


func (m *MagicDictionary) BuildDict(dictionary []string)  {
	for _, word := range dictionary {
		m.trie.Insert(word)
	}
}


func (m *MagicDictionary) Search(searchWord string) bool {
	node := m.trie
	sba := []byte(searchWord)
	for i, ch := range sba {
		for j := 0; j < 26; j++ {
			if byte('a' + j) == ch {
				continue
			}
			sba[i] = byte('a' + j)
			if node.Search(string(sba[i:])) {
				return true
			}
		}
		sba[i] = ch
		if node.children[ch - 'a'] == nil {
			return false
		}
		node = node.children[ch - 'a']
	}
	return false
}


/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */