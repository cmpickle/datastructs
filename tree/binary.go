package tree

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(v int) *Node {
	return &Node{Value: v}
}

func AreSame(left *Node, right *Node) bool {
	for {
		if left.Value != right.Value {
			return false
		} else if left.Left != nil && right.Left != nil && !AreSame(left.Left, right.Left) {
			return false
		} else if left.Right != nil && right.Right != nil && !AreSame(left.Right, right.Right) {
			return false
		} else {
			return true
		}
	}
}

func (n Node) Sum() int {
	sum := 0
	if n.Left != nil {
		sum += n.Left.Sum()
	}
	if n.Right != nil {
		sum += n.Right.Sum()
	}
	sum += n.Value
	return sum
}

func (n Node) UnivalCount() int {
	sum := 0
	if n.isUnival() {
		sum += 1
	}
	if n.Left != nil {
		sum += n.Left.UnivalCount()
	}
	if n.Right != nil {
		sum += n.Right.UnivalCount()
	}
	return sum
}

func (n Node) isUnival() bool {
	if n.Left == nil && n.Right == nil {
		return true
	}
	if n.Left != nil && n.Right != nil && n.Left.Value == n.Right.Value && n.Left.Value == n.Value {
		return n.Left.isUnival() && n.Right.isUnival()
	}
	return false
}

func CountUnivals(n Node) int {
	sum, _ := univalCounter(n)
	return sum
}

func univalCounter(n Node) (int, bool) {
	sum := 0
	unival := true
	if n.Left != nil {
		s, u := univalCounter(*n.Left)
		sum += s
		unival = unival && u
	}
	if n.Right != nil {
		s, u := univalCounter(*n.Right)
		sum += s
		unival = unival && u
	}
	if n.Left == nil && n.Right == nil {
		return 1, true
	}
	if unival {
		sum += 1
	}
	return sum, unival
}

func InvertBinaryTree(n *Node) *Node {
	leftTmp := n.Left
	n.Left = n.Right
	n.Right = leftTmp
	if n.Left != nil {
		InvertBinaryTree(n.Left)
	}
	if n.Right != nil {
		InvertBinaryTree(n.Right)
	}
	return n
}

func FlattenInPlace(n *Node) []int {
	list := []int{}
	list = flattenHelper(n, list)
	return list
}

func flattenHelper(tree *Node, list []int) []int {
	if tree.Left != nil {
		list = flattenHelper(tree.Left, list)
	}
	list = append(list, tree.Value)
	if tree.Right != nil {
		list = flattenHelper(tree.Right, list)
	}
	return list
}
