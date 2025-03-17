// calculate/entire.go

package calculate

import (
	"fmt"
	"myproject/calculate/calculate_user" 
)

func Show() {
	var menu int
	fmt.Println(`
1. Get Your Years Born
2. Get Your Age now
Please choose the menu:  
	`)
	_, err := fmt.Scan(&menu)
	if err != nil {
		fmt.Println("Only Number")
		Show()
		return
	}

    if (menu == 1) {
           calculateuser.UserBorn()
	} else {
		calculateuser.UserAge()
	}
}