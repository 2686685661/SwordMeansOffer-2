/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/9 6:13 下午, by lishanlei, create
*/

/*
写出以下逻辑，要求每秒钟调用一次proc并保证程序不退出?
```
package main

func main() {
    go func() {
        // 1 在这里需要你写算法
        // 2 要求每秒钟调用一次proc函数
        // 3 要求程序不能退出
    }()

    select {}
}

func proc() {
    panic("ok")
}
```

解析
题目主要考察了两个知识点：
1.定时执行执行任务
2.捕获 panic 错误

题目中要求每秒钟执行一次，首先想到的就是 time.Ticker对象，该函数可每秒钟往chan中放一个Time,正好符合我们的要求。
在 golang 中捕获 panic 一般会用到 recover() 函数。
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出

		t := time.NewTicker(time.Second * 1)

		for {
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					proc()
				}()
			}
		}
	}()
	select {}
}

func proc() {
	panic("ok")
}

