/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/24 3:12 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _66_单词之和



type MapSum struct {
	children [26]*MapSum
	isEnd bool
	val int
}



/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{}

}


func (ms *MapSum) Insert(key string, val int)  {
	node := ms
	for _, ch := range key {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &MapSum{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
	node.val = val
}


func (ms *MapSum) Sum(prefix string) int {
	node :=ms
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return 0
		}
		node = node.children[ch]
	}
	if node == nil {
		return 0
	}

	var res int
	var dfs func(node *MapSum)
	dfs = func(node *MapSum) {
		for i := 0; i < 26; i++ {
			if node.children[i] != nil {
				dfs(node.children[i])
			}
		}
		if node != nil && node.isEnd {
			res += node.val
		}
	}

	dfs(node)
	return res
}



/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */