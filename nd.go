// To execute Go code, please declare a func main() in a package "main"

package main

type Ndarray struct {
	cols int
	rows int
	arr  []int
}

func New(row int, col int) *Ndarray {
	rowArr := make([]int, row*col)
	return &Ndarray{arr: rowArr, cols: col, rows: row}
}

func (ndarr *Ndarray) GetElement(row int, col int) int {
	return ndarr.arr[row*ndarr.cols+col]
}

func (ndarr *Ndarray) SetElement(row int, col int, val int) {
	ndarr.arr[row*ndarr.cols+col] = val
}

// func main() {
// 	col := 5
// 	row := 5
// 	ndarr := New(row, col)
// 	for i := 0; i < row; i++ {
// 		for j := 0; j < col; j++ {
// 			num, _ := strconv.Atoi(strconv.Itoa(i) + strconv.Itoa(j))
// 			ndarr.SetElement(i, j, num)
// 		}
// 	}
// 	for i := 0; i < row; i++ {
// 		for j := 0; j < col; j++ {
// 			fmt.Printf("%02d ", ndarr.GetElement(i, j))
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Println(ndarr.arr)
// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Printf("%02d\n", ndarr.GetElement(2, 3))
// }
