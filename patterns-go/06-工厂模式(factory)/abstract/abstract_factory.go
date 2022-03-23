/* *****. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 10:58 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package abstract

import "fmt"

// 女孩
type girl interface {
	weight()
}


// 中国胖女孩
type fatGirl struct {}

func (*fatGirl) weight() {
	fmt.Println("chinese girl, weight:80kg")
}

// 中国瘦女孩
type thinGirl struct {}

func (*thinGirl) weight() {
	fmt.Println("chinese girl, weight:50kg")
}


// 美国胖女孩
type AmericanFatGirl struct {}

func (AmericanFatGirl) weight() {
	fmt.Println("American girl, weight: 80kg")
}

// 美国瘦女孩
type AmericanThainGirl struct {
}

func (AmericanThainGirl) weight() {
	fmt.Println("American girl, weight: 50kg")
}



// 抽象工厂接口
type factory interface {
	createGirl(like string) girl
}

// 中国工厂
type ChineseGirlFactory struct {}
func (ChineseGirlFactory) createGirl(like string) girl {
	if like == "fat" {
		return &fatGirl{}
	} else if like == "thin" {
		return &thinGirl{}
	}
	return nil
}

// 美国工厂
type AmericanGirlFactory struct {}
func (AmericanGirlFactory) createGirl(like string) girl {
	if like == "fat" {
		return &AmericanFatGirl{}
	} else if like == "thin" {
		return &AmericanThainGirl{}
	}
	return nil
}


// 工厂提供者
type GirlFactoryStore struct {
	factory factory
}
func (store *GirlFactoryStore) createGirl(like string) girl {
	return store.factory.createGirl(like)
}