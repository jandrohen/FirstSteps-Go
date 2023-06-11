package exercices

import (
	"strconv"
)

func TextToConverterAndValidateAHundred(numberToValidate string) (int, string) {
	numberConverted, err := strconv.Atoi(numberToValidate)
	if err != nil {
		return 0, "The value is not a number"
	}
	if numberConverted >= 100 {
		return numberConverted, "The number is greater than 100 or equal"
	} else {
		return numberConverted, "The number is less than 100"
	}

}
