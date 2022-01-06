/* Copyright 2021 Baidu Inc. All Rights Reserved. */
/* - please input the go file action-  */
/*
modification history
--------------------
2021/8/5 3:40 下午, by lishanlei, create
*/

/*
DESCRIPTION
please input description
*/

package tree

import "strings"

type HuffmanNode struct {
	id int
	weight int
	parent int
	lChild int
	rChild int
}

type HuffmanTree struct {
	nodes []*HuffmanNode
	nodesNumber int
	weightLength int
	char2CodeMap map[byte]string
	code2CharMap map[string]byte
}

func NewHuffmanTree(weight []int) *HuffmanTree {
	// 节点数：2n-1
	max := 2 * len(weight) - 1
	nodes := make([]*HuffmanNode, max)
	for i := 0; i < len(weight); i++ {
		nodes[i] = &HuffmanNode{
			id:     i,
			weight: weight[i],
			lChild: -1,
			rChild: -1,
		}
	}
	return &HuffmanTree{
		nodes:        nodes,
		nodesNumber:  max,
		weightLength: len(weight),
	}
}

func (h *HuffmanTree) CreateTree() *HuffmanTree {
	for i := h.weightLength; i < h.nodesNumber; i++ {
		newNode := new(HuffmanNode)
		newNode.id = i
		s1, s2 := h.FindMinNode(i)
		newNode.weight = s1.weight + s2.weight
		newNode.lChild, newNode.rChild = s1.id, s2.id
		s1.parent, s2.parent = newNode.id, newNode.id
		h.nodes[i] = newNode
	}
	return h
}

func (h *HuffmanTree) CreateCode(chars []byte) {
	codes := make(map[byte]string)
	for i := 0; i < h.weightLength; i++ {
		parent := h.nodes[i].parent
		curId := i
		code := ""
		for parent != 0 {
			prevNode := h.nodes[parent]
			if prevNode.lChild == curId {
				code += "0"
			} else if prevNode.rChild == curId {
				code += "1"
			}
			parent = prevNode.parent
			curId = prevNode.id
		}
		codes[chars[i]] = h.StrTraverser(code) // 从叶子节点开始找的，正常编码是从根节点开始找，所以需要反转
	}
	h.char2CodeMap = codes
	h.code2CharMap = h.MapTraverse(codes)
}

func (h *HuffmanTree) Encode(str string) string {
	var b strings.Builder
	for i := 0; i < len(str); i++ {
		s := byte(str[i])
		if code, ok := h.char2CodeMap[s]; ok {
			b.WriteString(code)
		}
	}
	return b.String()
}

func (h *HuffmanTree) Decode(codes string) string {
	var b strings.Builder
	i, j := 0, 1
	for j <= len(codes) {
		c := codes[i:j]
		if char, ok := h.code2CharMap[c]; ok {
			b.WriteByte(char)
			i = j
		} else {
			j++
		}
	}
	return b.String()
}

func (h *HuffmanTree) StrTraverser(str string) string {
	l := len(str)
	b := []byte(str)
	for i := 0; i < l / 2; i++ {
		b[i], b[l-i-1] = b[l-i-1], b[i]
	}
	return string(b)
}

func (h *HuffmanTree) MapTraverse(m map[byte]string) map[string]byte {
	codeMap := make(map[string]byte)
	for char, code := range m {
		codeMap[code] = char
	}
	return codeMap
}

func (h *HuffmanTree) FindMinNode(index int) (s1 *HuffmanNode, s2 *HuffmanNode) {
	s1, s2 = new(HuffmanNode), new(HuffmanNode)
	s1.weight, s2.weight = 1<<32 - 1, 1<<32 - 1

	for i := 0; i < index; i++ {
		if h.nodes[i].parent > 0 {
			continue
		}
		if s1.weight > h.nodes[i].weight {
			s2 = s1
			s1 = h.nodes[i]
		} else if s2.weight > h.nodes[i].weight {
			s2 = h.nodes[i]
		}
	}
	return
}
