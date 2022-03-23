/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/6 7:55 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _41_滑动窗口的平均值


type MovingAverage struct {
	size int
	sum int
	nums []int
}


/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		size: size,
		sum:  0,
		nums: make([]int, 0),
	}

}


func (this *MovingAverage) Next(val int) float64 {
	this.sum += val
	this.nums = append(this.nums, val)
	for len(this.nums) > this.size {
		r := this.nums[0]
		this.sum -= r
		this.nums = this.nums[1:]
	}
	return float64(this.sum) / float64(len(this.nums))

}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */