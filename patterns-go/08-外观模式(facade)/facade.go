/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/18 12:17 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _8_外观模式_facade_

import "fmt"

// CPU
type CPU struct {
}

func (CPU) start() {
	fmt.Println("启动CPU。。。")
}

// 内存
type Memory struct {
}

func (Memory) start() {
	fmt.Println("启动内存管理。。。")
}

// 硬盘
type Disk struct {
}

func (Disk) start() {
	fmt.Println("启动硬盘。。。")
}

// 开机键
type StartBtn struct {
}

func (StartBtn) start() {
	cpu := &CPU{}
	cpu.start()
	memory := &Memory{}
	memory.start()
	disk := &Disk{}
	disk.start()
}