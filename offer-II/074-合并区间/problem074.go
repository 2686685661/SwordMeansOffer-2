/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/26 4:27 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _74_合并区间

import "sort"

/**
首先，我们将列表中的区间按照左端点升序排序。然后我们将第一个区间加入 merged 数组中，并按顺序依次考虑之后的每个区间：
如果当前区间的左端点在数组 merged 中最后一个区间的右端点之后，那么它们不会重合，我们可以直接将这个区间加入数组 merged 的末尾；
否则，它们重合，我们需要用当前区间的右端点更新数组 merged 中最后一个区间的右端点，将其置为二者的较大值。
 */
func merge(intervals [][]int) [][]int {
	if intervals == nil || len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := make([][]int, 0)
	Max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > merged[len(merged) - 1][1] {
			merged = append(merged, intervals[i])
		} else {
			merged[len(merged) - 1][1] = Max(merged[len(merged) - 1][1], intervals[i][1])
		}
	}
	return merged

}