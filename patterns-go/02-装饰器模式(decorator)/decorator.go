/* Copyright 2022 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2022/2/17 1:09 上午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package _2_装饰器模式_decorator_

import "fmt"

type Person interface {
	cost() int
	show()
}

type laowang struct {}

func (*laowang) cost() int {
	return 0
}

func (*laowang) show() {
	fmt.Println("光溜溜的老王")
}


type clothesDecorator struct {
	person Person
}

func (*clothesDecorator) cost() int {
	return 0
}
func (*clothesDecorator) show() {}

type Jacket struct {
	clothesDecorator
}

func (j *Jacket) cost() int {
	return j.person.cost() + 10
}

func (j *Jacket) show() {
	j.person.show()
	fmt.Println("穿上夹克的老王")
}

type Hat struct {
	clothesDecorator
}

func (h *Hat) cost() int {
	return h.person.cost() + 5
}

func (h *Hat) show() {
	fmt.Println("戴上帽子的老王")
}






