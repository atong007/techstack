package doubly

import (
	"bytes"
	"fmt"
	"strings"
)

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func (n *Node) String() string {
	head := n
	var buf bytes.Buffer
	for head != nil {
		buf.WriteString(fmt.Sprintf("%v<->", head.Val))
		head = head.Next
	}
	return strings.TrimRight(buf.String(), "<->")
}

func randomDoublyNode(count int) *Node {
	root := &Node{Val: -1}
	node := root
	var prev *Node
	for i := 1; i <= count; i++ {
		prev = node
		node.Next = &Node{
			Val:  i,
			Prev: prev,
		}
		node = node.Next
	}
	root.Next.Prev = nil
	return root.Next
}
