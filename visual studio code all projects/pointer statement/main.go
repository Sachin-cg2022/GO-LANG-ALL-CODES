package main

import "fmt"

// type Employee struct {
// 	EmpId int
// }

// func main() {
// 	var e Employee
// 	var emp *Employee

// 	fmt.Println(emp)

// 	emp = new(Employee) // allocate memory location
// 	fmt.Println(emp)    // o/p--> &{0}-->address
// 	fmt.Println(e)      // o/p--> {0}

// 	fmt.Println(*emp) // o/p--> {0}-->de-refering (value)

// 	// to obtain the values
// 	fmt.Println((*emp).EmpId)

// 	fmt.Println(e.EmpId)

// }

func main() {
	var x int = 90
	var b *int = &x //address of x

	fmt.Println(x)
	fmt.Println(b) // o/p address of x//refering memory location

	fmt.Println(x)
	fmt.Println(*b) // o/p value of x//de-refering the memory location

}
