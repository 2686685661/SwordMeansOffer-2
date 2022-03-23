/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/7/19 9:23 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package stack

import "testing"

func TestItemStack_Push(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if size := len(s.items); size != 3{
		t.Errorf("test failed ")
	}
}

func TestItemStack_Pop(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	i := s.Pop()
	if size := len(s.items); size != 2 || *i != 3 {
		t.Errorf("test failed, excepted 2 and got %d", size)
	}

	i = s.Pop()
	if size := len(s.items); size != 1 || *i != 2 {
		t.Errorf("test failed, excepted 2 and got %d", size)
	}
	s.Pop()

	if size := len(s.items); size != 0{
		t.Errorf("test failed, excepted 0 and got %d", size)
	}
}