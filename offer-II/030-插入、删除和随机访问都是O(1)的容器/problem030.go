/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/30 2:18 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _30_插入_删除和随机访问都是_O_1__的容器

import "math/rand"

type RandomizedSet struct {
	dict map[int]int
	list []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	m := make(map[int]int)
	var l []int
	return RandomizedSet{
		m,
		l,
	}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dict[val]; ok {
		return false
	}
	this.dict[val] = len(this.list)
	this.list = append(this.list, val)
	return true

}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if i, ok := this.dict[val]; !ok {
		return false
	} else {
		li := len(this.list) - 1
		lv := this.list[li]
		this.list[i] = this.list[li]
		this.dict[lv] = i
		this.list = this.list[:li]
		delete(this.dict, val)
	}
	return true
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */