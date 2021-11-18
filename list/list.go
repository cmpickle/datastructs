package list

func ReverseLinkedList(list *Node) *Node {
	var prev *Node
	curr := list
	var next *Node

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}
