package iterations

import (
	"fmt"
)

func Iteration() {
	for i := 10; i > 1; i-- {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
