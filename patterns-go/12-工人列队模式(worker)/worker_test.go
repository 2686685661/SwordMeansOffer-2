/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/18 12:59 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _2_工人列队模式_worker_

import (
	"fmt"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	go makeJob()

	dispatcher := NewDispatcher(10)
	dispatcher.Run()

	time.Sleep(time.Second * 1)
}

func makeJob() {
	for i := 0; i < 100; i++ {
		job := Job{Content: fmt.Sprintf("第%d个任务", i)}
		JobQueue <- job
	}
}