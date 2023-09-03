package arrays_slices

import "fmt"

var table [10]int = [10]int{10, 0, 59}
var matrix [5][7]int

func ViewArrays() {
	table[7] = 33
	table[5] = 15

	table2 := [10]string{"España", "Francia", "Portugal"}
	fmt.Println(table)
	fmt.Println(table2)

	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}
}

func ViewMatrix() {
	matrix[3][5] = 15
	fmt.Println(matrix)
}
