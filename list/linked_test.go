package list

import "testing"

func TestToStringLinkedList(t *testing.T) {
	expected := "1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> "
	var head *Node
	curr := head
	prev := head
	for i := 0; i < 10; i++ {
		prev = curr
		curr = &Node{Value: i + 1}
		if i == 0 {
			head = curr
		}
		if prev != nil {
			prev.next = curr
		}
	}
	result := ToString(head)

	if result != expected {
		t.Errorf("To string failed, expected %v, but got %v", expected, result)
	} else {
		t.Logf("To string success, got %v", result)
	}
}

func TestLinkedListReversal(t *testing.T) {
	expected := "10 -> 9 -> 8 -> 7 -> 6 -> 5 -> 4 -> 3 -> 2 -> 1 -> "
	var head *Node
	curr := head
	prev := head
	for i := 0; i < 10; i++ {
		prev = curr
		curr = &Node{Value: i + 1}
		if i == 0 {
			head = curr
		}
		if prev != nil {
			prev.next = curr
		}
	}
	result := ToString(Reverse(head))

	if result != expected {
		t.Errorf("To string failed, expected %v, but got %v", expected, result)
	} else {
		t.Logf("To string success, got %v", result)
	}
}

func TestLinkedListReversal2(t *testing.T) {
	expected := "10 -> 9 -> 8 -> 7 -> 6 -> 5 -> 4 -> 3 -> 2 -> 1 -> "
	var head *Node
	curr := head
	prev := head
	for i := 0; i < 10; i++ {
		prev = curr
		curr = &Node{Value: i + 1}
		if i == 0 {
			head = curr
		}
		if prev != nil {
			prev.next = curr
		}
	}
	result := ToString(ReverseLinkedList(head))

	if result != expected {
		t.Errorf("To string failed, expected %v, but got %v", expected, result)
	} else {
		t.Logf("To string success, got %v", result)
	}
}
