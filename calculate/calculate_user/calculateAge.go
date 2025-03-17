// calculate/calculate_age.go

package calculateuser


import (
	"fmt"

)



func UserBorn() {
	/*
       Calculate user user born	
	*/
	showUserDate()
}


func getUserDateYears() int {
	var years int
	fmt.Print("Enter Your Date Years: ")
	_, err := fmt.Scan(&years)
	if err != nil {
		fmt.Println("Please only number")
		return getUserDateYears()
	} 
	return years
}


func calculateAge() int {
	var years = getUserDateYears()
	return 2025 - years
}


func showUserDate() {
	fmt.Println("Your Age is", calculateAge())
}

