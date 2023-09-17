package defer_panic

import (
	"fmt"
	"log"
)

func ViewDefer() {
	fmt.Println("this is the first message")
	defer fmt.Println("this is the final message")

	fmt.Println("this is the second message")
}

func ExamplePanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("this is the error: %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("this is a panic")
	}
}
