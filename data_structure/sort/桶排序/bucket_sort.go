/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/9 7:22 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package 桶排序

// 桶排序
// 非比较类算法
func BucketSort(elem []int) []int {
	if len(elem) < 1 {
		return elem
	}

	minVal, maxVal := elem[0], elem[0]
	for i := 0; i < len(elem); i++ {
		if minVal > elem[i] {
			minVal = elem[i]
		}
		if maxVal < elem[i] {
			maxVal = elem[i]
		}
	}

	//初始化桶
	bucketSize := (maxVal - minVal) / len(elem) + 1
	buckets := make([][]int, bucketSize)

	//数据入桶
	for i := 0; i < len(elem);i++ {
		buckets[(elem[i] - minVal) / len(elem)] = append(buckets[(elem[i] - minVal) / len(elem)], elem[i])
	}

	// 每个桶内排序，这里采用了直插
	for i, bucket := range buckets {
		if len(bucket) < 1 {
			continue
		}
		buckets[i] = insertSort(bucket)
	}

	// 出桶
	i := 0
	for _, b := range buckets {
		for _, v := range b {
			elem[i] = v
			i++
		}
	}
	return elem
}

func insertSort(elem []int) []int {
	for i := 1; i < len(elem); i++ {
		if elem[i] < elem[i - 1] {
			j, tmp := 0, elem[i]
			for j = i - 1; j >= 0 && elem[j] > tmp;j-- {
				elem[j + 1] = elem[j]
			}
			elem[j + 1] = tmp
		}
	}
	return elem
}