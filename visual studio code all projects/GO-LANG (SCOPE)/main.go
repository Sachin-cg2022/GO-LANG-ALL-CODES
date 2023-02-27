package main

import (
	"fmt"
	"myapp/packageone" // always remember bend rgt side (/)// don't forget to use
	// go mod init myapp do everytime from now on
)

// package level variables are avilable everywhere
//in that package,where package are declared
var one = "one gffgf " // applicable for both func

func main() {
	//var one = "this is a block level variable "// u should not do ,this is variable shadowing
	//where level var name is same
	//as func var name

	var something = "this is a block level variable" //always main func variable
	// execute first
	fmt.Println(something)

	funcOne()
	// package variable exporting
	newString := packageone.PublicVar
	fmt.Println("from packageone ", newString)

	packageone.Exported()

}

func funcOne() {
	//var one = "one" //- by creating a variable

	// if i don't assign variable we can use funcOne(one string)
	// then also no issue i am using package level variable
	fmt.Println(one)
}
