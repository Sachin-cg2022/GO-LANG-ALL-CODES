package main

// ---------------------------------------------------------
// EXERCISE: Print Your Name
//
//  Get it from the first command-line argument
//
// INPUT
//  Call the program using your name
//
// EXPECTED OUTPUT
//  It should print your name
//
// EXAMPLE
//  go run main.go inanc
//
//    inanc
//
// BONUS: Make the output like this:
//
//  go run main.go inanc
//    Hi inanc
//    How are you?
// ---------------------------------------------------------

import (
	"fmt"
	"math"
)

func main() {
	// var b1 bool = true
	// b2 := false
	// fmt.Println("its", b1, b2)

	// data := "hello my friend"
	// fmt.Println(len(data))
	// var inter, floater, stringer = 2, 2.1, "HI"
	// fmt.Println(inter, floater, stringer)

	//d := FOUR
	//fmt.Println(d)

	fmt.Println(sqrt(2), sqrt(-4))
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	} else {
		return fmt.Sprint(math.Sqrt(x))
	}

}
