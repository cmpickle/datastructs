package tree

import "testing"

func TestTreeSum(t *testing.T) {
	root := NewNode(2)
	child1 := NewNode(2)
	child2 := NewNode(2)
	root.Left = child1
	root.Right = child2
	child3 := NewNode(2)
	child4 := NewNode(2)
	child1.Left = child3
	child1.Right = child4
	child5 := NewNode(2)
	child6 := NewNode(2)
	child3.Left = child5
	child3.Right = child6

	expected := 14
	result := root.Sum()

	if result != expected {
		t.Errorf("Tree sum failed, expected %d, but got %d", expected, result)
	} else {
		t.Logf("Tree sum success, got %d", result)
	}
}

func TestUnivalTree(t *testing.T) {
	root := NewNode(2)
	child1 := NewNode(2)
	child2 := NewNode(2)
	root.Left = child1
	root.Right = child2
	child3 := NewNode(2)
	child4 := NewNode(2)
	child1.Left = child3
	child1.Right = child4
	child5 := NewNode(2)
	child6 := NewNode(2)
	child3.Left = child5
	child3.Right = child6

	expected := 7
	result := root.UnivalCount()

	if result != expected {
		t.Errorf("Tree unival count failed, expected %d, but got %d", expected, result)
	} else {
		t.Logf("Tree unival count success, got %d", result)
	}
}

func TestAreSameBinaryTree(t *testing.T) {
	root := NewNode(1)
	child1 := NewNode(2)
	child2 := NewNode(3)
	root.Left = child1
	root.Right = child2
	child3 := NewNode(4)
	child4 := NewNode(5)
	child1.Left = child3
	child1.Right = child4
	child5 := NewNode(6)
	child6 := NewNode(7)
	child3.Left = child5
	child3.Right = child6

	broot := NewNode(1)
	bchild1 := NewNode(2)
	bchild2 := NewNode(3)
	broot.Left = bchild1
	broot.Right = bchild2
	bchild3 := NewNode(4)
	bchild4 := NewNode(5)
	bchild1.Left = bchild3
	bchild1.Right = bchild4
	bchild5 := NewNode(6)
	bchild6 := NewNode(7)
	bchild3.Left = bchild5
	bchild3.Right = bchild6

	expected := true
	result := AreSame(root, broot)

	if result != expected {
		t.Errorf("Tree unival count failed, expected %v, but got %v", expected, result)
	} else {
		t.Logf("Tree unival count success, got %v", result)
	}
}

func TestInvertBinaryTree(t *testing.T) {
	root := NewNode(1)
	child1 := NewNode(2)
	child2 := NewNode(3)
	root.Left = child1
	root.Right = child2
	child3 := NewNode(4)
	child4 := NewNode(5)
	child1.Left = child3
	child1.Right = child4
	child5 := NewNode(6)
	child6 := NewNode(7)
	child3.Left = child5
	child3.Right = child6

	broot := NewNode(1)
	bchild1 := NewNode(3)
	bchild2 := NewNode(2)
	broot.Left = bchild1
	broot.Right = bchild2
	bchild3 := NewNode(5)
	bchild4 := NewNode(4)
	bchild1.Left = bchild3
	bchild1.Right = bchild4
	bchild5 := NewNode(7)
	bchild6 := NewNode(6)
	bchild3.Left = bchild5
	bchild3.Right = bchild6

	expected := true
	result := AreSame(root, InvertBinaryTree(broot))

	if result != expected {
		t.Errorf("Tree unival count failed, expected %v, but got %v", expected, result)
	} else {
		t.Logf("Tree unival count success, got %v", result)
	}
}

//        5
//       /  \
//      3    8
//     / \   / \
//    2   4 6   10
//   /       \  /
//  1         7 9
func TestFlattenBinaryTree(t *testing.T) {
	root := NewNode(5)
	child1 := NewNode(3)
	child2 := NewNode(8)
	child3 := NewNode(2)
	child4 := NewNode(4)
	child5 := NewNode(6)
	child6 := NewNode(10)
	child7 := NewNode(1)
	child8 := NewNode(7)
	child9 := NewNode(9)
	root.Left = child1
	root.Right = child2
	child1.Left = child3
	child1.Right = child4
	child3.Left = child7
	child2.Left = child5
	child2.Right = child6
	child5.Right = child8
	child6.Left = child9

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := FlattenInPlace(root)

	if !arrSame(result, expected) {
		t.Errorf("Tree unival count failed, expected %v, but got %v", expected, result)
	} else {
		t.Logf("Tree unival count success, got %v", result)
	}
}

func arrSame(left []int, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
