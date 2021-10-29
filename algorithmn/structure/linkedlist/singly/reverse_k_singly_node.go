package singly

// ReverseKNodes reverse part of the linked list
// 单链表部分反转
// 1->2->3->4 from 1 to 3
// result: 3->2->1->4
func ReverseKNodes(head *Node, from, to int) *Node {
	if from < 1 || to < 1 || from >= to {
		return head
	}
	root := &Node{Val: -1}
	root.Next = head
	// 找到from和to位置的节点
	node := root
	var beginNode, endNode *Node
	for i := 1; i < from && node != nil; i++ {
		node = node.Next
	}
	beforeNode := node
	node = node.Next
	beginNode = node
	for i := 1; i < to-from+1 && node != nil; i++ {
		node = node.Next
	}
	endNode = node
	// 先把to位置之后节点保存起来
	last := endNode.Next
	// 把endNode指向nil
	endNode.Next = nil
	// 反转from到to的节点
	rNode := ReverseWithRecursive(beginNode)
	// 前后拼接起来
	// 由于beforeNode.Next指向了被反转后的rNode的尾部
	// 所以需要先将rNode的尾部指向剩余的尾部节点
	beforeNode.Next.Next = last
	beforeNode.Next = rNode
	return root.Next
}
