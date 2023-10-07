
package utils

import (
	"fmt"
)

func View_userBy_id(employees [] Employee) {
	fmt.Println("Please enter the id of user to see its details")
	var _id int
	fmt.Scanln(&_id)

	index := -1

	for i, employee := range employees {
		if employee.ID == _id {
			index = i
			fmt.Println("So the person with id is below")
			fmt.Println(employee.FirstName)
			fmt.Println(employee.LastName)
			fmt.Println(employee.Email)
			fmt.Println(employee.PhoneNo)
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


