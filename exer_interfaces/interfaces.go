package exer_interfaces

import (
	"WebstormProjects/UDEMY/GO/FirstSteps-GO/interfaces"
	"fmt"
)

func BreatheHumans(hu interfaces.Human) {
	hu.Breathe()
	fmt.Printf("I'm %s, and i`m breathing \n", hu.Sex())
}
