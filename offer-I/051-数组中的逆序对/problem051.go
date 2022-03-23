/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/11/8 2:00 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _51_数组中的逆序对


func InversePairs( data []int ) int {
	if data == nil || len(data) == 0 {
		return 0
	}
	copies := make([]int, len(data))
	for i := 0; i< len(data); i++ {
		copies[i] = data[i]
	}

	count := InversePairsCore(data, copies, 0, len(data) - 1)

	return count % 1000000007

}

func InversePairsCore(data, copy []int, low, high int) int {
	if low == high {
		return 0
	}

	mid := (low + high) >> 1
	leftCount := InversePairsCore(data, copy, low, mid)
	rightCount := InversePairsCore(data, copy, mid + 1, high)

	i, j, loc, count := mid, high, high, 0

	for ;i >= low && j > mid; {
		if data[i] > data[j] {
			count += j - mid
			copy[loc] = data[i]
			loc--
			i--
		} else {
			copy[loc] = data[j]
			loc--
			j--
		}
	}

	for ; i >= low; i-- {
		copy[loc] = data[i]
		loc--
	}
	for ; j > mid; j-- {
		copy[loc] = data[j]
		loc--
	}

	for s := low; s <= high; s++ {
		data[s] = copy[s]
	}
	return leftCount + rightCount + count

}