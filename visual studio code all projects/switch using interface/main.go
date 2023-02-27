package main

import "fmt"

func main() {
	var x interface{} = "progrank"
	switch x.(type) {
	case int:
		fmt.Println("hello")
	case float64:
		fmt.Println("hi")
	case string:
		fmt.Println("hey")
	default:
		fmt.Println("none")

	}
}
