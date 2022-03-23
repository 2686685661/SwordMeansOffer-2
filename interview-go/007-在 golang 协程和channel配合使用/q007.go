/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/9 4:55 下午, by lishanlei, create
*/

/*
题目
写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。

解析
这是一道很简单的golang基础题目，实现方法也有很多种，一般想答让面试官满意的答案还是有几点注意的地方。
- goroutine 在golang中式非阻塞的
- channel 无缓冲情况下，读写都是阻塞的，且可以用for循环来读取数据，当管道关闭后，for 退出。
- golang 中有专用的select case 语法从管道读取数据。
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	out := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			out <- rand.Intn(5)
		}
		close(out)
	}()
	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}

