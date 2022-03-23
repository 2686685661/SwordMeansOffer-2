/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/7/19 9:19 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package stack


type Item interface {

}
type ItemStack struct {
	items []Item
}

func New() *ItemStack {
	s := &ItemStack{items: []Item{}}
	return s
}

func (s *ItemStack) Push(t Item) {
	s.items = append(s.items, t)
}

func (s *ItemStack) Pop() *Item {
	i := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items) - 1]
	return &i
}
