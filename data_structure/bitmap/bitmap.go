/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/3/18 12:07 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package bitmap

import "fmt"

/**
BitMap 算法实现
https://www.jianshu.com/p/6082a2f7df8e
 */

/**

1byte=8bit
n >> 3 就是 n / 8，获得数字n在bits数组中的下标索引
n & 0x07 就是 n % 8，获得n在指定bit上的位置下标索引
add：将1左移position后，那个位置自然就是1，然后和以前的数据做|，这样，那个位置就替换成1了
clear：将1左移position后，进行位取反，则指定位置为0，其他位置是1，然后和以前的数据做&，这样指定位置就是1&0=0了
Contain：将1左移position后，指定位置是1，然后和原数据进行&，如果指定位置有值，则1&1=1，如果指定位置无值，则1&0=0，即可判断是否存在
*/
type BitMap struct {
	bits []byte
	cap int

}

func NewBitMap(cap int) *BitMap {
	return &BitMap{
		bits:make([]byte, cap >> 3 + 1),
		cap : cap,
	}
}


func (b *BitMap) Add(n int) {
	if n > b.cap {
		fmt.Printf("int=%d out of range!\n", n)
		return
	}
	b.bits[b.index(n)] |= 1 << b.position(n)
}

func (b *BitMap) Clear(n int) {
	b.bits[b.index(n)] &= ^(1 << b.position(n))
}

func (b *BitMap) Contain(n int) bool {
	return (b.bits[b.index(n)] & (1 << b.position(n))) != 0
}

func (b BitMap) index(n int) int {
	return n >> 3
}
func (b BitMap) position(n int) int {
	return n & 0x07
}
