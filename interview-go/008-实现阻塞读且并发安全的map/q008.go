/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/9 5:20 下午, by lishanlei, create
*/

/*
GO里面MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全

解析：
看到阻塞协程第一个想到的就是channel，题目中要求并发安全，那么必须用锁，还要实现多个goroutine读的时候如果值不存在则阻塞，直到写入值，那么每个键值需要有一个阻塞goroutine 的 channel。
*/

package main

import (
	"log"
	"sync"
	"time"
)

type sp interface {
	Out(key string, val interface{})  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type Map struct {
	c map[string]*entry
	rmx *sync.RWMutex
}

type entry struct {
	ch chan struct{}
	value interface{}
	isExist bool
}

func (m *Map) Out(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	item, ok := m.c[key]
	if !ok {
		m.c[key] = &entry{
			value:   val,
			isExist: true,
		}
		return
	}
	item.value = val
	if !item.isExist && item.ch != nil {
		close(item.ch)
		item.ch = nil
	}
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	m.rmx.RLock()
	if item, ok := m.c[key]; ok && item.isExist {
		m.rmx.RUnlock()
		return item.value
	} else if !ok {
		m.rmx.RUnlock()
		m.rmx.Lock()
		item = &entry{
			ch:      make(chan struct{}),
			isExist: false,
		}
		m.c[key] = item
		m.rmx.Unlock()
		select {
		case <- item.ch:
			return item.value
		case <-time.After(timeout):
			return nil
		}
	} else {
		m.rmx.RUnlock()
		select {
		case <- item.ch:
			return item.value
		case <-time.After(timeout):
			return nil
		}
	}
}


func main() {
	m := Map{
		c:   make(map[string]*entry),
		rmx: &sync.RWMutex{},
	}
	for i := 0; i < 10; i++ {
		go func() {
			val := m.Rd("key", time.Second * 6)
			log.Println("读取值为->", val)
		}()
	}

	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		go func(val int) {
			m.Out("key", i)
		}(i)
	}
	time.Sleep(time.Second * 30)
}

