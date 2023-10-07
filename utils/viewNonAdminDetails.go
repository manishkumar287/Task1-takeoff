
package utils

import (
	"fmt"
)

func View_NonAdminDetails(employees [] Employee) {
	for _, employee := range employees {
		if employee.Role != "Admin" {
			fmt.Println("So the Employee id is", employee.ID)
			fmt.Println(employee.FirstName)
			fmt.Println(employee.LastName)
			fmt.Println(employee.Email)
			fmt.Println(employee.Role)
		}
	}
	return
}


