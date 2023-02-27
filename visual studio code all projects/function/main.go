package main

import "fmt"

///////////////////////return type
// func main() {

// 	fmt.Println(show())
// }
// func show() string {
// 	return "hi"
// }
// o/p--> hi

////////////////////no return type direct
// func main() {

// 	show()
// }
// func show() {
// 	fmt.Println("HELLO")
// }
// o/p-->hello

////////////////////////summation
// func main(){
// 	summation(2,5)
// }
// func summation(x int,y int){
// 	fmt.Println(x+y)
// }
// o/p-->7

//

// //////////////////////multiple value return statement

// func main() {
// 	add, sub, div := cals(5, 3)
// 	fmt.Println(add)
// 	fmt.Println(sub)
// 	fmt.Println(div)
// }

// func cals(x int, y int) (int, int, int) {
// 	return (x + y), (x - y), (x / y)
// }
// // o/p---> 8,2,1

// //////////////boolean o/p
// func main() {
// 	f := isEven(5)
// 	fmt.Println(f)

// }

// func isEven(x int) bool {
// 	return x%2 == 0
// }

// //o/p-->false

// /////////////////////////

// func main() {
// 	num := 4
// 	sum(&num)
// 	fmt.Println(num)
// }
// func sum(num *int) {
// 	*num = 40
// 	fmt.Println(*num)
// }
// // o/p -->40,40

// //////////////////////struct associated with func
// type Employee struct {
// 	EmpName string
// }

// func main() {
// 	emp := Employee{"progrank"}
// 	emp.showEmpDetails()

// }

// func (emp Employee) showEmpDetails() {
// 	fmt.Println(emp.EmpName)
// }
// //// o/p--> progrank

/////////////////anonaoymous function

// func main() {

// 	func() {
// 		fmt.Println("hello")
// 	}()
// }

// o/p --> hello

// func main() {

// 	f := func() {
// 		fmt.Println("hi")
// 	}

// 	f()
// }
//// o/p --> hi

func main() {

	f := func() string {
		return "bye"
	}

	fmt.Println(f())
}

/// o/p---> bye
