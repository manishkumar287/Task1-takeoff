
package utils

import (
	"log"
)

func NonAdmin_UpdateEmployee(employees *[] Employee,_id int) {
		index := -1
	var updatedEmp Employee
	for i, employee := range *employees {
		if employee.ID == _id && employee.Role != "Admin" {
			index = i
			updatedEmp = employee
			log.Println("So the person with id is below")
			log.Println(employee.FirstName)
			log.Println(employee.LastName)
			log.Println(employee.Email)
			log.Println(employee.Role)

			break
		}
	}

	// If the employee is not found, then return an error.
	if index == -1 {
		log.Println("employee not found with this id: ", _id)
		return
	}

	// Update the employee's information.
	log.Println("To update the employee please enter these details")
	log.Println("Get the employee's first name.")
	firstName := readLine()

	log.Println("Get the employee's last name.")
	lastName := readLine()

	log.Println("Get the employee's email.")
	email := readLine()

	log.Println("Get the employee's role.")
	role := readLine()

	updatedEmp.FirstName = firstName
	updatedEmp.LastName = lastName
	updatedEmp.Email = email
	updatedEmp.Role = role

	(*employees)[index] = updatedEmp

	return
	
}


