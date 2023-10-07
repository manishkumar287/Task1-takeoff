package utils

import (
	"fmt"
	"sort"
)

func Sorted_userBy_salary(employees []Employee) {
	sort.Slice(employees, func(i, j int) bool {
		return employees[i].Salary < employees[j].Salary
	})

	for _, employee := range employees {
		fmt.Println("")
		fmt.Println("So the person with id is", employee.ID)
		fmt.Println(employee.FirstName)
		fmt.Println(employee.LastName)
		fmt.Println(employee.Email)
		fmt.Println(employee.PhoneNo)
		fmt.Println(employee.Role)

	}
	return
}
