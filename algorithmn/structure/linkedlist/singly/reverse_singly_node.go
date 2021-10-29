package singly

// ReverseWithTraverse reverse the linked list with traversing
// 遍历实现单链表反转
func ReverseWithTraverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	node := head
	var prev *Node
	for {
		// backup the next
		next := node.Next
		// change next pointer to prev one
		node.Next = prev
		// current node become prev node of the next
		prev = node
		// go to the next one
		if next == nil {
			return node
		}
		node = next
	}
}

// ReverseWithRecursive reverse the linked list with recursive
// 递归实现单链表反转
func ReverseWithRecursive(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	// reverse the node from next one
	node := ReverseWithTraverse(head.Next)
	// the next node of head now point to the last
	// check the image if you are confused about this
	head.Next.Next = head
	head.Next = nil
	return node
}
