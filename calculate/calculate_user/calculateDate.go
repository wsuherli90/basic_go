// calculate/calculate_date.go
package calculateuser

import (
	"fmt"
)

func UserAge() {
	/*
	
	Calculate user Years Age
	
	*/

	showUserAge()
    
}

func getUserAge() int {
	var age int
	fmt.Print("Enter your age: ")
	_, err := fmt.Scan(&age)
	if err != nil {
		fmt.Println("Only Number ")
		return getUserAge()
	}

	return age
}

func calculateUserAge() int {
	var age = getUserAge()
	return 2025 - age
}

func showUserAge() {
	fmt.Println("Your born  ", calculateUserAge())
}