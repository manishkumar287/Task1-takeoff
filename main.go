package main

import (
	"ProjectGo/utils"
	"fmt"
	"os"
)


func readLine() string {
	var input string
	fmt.Scanln(&input)

	return input
}

func main() {
	// Display the login screen
	fmt.Println("Hello and welcome to the management system")

	loginScreen()
	selectRole()

	
}

func selectRole()  {
	// Get the user's role
	fmt.Println("Please select your identity")
	fmt.Println("1 -> Admin")
	fmt.Println("2 -> Manager/Developer/Tester")
	fmt.Println("Press ctrl + c or any other key to exit the program")

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
		fmt.Println("7 -> Go to previous Menu")
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
		case "7":
			selectRole()
		default:
			fmt.Println("Invalid Input")
			os.Exit(1)
			break
		}

	}

}

func addEmployee() {	
	utils.Add_user(&utils.EmployeesList)
}

func deleteEmployee() {
	utils.Delete_user(&utils.EmployeesList)
}

func updateEmployee() {
	utils.Update_user(&utils.EmployeesList)
}

func viewEmployeeDetails() {
	utils.View_userBy_id(utils.EmployeesList)

}

func listEmployees() {
	utils.View_allUsers(utils.EmployeesList)
}

func sortedPrintEmployee() {
	utils.Sorted_userBy_salary(utils.EmployeesList)
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
		fmt.Println("4 -> Go to Previous Menu")
		fmt.Println("Press ctrl + c to exit the program")
		operation := readLine()

		switch operation {
		case "1":
			viewEmployeeDetailsNonAdmin(_id)
		case "2":
			updateEmployeeNonAdmin(_id)
		case "3":
			viewNonAdminDetails()
		case "4":
			selectRole()

		default:
			fmt.Println("Invalid Input")
			os.Exit(1)
			break
		}

	}

}

func viewEmployeeDetailsNonAdmin(_id int) {

	utils.View_NonAdminEmployeesDetails(utils.EmployeesList,_id)

}

func updateEmployeeNonAdmin(_id int) {
	utils.NonAdmin_UpdateEmployee(&utils.EmployeesList,_id)
}

func viewNonAdminDetails() {
	utils.View_NonAdminDetails(utils.EmployeesList)
}
