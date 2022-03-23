/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 11:04 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package abstract

import "testing"

func TestAbstractFactory(t *testing.T) {

	store := new(GirlFactoryStore)
	// 提供美国工厂
	store.factory = new(AmericanGirlFactory)
	americanFatGirl := store.createGirl("fat")
	americanFatGirl.weight()

	// 提供中国工厂
	store.factory = new(ChineseGirlFactory)
	chineseFatGirl := store.createGirl("fat")
	chineseFatGirl.weight()
}