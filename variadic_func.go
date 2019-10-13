package main

import "fmt"

func FindBiggestNumber(numbers ...int) int {
	big := numbers[0]
	for _, i := range numbers {
		if i > big {
			big = i
		}
	}
	return big
}
func main() {
	biggestnum := FindBiggestNumber(1, 23, 43, 200, 66, 23, 54, 22)
	fmt.Println(biggestnum)
}
