package main

import (
	"WebstormProjects/UDEMY/GO/FirstSteps-GO/middleware"
)

// "WebstormProjects/UDEMY/GO/FirstSteps-GO/exercices"
// "WebstormProjects/UDEMY/GO/FirstSteps-GO/files"
// "WebstormProjects/UDEMY/GO/FirstSteps-GO/iterations"
// "WebstormProjects/UDEMY/GO/FirstSteps-GO/keyboard"
// "WebstormProjects/UDEMY/GO/FirstSteps-GO/variables"
// "fmt"
// "runtime"
// "WebstormProjects/UDEMY/GO/FirstSteps-GO/functions"
// "WebstormProjects/UDEMY/GO/FirstSteps-GO/arrays_slices"
// "WebstormProjects/UDEMY/GO/FirstSteps-GO/maps"
// "WebstormProjects/UDEMY/GO/FirstSteps-GO/users"
// e "WebstormProjects/UDEMY/GO/FirstSteps-GO/exer_interfaces"
// "WebstormProjects/UDEMY/GO/FirstSteps-GO/models"
// "WebstormProjects/UDEMY/GO/FirstSteps-GO/defer_panic"

func main() {
	// variables.ShowInt()
	// variables.VariablesRest()
	// state, text := variables.TextToConverter(1880)
	// fmt.Println(state, text)

	// if os := runtime.GOOS; os == "windows." || os == "OS X." {
	// 	fmt.Println("Esto no es Windows, es", os)

	// } else {
	// 	fmt.Println("Esto es linux")
	// }

	// switch os := runtime.GOOS; os {
	// case "windows":
	// 	fmt.Println("Esto es Windows")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os)
	// }

	// fmt.Println(exercices.TextToConverterAndValidateAHundred("100"))

	// keyboard.InsertNumbers()

	// iterations.Iteration()

	// exercices.MultiplicationOfTenTable()

	// files.SaveTable()
	// files.AddTable()
	// files.ReadFile()

	// functions.Calculates()
	// functions.CallClosure()
	// functions.Exponential(2)

	// arrays_slices.ViewArrays()
	// arrays_slices.ViewMatrix()
	// arrays_slices.ViewSlice()

	// arrays_slices.Capacity()

	// maps.ViewMaps()

	// users.CreateUser()

	// Pepe := new(models.Man)
	// e.BreatheHumans(Pepe)

	// Maria := new(models.Woman)
	// e.BreatheHumans(Maria)

	// defer_panic.ViewDefer()
	// defer_panic.ExamplePanic()

	// channel1 := make(chan bool)
	// go goroutines.MySlowName("Jandrohen", channel1)
	// defer func() {
	// 	<-channel1
	// }()
	// fmt.Println("I am here")

	// webserver.MiWebServer()

	middleware.MyMiddleware()

}
