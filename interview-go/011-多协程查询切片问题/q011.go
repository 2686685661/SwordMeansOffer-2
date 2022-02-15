/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/9 9:43 下午, by lishanlei, create
*/

/*
假设有一个超长的切片，切片的元素类型为int，切片中的元素为乱序排序。限时5秒，使用多个goroutine查找切片中是否存在给定的值，在查找到目标值或者超时后立刻结束所有goroutine的执行。
比如，切片 [23,32,78,43,76,65,345,762,......915,86]，查找目标值为 345 ，如果切片中存在，则目标值输出"Found it!"并立即取消仍在执行查询任务的goroutine。
如果在超时时间未查到目标值程序，则输出"Timeout！Not Found"，同时立即取消仍在执行的查找任务的goroutine。

答案: https://mp.weixin.qq.com/s/GhC2WDw3VHP91DrrFVCnag
*/

package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 5)
	data := []int{1, 2, 3, 10, 999, 8, 345, 7, 98, 33, 66, 77, 88, 68, 96}
	dl := len(data)
	size := 3
	target := 345
	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan bool)

	SearchTarget := func(ctx context.Context, data []int, target int, resultChan chan bool) {
		for _, v := range data {
			select {
			case <-ctx.Done():
				_, _ = fmt.Fprintf(os.Stdout, "Task cancelded! \n")
				return
			default:
			}
			// 模拟一个耗时查找，这里只是比对值，真实开发中可以是其他操作
			_, _ = fmt.Fprintf(os.Stdout, "v: %d \n", v)
			time.Sleep(time.Millisecond * 1500)
			if target == v {
				resultChan <- true
				return
			}
		}
	}

	for i := 0; i < dl; i += size {
		end := i + size
		if end >= dl {
			end = dl - 1
		}
		go SearchTarget(ctx, data[i:end], target, resultChan)
	}

	select {
	case <-timer.C:
		_, _ = fmt.Fprintln(os.Stderr, "Timeout! Not Found")
		cancel()
	case <-resultChan:
		_, _ = fmt.Fprintf(os.Stdout, "Found it!\n")
		cancel()
	}
	time.Sleep(time.Second * 2)

}