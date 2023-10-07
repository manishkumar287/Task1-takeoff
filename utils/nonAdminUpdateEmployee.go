
package utils

import (
	"fmt"
)

func NonAdmin_UpdateEmployee(employees *[] Employee,_id int) {
		index := -1
	var updatedEmp Employee
	for i, employee := range *employees {
		if employee.ID == _id && employee.Role != "Admin" {
			index = i
			updatedEmp = employee
			fmt.Println("So the person with id is below")
			fmt.Println(employee.FirstName)
			fmt.Println(employee.LastName)
			fmt.Println(employee.Email)
			fmt.Println(employee.Role)

			break
		}
	}

	// If the employee is not found, then return an error.
	if index == -1 {
		fmt.Println("employee not found with this id: ", _id)
		return
	}

	// Update the employee's information.
	fmt.Println("To update the employee please enter these details")
	fmt.Println("Get the employee's first name.")
	firstName := readLine()

	fmt.Println("Get the employee's last name.")
	lastName := readLine()

	fmt.Println("Get the employee's email.")
	email := readLine()

	fmt.Println("Get the employee's role.")
	role := readLine()

	updatedEmp.FirstName = firstName
	updatedEmp.LastName = lastName
	updatedEmp.Email = email
	updatedEmp.Role = role

	(*employees)[index] = updatedEmp

	return
	
}


