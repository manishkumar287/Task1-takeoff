
package utils

import (
	"log"
)

func View_NonAdminDetails(employees [] Employee) {
	for _, employee := range employees {
		if employee.Role != "Admin" {
			log.Println("So the Employee id is", employee.ID)
			log.Println(employee.FirstName)
			log.Println(employee.LastName)
			log.Println(employee.Email)
			log.Println(employee.Role)
		}
	}
	return
}


