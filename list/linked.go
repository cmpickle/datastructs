package list

import (
	"fmt"
	"strings"
)

type Node struct {
	Value int
	prev  *Node
	next  *Node
}

func ToString(n *Node) string {
	var output strings.Builder
	for i := 0; n != nil; i++ {
		fmt.Fprintf(&output, "%d -> ", n.Value)
		n = n.next
	}
	return output.String()
}

func Reverse(n *Node) *Node {
	var prev *Node
	head := n
	curr := n
	next := n.next

	for curr != nil {
		if prev == nil {
			curr.next = nil
		} else {
			curr.next = prev
		}
		prev = curr
		curr = next
		if curr != nil {
			next = next.next
		}
		if curr == nil {
			head = prev
		}
	}
	return head
}
