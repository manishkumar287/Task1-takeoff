package utils

import (
	"log"
)

func View_NonAdminEmployeesDetails(employees []Employee, _id int) {
	index := -1

	for i, employee := range employees {
		if employee.ID == _id && employee.Role != "Admin" {
			index = i
			log.Println("So the User whose id id ", _id)
			log.Println(employee.FirstName)
			log.Println(employee.LastName)
			log.Println(employee.Email)
			log.Println(employee.Role)
			break
		}
	}
	if index == -1 {
		log.Println("employee not found with this id: ", _id)
		return
	}
	return
}
