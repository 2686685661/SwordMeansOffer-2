# 基础知识
《Go 语言设计与实现》：https://draveness.me/golang/  
go大汇集：http://www.topgoer.com/  
《channel的实现原理.pdf》

## go面试题
[Golang 常见面试题目解析](https://github.com/lifei6671/interview-go#golang-%E5%B8%B8%E8%A7%81%E9%9D%A2%E8%AF%95%E9%A2%98%E7%9B%AE%E8%A7%A3%E6%9E%90)


## GMP模型和并发调度
[Golang 理论](https://github.com/lifei6671/interview-go#golang-%E7%90%86%E8%AE%BA)
- Go语言的GPM调度器是什么？
- Goroutine调度策略
- goroutine调度器概述


## go内存分配和回收
《go的内存分配.pdf》
[Golang内存分配器](https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-memory-allocator/)  
[Golang 垃圾收集器](https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-garbage-collector/)  
[golang 栈内存管理](https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-stack-management/)  


## 性能分析和优化
[Go 语言中的大杀器](https://golang2.eddycjy.com/posts/ch6/01-pprof-1/)
[golang pprof 实战](https://blog.wolfogre.com/posts/go-ppof-practice/#%E5%89%8D%E8%A8%80)



## golang 设计模式
https://github.com/sevenelevenlee/go-patterns
https://github.com/pibigstar/go-demo/tree/master/design



## 流量控制算法
https://xie.infoq.cn/article/eca91dd0bd2455d7b3ea0f3b9
https://blog.csdn.net/xushiyu1996818/article/details/106764890/




[toc]
# go底层

## slice底层

> [Go语言中slice作为参数传递时遇到的问题有哪些]([Go语言中slice作为参数传递时遇到的问题有哪些 - 编程语言 - 亿速云 (yisu.com)](https://www.yisu.com/zixun/182790.html))

## map底层

>[解剖Go语言map底层实现 - Go语言中文网 - Golang中文社区 (studygolang.com)](https://studygolang.com/articles/14583)
>
>[Golang map 的底层实现 - 简书 (jianshu.com)](https://www.jianshu.com/p/aa0d4808cbb8)

go语言中的map采用的是**哈希查找表hashtable**，由一个key通过哈希函数得到哈希值，64位系统中就生成一个64bit的哈希值，由这个哈希值将key对应到不同的桶（bucket）中，当有多个哈希映射到相同的的桶中时，使用链表解决哈希冲突。

map无序的原因也是因为底层是通过hash来实现的

哈希表的特点是会有一个哈希函数，对你传来的key进行哈希运算，得到唯一的值，一般情况下都是一个数值。Golang的`map`中也有这么一个哈希函数，也会算出唯一的值，对于这个值的使用，Golang也是很有意思。

Golang把求得的值按照用途一分为二：高位和低位。

然后低位用于寻找当前key属于`hmap`中的哪个bucket，而高位用于寻找bucket中的哪个key。上文中提到：bucket中有个属性字段是“高位哈希值”数组，这里存的就是蓝色的高位值，用来声明当前bucket中有哪些“key”，便于搜索查找。 需要特别指出的一点是：我们`map`中的key/value值都是存到同一个数组中的。

---

## chan的底层实现

>[chan数据结构实现原理 - failymao - 博客园 (cnblogs.com)](https://www.cnblogs.com/failymao/p/14891813.html)

`src/runtime/chan.go:hchan`定义了channel的数据结构

```go
type hchan struct {
    qcount   uint           // 当前队列中剩余元素个数
    dataqsiz uint           // 环形队列长度，即可以存放的元素个数
    buf      unsafe.Pointer // 环形队列指针
    elemsize uint16         // 每个元素的大小
    closed   uint32         // 标识关闭状态
    elemtype *_type         // 元素类型
    sendx    uint           // 队列下标，指示元素写入时存放到队列中的位置
    recvx    uint           // 队列下标，指示元素从队列的该位置读出
    recvq    waitq          // 等待读消息的goroutine队列
    sendq    waitq          // 等待写消息的goroutine队列
    lock mutex              // 互斥锁，chan不允许并发读写
}
```

从数据结构可以看出channel由队列、类型信息、goroutine等待队列组成

1. `qcount`:当前队列中剩余元素的个数，我们在判断一个有缓存是否有可读数据时，**使用len(chan)**
2. `daraqsize`: 环形队列长度，可存放的元素个数，cap(chan)
3. `closed`: 标识关闭的状态，通过close(chan)来改变状态

### channel数据结构

#### 环形队列

chan内部实现了一个环形队列作为其缓冲区，队列的长度是创建chan时指定的： `make(chan int, 10)`

图展示了一个可缓存6个元素的channel示意图

[![bFqmrD.png](https://s4.ax1x.com/2022/02/24/bFqmrD.png)](https://imgtu.com/i/bFqmrD)

- dataqsiz指示了队列长度为6，即可缓存6个元素；
- buf指向队列的内存，队列中还剩余两个元素；
- qcount表示队列中还有两个元素；
- sendx指示后续写入的数据存储的位置，取值[0, 6)；
- recvx指示从该位置读取数据, 取值[0, 6)；

#### 等待队列

从channel读数据，如果channel缓冲区为空或者没有缓冲区，当前goroutine会被阻塞。
向channel写数据，如果channel缓冲区已满或者没有缓冲区，当前goroutine会被阻塞。

被阻塞的gouroutie将会挂在channel的等待队列中：

- 因读阻塞的goroutine会被像channel写入数据的goroutine唤醒
  - 常见有缓冲已经消费完数据数据后，如使用range读取时无数据时会进行阻塞
- 因写阻塞的goroutine会被从channel读数据的goroutine唤醒；
  - 无缓冲chan需要在写入前，读操作准备好
  - 有缓冲的chan在缓冲填满后，只有当读操作消费数据后，chan有空间，写入数据才能进行

下图展示了一个没有缓冲区的channel,有几个goroutine阻塞等待读数据：

[![bFqOdH.png](https://s4.ax1x.com/2022/02/24/bFqOdH.png)](https://imgtu.com/i/bFqOdH)

注意，一般情况下recvq和sendq至少有一个为空。只有一个例外，那就是同一个goroutine使用select语句向channel一边写数据，一边读数据。

#### 类型消息

- elemtype *_type

一个channel只能传递一种类型的值，类型信息存储在hchan数据结构中。

- elemtype代表类型，用于数据传递过程中的赋值；
- elemsize代表类型大小，用于在buf中定位元素位置。

#### 锁

一个channel同时仅允许被一个goroutine读写

### channel读写

#### 创建channel

创建channel的过程实际上是初始化hchan结构。其中类型信息和缓冲区长度由make语句传入，buf的大小则与元素大小和缓冲区长度共同决定。

#### 向channel中写入数据

向一个channel中写数据简单过程如下：

- 如果等待接收队列recvq不为空，说明缓冲区中没有数据或者没有缓冲区，此时直接从recvq取出G,并把数据写入，最后把该G唤醒，结束发送过程；
- 如果缓冲区中有空余位置，将数据写入缓冲区，结束发送过程；
- 如果缓冲区中没有空余位置，将待发送数据写入G，将当前G加入sendq，进入睡眠，等待被读goroutine唤醒；

#### 从channel中读数据

从一个channel读数据简单过程如下：

1. 如果等待发送队列sendq不为空，且没有缓冲区，直接从sendq中取出G，把G中数据读出，最后把G唤醒，结束读取过程；

2. 如果等待发送队列sendq不为空，此时说明缓冲区已满，从缓冲区中首部读出数据，把G中数据写入缓冲区尾部，把G唤醒，结束读取过程；

3. 如果缓冲区中有数据，则从缓冲区取出数据，结束读取过程；

4. 将当前goroutine加入recvq，进入睡眠，等待被写goroutine唤醒；

   [![bFOdCn.md.png](https://s4.ax1x.com/2022/02/24/bFOdCn.md.png)](https://imgtu.com/i/bFOdCn)

#### 关闭channel

关闭channel时会把recvq中的G全部唤醒， 本该写入G的数据位置为nil。把sendq中的G全部唤醒，但这些G会panic

- 简单说就是，向关闭的channel写数据会panic,继续读数据会读出nil

除此之外， panic出现的常见场景还有：

1. 关闭值为nil的通道channel
2. 关闭已经被关闭的channel
3. 像已经关闭的channel写数据



---

## 锁的底层原理

>[(205条消息) Go面试题(四)：锁的实现原理sync-mutex 篇_小道哥的博客-CSDN博客_golang sync.mutex 实现原理](https://blog.csdn.net/xiaodaoge_it/article/details/121435160)

在go中主要实现了两种锁：互斥锁(sync.Mutex)和读写锁(sync.RWMutex)

为什么需要锁？

在高并发的背景下，多个goroutine同时执行，可能会读写到同一块的内存，比如：

```go
var count int
var mu sync.Mutex

func func1() {
	for i := 0; i < 1000; i++ {
		go func() {
			count = count + 1
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(count)
}
```

**在Go中对于并发程序进行公共资源的访问的限制最常用的就是互斥锁（sync.mutex）的方式**

sync.mutex的常用方法有两个：

```go
Mutex.lock()用来获取锁
Mutex.Unlock()用于释放锁
```

在 Lock 和 Unlock 方法之间的代码段称为资源的临界区，这一区间的代码是严格被锁保护的，是线程安全的，任何一个时间点最多只能有一个goroutine在执行。
上面的代码加锁之后

```go
var count int
var mutex sync.Mutex

func func2() {
	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			count = count + 1
			mutex.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(count)
}
```

当某一goroutine执行了mutex.lock()方法后，如果有其他的goroutine来执行上锁操作，会被阻塞，直到当前的goroutine执行mutex.unlock()方法释放锁后其他的goroutine才会继续抢锁执行。

### 实现原理

**Sync.Mutex由两个字段构成，state用来表示当前互斥锁处于的状态，sema用于控制锁状态的信号量**。

**互斥锁state主要记录了如下四种状态**

[![HxQZFK.md.png](https://s4.ax1x.com/2022/02/21/HxQZFK.md.png)](https://imgtu.com/i/HxQZFK)

waiter_num： 记录了当前等待抢这个锁的goroutine数量

starving： 当前锁是否处于饥饿状态 (后文会详解锁的饥饿状态) 0: 正常状态 1: 饥饿状态

woken： 当前锁是否有goroutine已被唤醒。 0：没有goroutine被唤醒； 1: 有goroutine正在加锁过程

locked： 当前锁是否被goroutine持有。 0: 未被持有 1: 已被持有

sema信号量的作用：当持有锁的gorouine释放锁后，会释放sema信号量，这个信号量会唤醒之前抢锁阻塞的gorouine来获取锁。

### 锁的两种模式

**互斥锁在设计上主要有两种模式： 正常模式和饥饿模式**。

之所以引入了饥饿模式，是为了保证goroutine获取[互斥锁](https://so.csdn.net/so/search?q=互斥锁&spm=1001.2101.3001.7020)的公平性。所谓公平性，其实就是多个goroutine在获取锁时，goroutine获取锁的顺序，和请求锁的顺序一致，则为公平。

**正常模式下**，所有阻塞在等待队列中的goroutine会按顺序进行锁获取，当唤醒一个等待队列中的goroutine时，此goroutine并不会直接获取到锁，而是会和新请求锁的goroutine竞争。 通常新请求锁的goroutine更容易获取锁，这是因为新请求锁的goroutine正在占用cpu片执行，大概率可以直接执行到获取到锁的逻辑。
**饥饿模式下**， 新请求锁的goroutine不会进行锁获取，而是加入到队列尾部阻塞等待获取锁。

**饥饿模式的触发条件：** 

- 当一个goroutine等待锁的时间超过1ms时，互斥锁会切换到饥饿模式

**饥饿模式的取消条件**：

- 当获取到锁的这个goroutine是等待锁队列中的最后一个goroutine，互斥锁会切换到正常模式
- 当获取到锁的这个goroutine的等待时间在1ms之内，互斥锁会切换到正常模式

### 注意事项

1. 在一个goroutine中执行Lock()加锁成功后，不要再重复进行加锁，否则会panic。
2. 在Lock() 之前 执行Unlock()释放锁 会panic
3. 对于同一把锁，可以在一个goroutine中执行Lock加锁成功后，可以在另外一个gorouine中执行Unlock释放锁。

### 读写锁的特点

读写锁区别与互斥锁的主要区别就是读锁之间是共享的，多个goroutine可以同时加读锁，但是写锁与写锁、写锁与读锁之间则是互斥的

### golang的读写锁的实现

[![Hx3XX4.md.png](https://s4.ax1x.com/2022/02/22/Hx3XX4.md.png)](https://imgtu.com/i/Hx3XX4)



## GMP模型

## 垃圾回收

> [Go语言内存管理三部曲（三）图解GC算法和垃圾回收原理 - InfoQ 写作平台](https://xie.infoq.cn/article/67cfd494e6e10cd0b40de95ab)
>
> [Go 垃圾回收（三）——三色标记法是什么鬼？ - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/105495961)

### 1.5之前采用**Mark-And-Sweep**（标记清除）

 **Mark-And-Sweep**，这个算法就是严格按照追踪式算法的思路来实现的。这个垃圾回收算法的执行流程可以用下面这

### 1.5 之后采用三色标记的并发垃圾收集器

三色标记清除算法背后的首要原则就是它把堆中的对象根据它们的颜色分到不同集合里面。

- 白色对象 — 潜在的垃圾，其内存可能会被垃圾收集器回收；
- 黑色对象 — 活跃的对象，包括不存在任何引用外部指针的对象以及从根对象可达的对象，垃圾回收器不会扫描这些对象的子对象；
- 灰色对象 — 活跃的对象，因为存在指向白色对象的外部指针，垃圾收集器会扫描这些对象的子对象；

相比传统的标记清扫算法，三色标记最大的好处是可以异步执行，从而可以以中断时间极少的代价或者完全没有中断来进行整个 GC。

三色标记法很简单[[2\]](https://zhuanlan.zhihu.com/p/105495961#ref_2)。首先将对象用三种颜色表示，分别是白色、灰色和黑色。最开始所有对象都是白色的，然后把其中**全局变量**和**函数栈里**的对象置为灰色。第二步把灰色的对象全部置为黑色，然后把原先灰色对象指向的变量都置为灰色，以此类推。等发现没有对象可以被置为灰色时，所有的白色变量就一定是需要被清理的垃圾了。

### 1.8后采用混合屏障





