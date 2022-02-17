/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 10:11 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _5_观察者模式_observer_

import "fmt"


// 报社 -- 客户
type Customer interface {
	update()
}

type customerA struct {}

func (*customerA) update() {
	fmt.Println("客户A，收到了报纸")
}

type customerB struct {}

func (*customerB) update() {
	fmt.Println("客户B，收到了报纸")
}

// 报社 （被观察者)
type newsOffice struct {
	customers []Customer
}

func (n *newsOffice) addCustomer(customer Customer) {
	// 通知所有客户
	n.customers = append(n.customers, customer)
}

func (n *newsOffice) newsPaperCome() {
	n.notifyAllCustomer()
}

func (n *newsOffice) notifyAllCustomer() {
	for _, customer := range n.customers {
		customer.update()
	}
}
