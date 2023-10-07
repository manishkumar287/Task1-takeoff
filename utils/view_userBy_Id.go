package utils

import (
	"fmt"
	"log"
)

func View_userBy_id(employees []Employee) {
	log.Println("Please enter the id of user to see its details")
	var _id int
	fmt.Scanln(&_id)

	index := -1

	for i, employee := range employees {
		if employee.ID == _id {
			index = i
			log.Println("So the person with id is below")
			log.Println(employee.FirstName)
			log.Println(employee.LastName)
			log.Println(employee.Email)
			log.Println(employee.PhoneNo)
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
