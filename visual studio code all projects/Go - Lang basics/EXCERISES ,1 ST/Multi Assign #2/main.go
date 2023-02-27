package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Multi Assign #2
//
//  1. Assign the correct values to the variables
//     to match to the EXPECTED OUTPUT below
//
//  2. Print the variables
//
// HINT
//  Use multiple Println calls to print each sentence.
//
// EXPECTED OUTPUT
//  Air is good on Mars
//  It's true
//  It is 19.5 degrees
// ---------------------------------------------------------

func main() {
	// UNCOMMENT THE CODE BELOW:

	var (
		planet string
		isTrue bool
		temp   float64
	)

	// ADD YOUR CODE BELOW
	planet, temp, isTrue = "Mars", 19.5, true

	fmt.Println("air is good on ", planet)
	fmt.Println("It's ", isTrue)
	fmt.Println("it is ", temp, " degrees")

}
