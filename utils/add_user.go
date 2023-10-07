package utils

import (
	"fmt"
)

func readLine() string {
	var input string
	fmt.Scanln(&input)

	return input
}

func Add_user(employees *[]Employee) {

	fmt.Println("Get the employee's first name.")
	firstName := readLine()

	fmt.Println("Get the employee's last name.")
	lastName := readLine()

	fmt.Println("Get the employee's email.")
	email := readLine()

	fmt.Println("Get the employee's password.")
	password := readLine()

	fmt.Println("Get the employee's Phone.")
	phoneNo := readLine()

	fmt.Println("Get the employee's role.")
	role := readLine()

	fmt.Println("Get the employee's salary.")
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
	fmt.Println("Employee added successfully!")

	fmt.Println("List of users currently inside the organisation")
	for _, employee := range *employees {
		fmt.Printf("Name is %s, and id is %d", employee.FirstName, employee.ID)

	}

}
