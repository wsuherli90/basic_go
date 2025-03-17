// calculate/calculate.go
package calculate

import (
	"fmt"
	"strconv"
)

func Result() { // Ubah result menjadi Result (huruf besar)
	age := askUser()
	yearsBorn := strconv.Itoa(CalculateFutureValue(age)) // Ubah calculateFutureValue
	outputText(yearsBorn)
}

func askUser() int {
	var age int
	fmt.Print("Enter your age: ")
	_, err := fmt.Scan(&age)
	if err != nil {
		fmt.Println("Only Number ")
		return askUser()
	}

	return age
}

func CalculateFutureValue(age int) int { // Ubah calculateFutureValue
	return 2025 - age
}

func outputText(text string) {
	fmt.Println("Your born  ", text)
}