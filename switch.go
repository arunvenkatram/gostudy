package main

import "fmt"

func main() {
	a, b := 1, 1
	c, d := 5, 7
	var sum int
	switch a + b {
	case 1:
		fmt.Println("Sum is 1")
	case 2:
		fmt.Println("Sum is 2")
		sum = c + d
		fmt.Println("c+d is: ", sum)
	case 3:
		fmt.Println("Sum is 3")
	default:
		fmt.Println("Printing default")
	}
}
