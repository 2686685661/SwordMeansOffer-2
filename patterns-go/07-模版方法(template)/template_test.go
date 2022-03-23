/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/18 12:13 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _7_模版方法_template_

import "testing"

func TestTemplate(t *testing.T) {

	// 做西红柿
	xihongshi := &XiHongShi{}
	doCook(xihongshi)

	// 做炒鸡蛋
	chaojidan := &ChaoJiDan{}
	doCook(chaojidan)

}