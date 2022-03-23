/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 8:54 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _3_适配器模式_adaptor_

import "testing"

func TestAdaptor(t *testing.T) {
	player := PlayerAdaptor{}
	player.play("mp3", "死了都要爱")
	player.play("wma", "滴滴")
	player.play("mp4", "复仇者联盟")
}