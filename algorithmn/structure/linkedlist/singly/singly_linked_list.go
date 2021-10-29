package singly

import (
	"bytes"
	"fmt"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

// String implement stringer interface for format printing
func (n *Node) String() string {
	node := n
	var buf bytes.Buffer
	for node != nil {
		buf.WriteString(fmt.Sprintf("%v->", node.Val))
		node = node.Next
	}
	return strings.TrimRight(buf.String(), "->")
}

func randomSinglyNode(count int) *Node {
	node := &Node{Val: -1}
	root := node
	for i := 0; i < count; i++ {
		node.Next = &Node{
			Val: i + 1,
		}
		node = node.Next
	}
	return root.Next
}

func (n *Node) Insert(node *Node, index int) {
	head := n
	for i := 1; i <= index && head != nil; i++ {
		head = head.Next
	}
	node.Next = head.Next.Next
	head.Next = node
}

func (n *Node) Delete(node *Node) {
	if node == nil {
		return
	}
	root := &Node{Val: -1}
	root.Next = n
	for root != nil {
		if root.Next == node {
			root.Next = root.Next.Next
			node.Next = nil
			return
		}
	}
}
