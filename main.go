package main

import (
	"fmt"
	"os"
	"sort"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
	PhoneNo   string
	Role      string
	Salary    float64
}

type Employees []Employee

var employees Employees = []Employee{{
	ID:        1,
	FirstName: "Manish",
	LastName:  "Kumar",
	Email:     "a@gmail.com",
	Password:  "abc",
	PhoneNo:   "64646",
	Role:      "Admin",
	Salary:    422343,
}, {
	ID:        2,
	FirstName: "Shashi",
	LastName:  "Kant",
	Email:     "b@gmail.com",
	Password:  "abc",
	PhoneNo:   "64646",
	Role:      "Employee",
	Salary:    41243,
}, {
	ID:        3,
	FirstName: "Ravi",
	LastName:  "Gupta",
	Email:     "c@gmail.com",
	Password:  "abc",
	PhoneNo:   "6464622",
	Role:      "Manager",
	Salary:    42413,
}}

var elements int = 3

func readLine() string {
	var input string
	fmt.Scanln(&input)

	return input
}

func main() {
	// Display the login screen
	fmt.Println("Hello and welcome to the management system")

	loginScreen()

	// Get the user's role
	fmt.Println("Please select your identity")
	fmt.Println("1 -> Admin")
	fmt.Println("2 -> Manager/Developer/Tester")

	role := readLine()

	// Perform operations based on the user's role
	switch role {
	case "1":
		// adminOperations()
		adminPanel()
	case "2":
		// employeeOperations()
		nonAdminPanel()
	default:
		fmt.Println("Invalid Input")
		os.Exit(1)
	}
}

func loginScreen() {
	fmt.Println("Employee Management System")
	fmt.Println("-------------------------")

	fmt.Print("Username: ")
	username := readLine()

	fmt.Print("Password: ")
	password := readLine()

	// TODO: Validate the username and password

	// If the username and password are valid, then log the user in
	fmt.Println("username : ", username)
	fmt.Println("Password : ", password)
	fmt.Println("Login successful!")

	// adminOperations()
}

func adminPanel() {
	fmt.Println("Welcome to the Elite Admin Panel 🎊")
	for true {
		fmt.Println("Please Enter your Input to Proceed...")
		fmt.Println("1 -> ADD an Employee")
		fmt.Println("2 -> DELETE an Employee")
		fmt.Println("3 -> UPDATE an Employee")
		fmt.Println("4 -> VIEW an Employee Details")
		fmt.Println("5 -> LIST all Employee")
		fmt.Println("6 -> Print all Employee in the shorted order of there salary")
		fmt.Println("Press ctrl + c to exit the program")
		operation := readLine()

		switch operation {
		case "1":
			addEmployee()
		case "2":
			deleteEmployee()
		case "3":
			updateEmployee()
		case "4":
			viewEmployeeDetails()
		case "5":
			listEmployees()
		case "6":
			sortedPrintEmployee()
		default:
			fmt.Println("Invalid Input")
			os.Exit(1)
			break
		}

	}

}

func addEmployee() {
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
	elements += 1
	employee := Employee{
		ID:        elements,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		PhoneNo:   phoneNo,
		Role:      role,
		Salary:    salary,
	}

	// Add the employee to the employee splice.
	employees = append(employees, employee)

	// Display a message to the user.
	fmt.Println("Employee added successfully!")

	fmt.Println("List of users")
	for _, employee := range employees {
		fmt.Println(employee.FirstName)
	}
}

func deleteEmployee() {
	fmt.Println("Please provide the id of the Person you want delete")
	var _id int
	fmt.Scanln(&_id)
	// Find the index of the employee with the specified ID.
	index := -1
	for i, employee := range employees {
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
	employees = append(employees[:index], employees[index+1:]...)
	elements -= 1
	fmt.Println("List of left out employees")
	for _, employee := range employees {
		fmt.Println(employee.FirstName)
	}

}

func updateEmployee() {
	fmt.Println("Please provide me the id of the user to update")
	var _id int
	fmt.Scanln(&_id)

	index := -1
	var updatedEmp Employee
	for i, employee := range employees {
		if employee.ID == _id {
			index = i
			updatedEmp = employee
			fmt.Println("So the person with id is below")
			fmt.Println(employee.FirstName)
			fmt.Println(employee.LastName)
			fmt.Println(employee.Email)
			fmt.Println(employee.PhoneNo)
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

	fmt.Println("Get the employee's Phone.")
	phoneNo := readLine()

	fmt.Println("Get the employee's role.")
	role := readLine()

	updatedEmp.FirstName = firstName
	updatedEmp.LastName = lastName
	updatedEmp.Email = email
	updatedEmp.PhoneNo = phoneNo
	updatedEmp.Role = role

	employees[index] = updatedEmp

	return
}

func viewEmployeeDetails() {
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

func listEmployees() {
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

func sortedPrintEmployee() {

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

func nonAdminPanel() {
	fmt.Println("Welcome to the Manager/Developer/Tester Panel 🎊")
	fmt.Println("Please Enter your ID ")
	var _id int
	fmt.Scanln(&_id)

	for true {
		fmt.Println("Please Enter your Input to Proceed...")
		fmt.Println("1 -> View your Details")
		fmt.Println("2 -> Update your details")
		fmt.Println("3 -> view other Employees details")
		fmt.Println("Press ctrl + c to exit the program")
		operation := readLine()

		switch operation {
		case "1":
			viewEmployeeDetailsNonAdmin(_id)
		case "2":
			updateEmployeeNonAdmin(_id)
		case "3":
			viewNonAdminDetails()

		default:
			fmt.Println("Invalid Input")
			os.Exit(1)
			break
		}

	}

}

func viewEmployeeDetailsNonAdmin(_id int) {

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

func updateEmployeeNonAdmin(_id int) {

	index := -1
	var updatedEmp Employee
	for i, employee := range employees {
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

	employees[index] = updatedEmp

	return
}

func viewNonAdminDetails() {
	for _, employee := range employees {
		if employee.Role != "Admin" {
			fmt.Println("So the Employee id is", employee.ID)
			fmt.Println(employee.FirstName)
			fmt.Println(employee.LastName)
			fmt.Println(employee.Email)
			fmt.Println(employee.Role)
		}
	}

}
