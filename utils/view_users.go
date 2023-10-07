
package utils

import (
	"log"
)

func View_allUsers(employees [] Employee) {
	for _, employee := range employees {
		log.Println("")
		log.Println("So the person with id is", employee.ID)
		log.Println(employee.FirstName)
		log.Println(employee.LastName)
		log.Println(employee.Email)
		log.Println(employee.PhoneNo)
		log.Println(employee.Role)

	}
	return
}


