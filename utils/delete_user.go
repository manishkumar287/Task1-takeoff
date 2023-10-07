
package utils

import (
	"fmt"
)

func Delete_user(employees *[] Employee) {
	fmt.Println("Please provide the id of the Person you want delete")
	var _id int
	fmt.Scanln(&_id)
	// Find the index of the employee with the specified ID.
	index := -1
	for i, employee := range *employees {
		if employee.ID == _id {
			index = i
			break
		}
	}

	// If the employee is not found, then return.
	if index == -1 {
		fmt.Println("no id matched")
		return
	}

	// Remove the employee from the slice.
	*employees = append((*employees)[:index], (*employees)[index+1:]...)
	fmt.Println("List of left out employees")
	for _, employee := range (*employees) {
		fmt.Println(employee.FirstName)
	}

}


