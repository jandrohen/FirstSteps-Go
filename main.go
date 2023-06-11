package main

import (
	"WebstormProjects/UDEMY/GO/FirstSteps-GO/exercices"
	"WebstormProjects/UDEMY/GO/FirstSteps-GO/variables"
	"fmt"
	"runtime"
)

func main() {
	variables.ShowInt()
	variables.VariablesRest()
	state, text := variables.TextToConverter(1880)
	fmt.Println(state, text)

	if os := runtime.GOOS; os == "windows." || os == "OS X." {
		fmt.Println("Esto no es Windows, es", os)

	} else {
		fmt.Println("Esto es linux")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Esto es Windows")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}

	fmt.Println(exercices.TextToConverterAndValidateAHundred("100"))

}
