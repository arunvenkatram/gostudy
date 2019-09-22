package main

import (
	"fmt"
)

func main() {
	arun := [...]int{1, 2, 3}
	fmt.Println(arun)
	slice := arun[1:2]
	fmt.Println("slice is:", slice)
	fmt.Println("Length of slice is: ", len(slice))
	fmt.Println("Appending slice")
	slice = append(slice, int(4), int(5), int(6))
	fmt.Println(slice)
	fmt.Println("Length of slice now is, ", len(slice))
	fmt.Println("Length of array now is, ", len(arun))
	fmt.Println("Capacity of parent array is: ", cap(slice))
	fmt.Println("Printing parent array: ", arun)
}
