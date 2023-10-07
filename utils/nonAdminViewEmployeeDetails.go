package utils

import (
	"fmt"
)

func View_NonAdminEmployeesDetails(employees []Employee, _id int) {
	index := -1

	for i, employee := range employees {
		if employee.ID == _id && employee.Role != "Admin" {
			index = i
			fmt.Println("So the User whose id id ", _id)
			fmt.Println(employee.FirstName)
			fmt.Println(employee.LastName)
			fmt.Println(employee.Email)
			fmt.Println(employee.Role)
			break
		}
	}
	if index == -1 {
		fmt.Println("employee not found with this id: ", _id)
		return
	}
	return
}
