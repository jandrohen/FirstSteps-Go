package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var leyend string
var err error

func InsertNumbers() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Insert number 1: ")
	if scanner.Scan() {
		number1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("The value is not a number " + err.Error())
		}
	}

	fmt.Println("Insert number 2: ")
	if scanner.Scan() {
		number2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("The value is not a number " + err.Error())
		}
	}

	fmt.Println("Insert leyend : ")
	if scanner.Scan() {
		leyend = scanner.Text()
	}

	fmt.Println(leyend, number1*number2)
}
