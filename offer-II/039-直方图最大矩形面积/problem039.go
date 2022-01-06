/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/1/6 7:29 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _39_直方图最大矩形面积

/**
参考：https://leetcode-cn.com/problems/0ynMMM/solution/jian-zhi-offer-2-mian-shi-ti-39-shu-zhon-qzaw/
 */

// 分治法
func largestRectangleArea(heights []int) int {
	if heights == nil {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}
	return helper(heights, 0, len(heights))
}

func helper(heights []int, start, end int) int {
	if start == end {
		return 0
	}

	if start + 1 == end {
		return heights[start]
	}

	mh, mi := heights[start], start
	for i := start + 1; i < end; i++ {
		if heights[i] < mh {
			mh, mi = heights[i], i
		}
	}

	midPlane := (end - start) * mh
	leftPlane, rightPlane := helper(heights, start, mi), helper(heights, mi + 1, end)
	return Max(midPlane, Max(leftPlane, rightPlane))

}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*-----------------------------------------------------------------*/

// 单调栈
type Stack []int
func (s Stack) Len() int {
	return len(s)
}
func (s Stack) Cap() int {
	return cap(s)
}

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

func (s Stack) Top() int {
	if len(s) == 0 {
		return 0
	}
	return s[len(s) - 1]
}

func (s *Stack) Pop() int {
	ts := *s
	if len(ts) == 0 {
		return 0
	}
	v := ts[len(ts) - 1]
	*s = ts[:len(ts) - 1]
	return v
}

func (s Stack) IsEmpty() bool  {
	return len(s) == 0
}

func largestRectangleArea2(heights []int) int {
	if heights == nil {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}

	stack, maxArea := new(Stack), 0

	stack.Push(-1)

	for i, v := range heights {
		for stack.Top() != -1 && heights[stack.Top()] >= v {
			h := heights[stack.Pop()]
			w := i - stack.Top() - 1
			maxArea = Max(maxArea, w * h)
		}
		stack.Push(i)
	}

	for stack.Top() != -1 {
		h := heights[stack.Pop()]
		w := len(heights) - stack.Top() - 1
		maxArea = Max(maxArea, w * h)
	}
	return maxArea

}