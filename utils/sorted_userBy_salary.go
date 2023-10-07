package utils

import (
	"log"
	"sort"
)

func Sorted_userBy_salary(employees []Employee) {
	sort.Slice(employees, func(i, j int) bool {
		return employees[i].Salary < employees[j].Salary
	})

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
