package ch02

// KthToLast returns kth to the last element
func KthToLast(node *Node, k int) *Node {
	n := node
	l := 0
	for ; n != nil; l++ {
		n = n.next
	}

	n = node
	for i := 1; i <= l-k-1; i++ {
		n = n.next
	}

	return n
}

// KthToLastRecursive returns kth to the last element
func KthToLastRecursive(node *Node, k int) (*Node, int) {
	if node.next == nil {
		return node, 0
	}

	n, i := KthToLastRecursive(node.next, k)
	i++

	if i == k {
		return node, i
	}

	return n, i
}

// KthToLastIterative returns kth to the last element
func KthToLastIterative(node *Node, k int) *Node {
	iterator := node
	runner := node

	for i := 0; i < k; i++ {
		iterator = iterator.next
	}

	for iterator.next != nil {
		iterator = iterator.next
		runner = runner.next
	}

	return runner
}
