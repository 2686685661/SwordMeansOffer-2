/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/7 10:59 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _42_最近请求次数


type RecentCounter struct {
	slot int
	queue []int
}


func Constructor() RecentCounter {
	return RecentCounter{
		slot:  3000,
		queue: make([]int, 0),
	}
}


func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	early := t - this.slot
	for i := 0; i < len(this.queue); i++ {
		if this.queue[i] < early {
			this.queue = append(this.queue[:i], this.queue[i+1:]...)
			i--
		}
	}
	return len(this.queue)

}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */