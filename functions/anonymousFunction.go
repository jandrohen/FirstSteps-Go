package functions

import "fmt"

func Calculates() {

	var num3 int = 5
	var num4 int = 5

	calculate := func(num1 int, num2 int) int {
		return num1 + num2 + num3 + num4
	}

	fmt.Println(calculate(5, 5))

	calculate = func(num1 int, num2 int) int {
		return num1 * num2 * num3
	}

	fmt.Println(calculate(5, 5))
}
