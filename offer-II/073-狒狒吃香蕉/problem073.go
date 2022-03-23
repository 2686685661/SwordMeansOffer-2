/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/25 1:51 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _73_狒狒吃香蕉

/**
二分查找
这道题典型的二分常考的【二分答案】，答案一定在某个区间内，所以可以二分检查快速接近正确答案
 */
func minEatingSpeed(piles []int, h int) int {
	if piles == nil || len(piles) < 1 || h < len(piles) {
		return 0
	}

	// 狒狒的速度肯定需要大于等于 1，同时也要小于等于最大的一堆香蕉数量 max，因为若大于 max 每小时也只能吃一堆，所以更大的速度是没有意义的
	low, high := 1, 0
	for _, p := range piles {
		if p > high {
			high = p
		}
	}

	// 返回狒狒速度k时需要吃完香蕉的时间t
	countTime := func(k int) (t int) {
		for _, p := range piles {
			t += p / k
			if p % k > 0 {
				t += 1
			}
		}
		return
	}

	for low <= high {
		mid := (low + high) >> 1
		if countTime(mid) <= h {
			// 如果时间t <= H，还需要判断k=mid-1速度吃完香蕉的时间t2是否大于H，如果t2大于H，说明k=mid为最慢吃完香蕉所需要的时间; 如果t2小于H，说明当前吃的速度太快，需要降低吃的速度
			if mid == 1 || countTime(mid - 1) > h {
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1 // t > H，说明狒狒吃的太慢了，得加快速度
		}
	}
	return -1
}