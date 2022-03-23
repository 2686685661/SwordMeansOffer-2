/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 8:46 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _3_适配器模式_adaptor_

import "fmt"

type MusicPlayer interface {
	play(fileType, fileName string)
}


type ExistPlayer struct {}

func (*ExistPlayer) playMp3(fileName string) {
	fmt.Println("paly mp3: ", fileName)
}

func (*ExistPlayer) playWma(fileName string) {
	fmt.Println("play wma: ", fileName)
}


type PlayerAdaptor struct {
	existPlayer ExistPlayer
}

func (player PlayerAdaptor) play(fileType, fileName string) {
	switch fileType {
	case "mp3":
		player.existPlayer.playMp3(fileName)
	case "wma":
		player.existPlayer.playWma(fileName)
	default:
		fmt.Println("不支持此类型文件播放")
	}
}