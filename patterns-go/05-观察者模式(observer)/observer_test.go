/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 10:15 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _5_观察者模式_observer_

import "testing"

func TestObserver(t *testing.T) {
	customerA := &customerA{}
	customerB := &customerB{}

	office := &newsOffice{}
	// 模拟客户订阅
	office.addCustomer(customerA)
	office.addCustomer(customerB)

	// 新的报纸
	office.newsPaperCome()
}