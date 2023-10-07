package main

import (
	"ProjectGo/utils"
	"fmt"
	"log"
	"os"
)

func readLine() string {
	var input string
	fmt.Scanln(&input)

	return input
}

func main() {
	// Display the login screen
	log.Println("Hello and welcome to the management system")

	loginScreen()
	selectRole()

}

func selectRole() {
	// Get the user's role
	log.Println("Please select your identity")
	log.Println("1 -> Admin")
	log.Println("2 -> Manager/Developer/Tester")
	log.Println("Press ctrl + c or any other key to exit the program")

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
		log.Println("Invalid Input")
		os.Exit(1)
	}

}

func loginScreen() {
	log.Println("Employee Management System")
	log.Println("-------------------------")

	log.Print("Username: ")
	username := readLine()

	log.Print("Password: ")
	password := readLine()

	// TODO: Validate the username and password

	// If the username and password are valid, then log the user in
	log.Println("username : ", username)
	log.Println("Password : ", password)
	log.Println("Login successful!")
	log.Println("-------------------------")

	// adminOperations()
}

func adminPanel() {
	log.Println("Welcome to the Elite Admin Panel 🎊")
	for true {
		log.Println("Please Enter your Input to Proceed...")
		log.Println("1 -> ADD an Employee")
		log.Println("2 -> DELETE an Employee")
		log.Println("3 -> UPDATE an Employee")
		log.Println("4 -> VIEW an Employee Details")
		log.Println("5 -> LIST all Employee")
		log.Println("6 -> Print all Employee in the shorted order of there salary")
		log.Println("7 -> Go to previous Menu")
		log.Println("Press ctrl + c to exit the program")
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
			log.Println("Invalid Input")
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
	log.Println("Welcome to the Manager/Developer/Tester Panel 🎊")
	log.Println("Please Enter your ID ")
	var _id int
	fmt.Scanln(&_id)

	for true {
		log.Println("Please Enter your Input to Proceed...")
		log.Println("1 -> View your Details")
		log.Println("2 -> Update your details")
		log.Println("3 -> view other Employees details")
		log.Println("4 -> Go to Previous Menu")
		log.Println("Press ctrl + c to exit the program")
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
			log.Println("Invalid Input")
			os.Exit(1)
			break
		}

	}

}

func viewEmployeeDetailsNonAdmin(_id int) {

	utils.View_NonAdminEmployeesDetails(utils.EmployeesList, _id)

}

func updateEmployeeNonAdmin(_id int) {
	utils.NonAdmin_UpdateEmployee(&utils.EmployeesList, _id)
}

func viewNonAdminDetails() {
	utils.View_NonAdminDetails(utils.EmployeesList)
}
