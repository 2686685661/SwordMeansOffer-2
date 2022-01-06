/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/12/30 6:26 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _35_最小时间差

import (
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	//时间点最多只有 24 * 60 个，因此，当 timePoints 长度超过 24 * 60，说明有重复的时间点，提前返回 0。
	if len(timePoints) < 2 || len(timePoints) > 24 * 60 {
		return 0
	}
	var mins = make([]int, 0)

	for _, v := range timePoints {
		time := strings.Split(v, ":")
		h, _ := strconv.Atoi(time[0])
		m, _ := strconv.Atoi(time[1])
		mins = append(mins, h * 60 + m)
	}
	//将“分钟制”列表按升序排列，然后将此列表的最小时间 mins[0] 加上 24 * 60 追加至列表尾部，用于处理最大值、最小值的差值这种特殊情况
	sort.Ints(mins)
	mins = append(mins, mins[0] + 24 * 60)
	res := 24 * 60
	for i := 1; i < len(mins); i++ {
		if mins[i] - mins[i - 1]  < res {
			res = mins[i] - mins[i - 1]
		}
	}
	return res
}