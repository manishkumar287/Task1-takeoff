package utils

import (
	"fmt"
	"log"
)

func readLine() string {
	var input string
	fmt.Scanln(&input)

	return input
}

func Add_user(employees *[]Employee) {

	log.Println("Get the employee's first name.")
	firstName := readLine()

	log.Println("Get the employee's last name.")
	lastName := readLine()

	log.Println("Get the employee's email.")
	email := readLine()

	log.Println("Get the employee's password.")
	password := readLine()

	log.Println("Get the employee's Phone.")
	phoneNo := readLine()

	log.Println("Get the employee's role.")
	role := readLine()

	log.Println("Get the employee's salary.")
	var salary float64
	fmt.Scanln(&salary)

	// Create a new Employee struct.

	employee := Employee{
		ID:        len(*employees) + 1,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		PhoneNo:   phoneNo,
		Role:      role,
		Salary:    salary,
	}

	// Add the employee to the employee splice.
	*employees = append(*employees, employee)

	// Display a message to the user.
	log.Println("Employee added successfully!")

	log.Println("List of users currently inside the organisation")
	for _, employee := range *employees {
		log.Println(employee.FirstName)

	}

}
