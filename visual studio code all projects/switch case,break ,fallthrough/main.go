package main

import "fmt"

func main() {
	////////////////////////////// switch case

	switch i := 7; i {

	case 5, 6, 7:
		fmt.Println("hello", i)

		// fallthrough will make second case also get printed
		//o/p -->hello 7
		//hi 7
		
		fallthrough


		//////////////////////break
		//below will not be printed
		//fmt.Println("my friend")

	case 4, 3, 2:
		fmt.Println("hi", i)

	case 10:
		fmt.Println("hey", i)

	default:
		fmt.Println("bye")

	}
}
