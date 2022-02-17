/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/18 12:26 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _9_代理模式_proxy_

import "testing"

func TestProxy(t *testing.T) {

	station := &Station{3}
	proxy := &StationProxy{station}
	station.sell("小华")
	proxy.sell("派大星")
	proxy.sell("小明")
	proxy.sell("小兰")
}
