package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("today is 5th")
	case 10:
		fmt.Println("today is 10th")
	case 15:
		fmt.Println("today is 15th")
	case 21, 25:
		fmt.Println("today is 25th")
	case 30:
		fmt.Println("today is 30th")
	default:
		fmt.Printf("no data avialable")
	}

}
