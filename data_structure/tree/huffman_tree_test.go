/* ***** */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/5 3:57 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package tree

import (
	"fmt"
	"testing"
)

func TestNewHuffmanTree(t *testing.T) {
	arr := []int{8, 2, 3, 9, 10}
	ht := NewHuffmanTree(arr)
	ht.CreateTree()
	for _, item := range ht.nodes {
		fmt.Printf("ID:%d,weight:%d,parent:%d,lChild:%v, rChild:%v\n",
			item.id, item.weight, item.parent, item.lChild, item.rChild)
	}
}

func TestHuffmanTree_CreateCode(t *testing.T) {
	arr := []int{8, 2, 3, 9}
	chars := []byte{'A', 'B', 'C', 'D'}
	h := NewHuffmanTree(arr)
	h.CreateTree()
	h.CreateCode(chars)
	fmt.Println(h.char2CodeMap)
	fmt.Println(h.code2CharMap)
}


func TestHuffmanTree_Encode(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	chars := []byte{'i', 'l', 'o', 'v', 'e', 'u', 'y', ' '}
	h := NewHuffmanTree(arr)
	h.CreateTree()
	h.CreateCode(chars)
	s := h.Encode("i love you")
	expect := "111100111111111010010101001110110"
	if s != expect {
		t.Errorf("expect:%s, actual:%s\n", expect, s)
	}
}

func TestHuffmanTree_Decode(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	chars := []byte{'i', 'l', 'o', 'v', 'e', 'u', 'y', ' '}
	h := NewHuffmanTree(arr)
	h.CreateTree()
	h.CreateCode(chars)
	s := "111100111111111010010101001110110"
	result := h.Decode(s)
	expect := "i love you"
	if result != expect {
		t.Errorf("expect:%s, actual:%s\n", expect, result)
	}
}