package utils

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

type Employees Employee

var EmployeesList = []Employee{{
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
