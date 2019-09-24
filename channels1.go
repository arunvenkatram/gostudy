package main

import (
	"fmt"
	"time"
)

func sendvalue(c chan int) {
	fmt.Println("Execuing sendvalue function")
	time.Sleep(1 * time.Second)
	for i := 10; i < 30; i++ {
		c <- i
	}
	fmt.Println("Finished execution")
}
func main() {
	fmt.Println("executing main function")
	value := make(chan int, 100)
	// value <- 6
	defer close(value)
	go sendvalue(value)
	printvalue := <-value
	fmt.Println("Channel value is: ", printvalue)
	printvalue2 := <-value
	fmt.Println("2nd channel value is: ", printvalue2)
	time.Sleep(1 * time.Second)
}
