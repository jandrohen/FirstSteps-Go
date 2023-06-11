package files

import (
	"WebstormProjects/UDEMY/GO/FirstSteps-GO/exercices"
	"bufio"
	"fmt"
	"os"
)

var fileName string = "./files/txt/table.txt"

func SaveTable() {
	var text string = exercices.MultiplicationOfTenTable()
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file" + err.Error())
		return
	}
	fmt.Fprintln(file, text)
	file.Close()
}

func AddTable() {
	var text string = exercices.MultiplicationOfTenTable()
	if !Append(fileName, text) {
		fmt.Println("Error appending file")
	}
}

func Append(filen string, text string) bool {
	file, err := os.OpenFile(filen, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file" + err.Error())
		return false
	}

	_, err = file.WriteString(text)

	if err != nil {
		fmt.Println("Error writing file" + err.Error())
		return false
	}

	file.Close()
	return true
}

func ReadFile() {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file" + err.Error())
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		register := scanner.Text()
		fmt.Println("> " + register)
	}

	file.Close()
}
