package main

import (
	"WebstormProjects/UDEMY/GO/FirstSteps-GO/variables"
	"fmt"
)

func main() {
	variables.ShowInt()
	variables.VariablesRest()
	state, text := variables.TextToConverter(1880)
	fmt.Println(state, text)
}
