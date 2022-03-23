/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/16 7:21 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _59_滑动窗口的最大值

func MaxInWindows( num []int ,  size int ) []int {

	var (
		queue []int
		res []int
	)

	for i := 0; i < len(num); i++ {
		for len(queue) > 0 && num[queue[len(queue)-1]] <= num[i] {
			queue = queue[:len(queue)-1]
		}
		for len(queue) > 0 && i - queue[0] + 1 > size {
			queue = queue[1:]

		}

		queue = append(queue, i)
		if size > 0 && i + 1 >= size {
			res = append(res, num[queue[0]])
		}
	}
	return res
}