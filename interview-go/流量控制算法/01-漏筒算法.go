/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/28 10:43 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

/**
参考：https://blog.csdn.net/liyunlong41/article/details/88429183
 */


type rateLimiter struct {
	m *sync.Mutex // 防止并发情况下出现错误
	rate float64 // 最大速率限制，即接口每秒限制多少个请求。也就是水滴从漏桶中流出的速度，同时也是余量增加的速度
	balance float64 // 漏桶的空闲余量，会随着漏桶滴水逐渐变大；如果将请求添加到漏桶中，会逐渐变小。当请求到来时，如果余量不足1，那么表明不能容下当前的请求，当前的请求会被拒绝。
	limit float64 // 漏桶的最大容量
	lastTime time.Time // 上次调用Check函数的时间。用于计算时间差dur，然后计算这段时间漏桶流出的水w，增加的余量=流出的水量w=时间*速率=dur*rate。
}

func NewRateLimiter(limitPerSecond int, balance int) *rateLimiter {
	return &rateLimiter{
		m:        new(sync.Mutex),
		rate:     float64(limitPerSecond),
		balance:  float64(balance),
		limit:    float64(balance),
		lastTime: time.Now(),
	}
}

func (r *rateLimiter) Check() (ok bool) {
	r.m.Lock()
	defer r.m.Unlock()
	now := time.Now()
	dur := now.Sub(r.lastTime).Seconds()
	r.lastTime = now
	water := dur * r.rate //计算这段时间内漏桶流出水的流量water
	r.balance += water //漏桶流出water容量的水，自然漏桶的余量多出water
	if r.balance > r.limit {
		r.balance = r.limit
	}
	//漏桶余量足够容下当前的请求
	if r.balance >= 1 {
		r.balance -= 1
		ok = true
	}
	return
}

func main() {
	limiter := NewRateLimiter(10, 1)
	start := time.Now()
	count := 0
	for i := 0; i < 1e3; i++ {
		if limiter.Check() {
			fmt.Println(i)
			count++
		}
		time.Sleep(time.Millisecond)
	}
	fmt.Println("count: ", count)
	fmt.Println(time.Now().Sub(start).Seconds())

}