package main

import "fmt"

func main() {

	// declaration
	//empSal := make(map[string]int)

	//declaration with initialization
	empSal := map[string]int{
		"Atul": 20000,
		"raj":  10000,
		"neha": 15000,
	}
	fmt.Println(empSal)
	// obtain particular key -- neha
	fmt.Println(empSal["neha"])

	//if i want to change the salary of raj
	empSal["raj"] = 18000
	fmt.Println(empSal["raj"])

	//if i want to check the length of given data
	fmt.Println(len(empSal))

	//if i want to add new member
	empSal["sachin"] = 90000
	fmt.Println(empSal["sachin"])
	fmt.Println(empSal)

	// if i want to delete particular key
	delete(empSal, "neha")
	fmt.Println(empSal)

	// but i don't know jap is really present or not :- i will get 0 answer
	// this is not good approach
	fmt.Println(empSal["jap"])

	//to check the avialibility of particular key
	// output_ false
	_, flag := empSal["jap"]
	fmt.Println(flag)

	// output_ true
	_, flag = empSal["raj"]
	fmt.Println(flag)

}
