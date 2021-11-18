package tree

import "fmt"

//        a
//    /       \
//   b         c
//  / \        /\
// d   e      f  g
//    / \     \
//   h   i     j

// a

func Bfs(curr *Node) {
	queue := []*Node{curr}
	for len(queue) > 0 {
		curr = queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", curr.Value)
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
}

func Dfs(node *Node) {
	// fmt.Printf("%d ", curr.Value)
	// if curr.Left != nil {
	// 	Dfs(curr.Left)
	// }
	// if curr.Right != nil {
	// 	Dfs(curr.Right)
	// }
	stack := []*Node{node}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", curr.Value)
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
}
