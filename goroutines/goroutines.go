package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MySlowName(name string) {
	letras := strings.Split(name, "")

	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
