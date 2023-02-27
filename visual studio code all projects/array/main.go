package main

import "fmt"

func main() {

	arr := [...]int{1, 2, 13, 4, 5, 6, 7}

	fmt.Println(arr)

	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[4])

	fmt.Println(len(arr))

	fmt.Println(arr[1:6])

	fmt.Println(arr[:5])

	fmt.Println(arr[4:])

	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matrix)
	fmt.Println(len(matrix))
	fmt.Println(matrix[2][1])
	matrix[1][0] = 99
	fmt.Println(matrix)

}
