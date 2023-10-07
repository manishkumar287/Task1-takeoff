
package utils

import (
	"fmt"
)

func View_allUsers(employees [] Employee) {
	for _, employee := range employees {
		fmt.Println("")
		fmt.Println("So the person with id is", employee.ID)
		fmt.Println(employee.FirstName)
		fmt.Println(employee.LastName)
		fmt.Println(employee.Email)
		fmt.Println(employee.PhoneNo)
		fmt.Println(employee.Role)

	}
	return
}


