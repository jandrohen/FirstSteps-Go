package middleware

import "fmt"

func sum(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func MyMiddleware() {
	fmt.Println("Middleware")

	result := operationsMidd(sum)(1, 2)
	fmt.Println(result)
	result = operationsMidd(subtract)(10, 5)
	fmt.Println(result)
	result = operationsMidd(multiply)(2, 3)
	fmt.Println(result)
}

func operationsMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Start")
		result := f(a, b)
		fmt.Println("End")
		return result
	}
}