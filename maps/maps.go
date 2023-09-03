package maps

import "fmt"

func ViewMaps() {
	countryCapital := make(map[string]string)
	fmt.Println(countryCapital)

	countryCapital["Mexico"] = "CDMX"
	countryCapital["Colombia"] = "Bogota"

	fmt.Println(countryCapital)
	fmt.Println(countryCapital["Mexico"])

	championship := map[string]int{
		"Barcelona":    25,
		"Real Madrid":  28,
		"Chivas":       12,
		"Boca Juniors": 18,
	}

	fmt.Println(championship)

	for team, score := range championship {
		fmt.Printf("The team %s has a score of %d \n", team, score)
	}

	delete(championship, "Chivas")

	fmt.Println(championship)

	score, exist := championship["Chivas"]

	fmt.Printf("The score is %d and the team exist %t \n", score, exist)
}
