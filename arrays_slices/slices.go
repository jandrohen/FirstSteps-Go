package arrays_slices

import "fmt"

var tableS []int = []int{10, 0, 59}
var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func ViewSlice() {
	fmt.Println(tableS)

	slice := array[3:]   // create a slice from array, from position 3 to end
	slice2 := array[:6]  // create a slice from array, from position 0 to 6
	slice3 := array[3:6] // create a slice from array, from position 3 to 6

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

func Capacity() {
	elements := make([]int, 5, 20) // create a slice with 5 elements and capacity for 20
	fmt.Printf("Length: %d\nCapacity: %d\n", len(elements), cap(elements))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("\nLength: %d\nCapacity: %d\n", len(nums), cap(nums))
}
