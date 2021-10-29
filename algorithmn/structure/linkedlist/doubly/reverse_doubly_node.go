package doubly

// Reverse 反转双向链表
func Reverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	node := head
	for {
		next := node.Next
		// 指针对调
		node.Prev, node.Next = node.Next, node.Prev
		if next == nil {
			return node
		}
		node = next
	}
}
