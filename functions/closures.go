package functions

import "fmt"

func Table(value int) func() int {
	number := value
	sequence := 0
	return func() int {
		sequence++
		return number * sequence
	}
}

func CallClosure() {
	tableOfTwo := 2
	MyTable := Table(tableOfTwo)
	for i := 0; i < 10; i++ {
		fmt.Println(MyTable())
	}
}
