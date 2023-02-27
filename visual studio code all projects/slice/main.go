package main

import "fmt"

func main() {

	var arr []int = []int{1, 2, 3, 4}

	arr2 := arr

	arr[3] = 55
	fmt.Println(arr)
	fmt.Println(arr2)

	// capacity
	fmt.Println(cap(arr))

	// 2 nd way of declaration
	// size is 2 and capacity is 4
	var arr4 []int = make([]int, 2, 4)

	fmt.Println(len(arr4))
	fmt.Println(cap(arr4))
	// append use to add new elements to another slice

	arr2 = append(arr, 1000)

	arr3 := append(arr, arr2...)
	fmt.Println(arr3)

}
