package main

import "fmt"

type Salary struct {
	Basic, HRA float64
}

type Employee struct {
	FirstName  string
	MonthlySalary              []Salary
}

func (e Employee) EmpInfo() string {
	fmt.Println(e.FirstName, )
	
	var basic float64
	var hra float64
	for _, info := range e.MonthlySalary {
		basic = info.Basic
		hra    = info.HRA
	}
	fmt.Print(basic,"\n")
	fmt.Print(hra,"\n")
	return "----------------------"
}

func main() {

	e := Employee{
		FirstName: "Mark",
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
			},
		},
	}
	fmt.Println(e.EmpInfo())
}