/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/30 3:32 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _31_最近最少使用缓存

import "container/list"

type LRUCache struct {
	capacity int
	list *list.List
	dict map[int]*list.Element
}


func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		dict:     make(map[int]*list.Element, capacity),
	}
}

type entry struct {
	key, val int
}

func (this *LRUCache) Get(key int) int {
	if elem, ok := this.dict[key]; ok {
		this.list.MoveToFront(elem)
		return elem.Value.(*entry).val
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if elem, ok := this.dict[key]; ok {
		e := elem.Value.(*entry)
		e.val = value
		this.list.MoveToFront(elem)
		return
	}
	//// 长度超限
	//if this.list.Len() >= this.capacity {
	//	delete(this.dict, this.list.Back().Value.(int))
	//	this.list.Remove(this.list.Back())
	//}

	this.dict[key] = this.list.PushFront(&entry{key, value})
	if this.capacity < len(this.dict) {
		delete(this.dict, this.list.Remove(this.list.Back()).(*entry).key)
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */