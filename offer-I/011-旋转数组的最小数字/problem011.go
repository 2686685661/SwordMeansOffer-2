/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/10/18 4:04 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _11_旋转数组的最小数字

// 二分查找思路
/**
需要考虑三种情况：
	(1)array[mid] > array[high]:
	出现这种情况的array类似[3,4,5,6,0,1,2]，此时最小数字一定在mid的右边。
	low = mid + 1

	(2)array[mid] == array[high]:
	出现这种情况的array类似 [1,0,1,1,1] 或者[1,1,1,0,1]，此时最小数字不好判断在mid左边
	还是右边,这时只好一个一个试 ，
	high = high - 1

	(3)array[mid] < array[high]:
	出现这种情况的array类似[2,2,3,4,5,6,6],此时最小数字一定就是array[mid]或者在mid的左
	边。因为右边必然都是递增的。
	high = mid
 */
func MinNumberInRotateArray( rotateArray []int ) int {
	if len(rotateArray) == 0 {
		return 0
	}

	l, r := 0, len(rotateArray) - 1

	for l < r {
		m := (l + r) / 2
		if rotateArray[m] > rotateArray[r] {
			l = m + 1
		} else if rotateArray[m] < rotateArray[r] {
			r = m
		} else if rotateArray[m] == rotateArray[r] {
			r = r - 1
		}
	}
	return rotateArray[l]
}