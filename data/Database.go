package data

var employees []Employee

func GetAllEmployees() []Employee {
	return employees
}

func SaveEmployee(newEmployee Employee) {
	employees = append(employees, newEmployee)
}

func Init() {

	employees = append(employees, Employee{Id: 1, Age: 51, Namn: "Stefan", City: "Test"})
	employees = append(employees, Employee{Id: 2, Age: 15, Namn: "Oliver", City: "Test"})
	employees = append(employees, Employee{Id: 3, Age: 21, Namn: "Josefine", City: "Test"})
}