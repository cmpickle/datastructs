package main

import (
	"fmt"
)

func main() {
	// root := tree.NewNode(2)
	// child1 := tree.NewNode(2)
	// child2 := tree.NewNode(2)
	// root.Left = child1
	// root.Right = child2
	// child3 := tree.NewNode(2)
	// child4 := tree.NewNode(2)
	// child1.Left = child3
	// child1.Right = child4
	// child5 := tree.NewNode(2)
	// child6 := tree.NewNode(2)
	// child3.Left = child5
	// child3.Right = child6

	// fmt.Println("The sum of the tree is:")
	// fmt.Println(root.Sum())
	// fmt.Println("The count of unival trees is:")
	// fmt.Println(root.UnivalCount())
	// fmt.Println(tree.CountUnivals(*root))

	// arr1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr2 := []int{-30, -15, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 42, 100}
	// fmt.Println("Index of 0:")
	// fmt.Println(search.Binary(arr1, 0))
	// fmt.Println("Index of 3:")
	// fmt.Println(search.Binary(arr1, 3))
	// fmt.Println("Index of 0:")
	// fmt.Println(search.Binary(arr1, 0))
	// fmt.Println("Index of -1:")
	// fmt.Println(search.Binary(arr2, -1))

	//        1
	//    /       \
	//   2         3
	//  / \        /\
	// 4   5      6  7
	//    / \     \
	//   8   9     10

	// n1 := tree.NewNode(1)
	// n2 := tree.NewNode(2)
	// n3 := tree.NewNode(3)
	// n4 := tree.NewNode(4)
	// n5 := tree.NewNode(5)
	// n6 := tree.NewNode(6)
	// n7 := tree.NewNode(7)
	// n8 := tree.NewNode(8)
	// n9 := tree.NewNode(9)
	// n10 := tree.NewNode(10)
	// n1.Left = n2
	// n1.Right = n3
	// n2.Left = n4
	// n2.Right = n5
	// n3.Left = n6
	// n3.Right = n7
	// n5.Left = n8
	// n5.Right = n9
	// n6.Right = n10
	// // fmt.Println(n1)
	// tree.Bfs(n1)
	// fmt.Println()
	// tree.Dfs(n1)

	// fmt.Println()
	// matrix := [][]int{[]int{1, 2, 3, 4}, []int{12, 13, 14, 5}, []int{11, 16, 15, 6}, []int{10, 9, 8, 7}}
	// matrix2 := [][]int{[]int{1, 2, 3, 4}, []int{10, 11, 12, 5}, []int{9, 8, 7, 6}}
	// matrix3 := [][]int{[]int{1, 2, 3}, []int{12, 13, 4}, []int{11, 14, 5}, []int{10, 15, 6}, []int{9, 8, 7}}
	// result := SpiralTraversal(matrix)
	// result2 := SpiralTraversal(matrix2)
	// result3 := SpiralTraversal(matrix3)
	// for i := 0; i < len(result); i++ {
	// 	fmt.Printf("%d ", result[i])
	// }
	// fmt.Println()
	// for i := 0; i < len(result2); i++ {
	// 	fmt.Printf("%d ", result2[i])
	// }
	// fmt.Println()
	// for i := 0; i < len(result3); i++ {
	// 	fmt.Printf("%d ", result3[i])
	// }
	// fmt.Println()

	array := [][]int{
		{1, 2, 3, 4},
		{5, 6, 8, 7},
		{9, 10, 11, 12}}
	// fmt.Println(array[2][1])
	myRow := array[1]
	// for _, row := range array {
	// 	fmt.Println(row)
	// }
	// fmt.Println(myRow)
	// fmt.Println("Pop")
	// fmt.Println(myRow[:len(myRow)-1])
	// fmt.Println("Dequeue")
	// fmt.Println(myRow[1:])
	i := 0
	j := len(myRow) - 1
	for i < j {
		myRow[i], myRow[j] = myRow[j], myRow[i]
		i++
		j--
	}
	fmt.Println(myRow)

	// tallyMap := make(map[string]int)
	// tallyMap["go"] = 40
	// tallyMap["html"] = 0
	// tallyMap["c#"] = 10

	// for i, value := range tallyMap {
	// 	fmt.Println(i)
	// 	fmt.Println(value)
	// 	fmt.Println(tallyMap[i])
	// }
}

func SpiralTraversal(input [][]int) []int {
	var left, right, top, bottom int
	right = len(input[0]) - 1
	bottom = len(input) - 1
	result := make([]int, 0)

	for left <= right && bottom >= top {
		for col := left; col < right+1; col++ {
			result = append(result, input[top][col])
		}
		for row := top + 1; row < bottom+1; row++ {
			result = append(result, input[row][right])
		}
		for col := right; col > left; col-- {
			if top == bottom {
				break
			}
			result = append(result, input[bottom][col])
		}
		for row := bottom; row >= top+1; row-- {
			if left == right {
				break
			}
			result = append(result, input[row][left])
		}
		top++
		right--
		bottom--
		left++
	}
	return result
}
