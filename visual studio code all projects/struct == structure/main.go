package main

import "fmt"

// struct is similiar as class but in struct we have only values(property)
// user defined datatype
type Employee struct {
	EmpId       int
	EmpName     string
	EmpMobileNo []string
}

func main() {

	emp := Employee{
		EmpId:       12,
		EmpName:     "raj",
		EmpMobileNo: []string{"98454646661", "44654544444"},
	}
	//values will be copied in struct not address
	emp2 := emp

	// if u want struct to be address base --> use &
	//emp2 = &emp // bad approach

	emp.EmpName = "ravi"
	fmt.Println(emp)
	fmt.Println(emp2)
}
