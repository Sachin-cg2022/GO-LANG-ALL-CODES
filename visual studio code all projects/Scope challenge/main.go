package main

//declare a package level variable for the main package named myVar
//declare a block variable for the main function called blockVar
//declare a package level variable in the packageone package
//named packageVar
// create a exported function in packageone called PrintMe
//in the main function,print out the values of myVar, blockVar,
// and packageVar on one line using the PrintMe function in packageone
import (
	"myapp/packageone"
)

var myVar = "its a package level variable , "

func main() {
	var blockVar = "its a block level variable"

	packageone.PrintMe(myVar, blockVar)

}
