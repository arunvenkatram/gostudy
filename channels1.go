package main

import (
	"fmt"
	"time"
)

func sendvalue(c chan int) {
	fmt.Println("Execuing sendvalue function")
	time.Sleep(1 * time.Second)
	c <- 10
	fmt.Println("Finished execution")
}
func main() {
	fmt.Println("executing main function")
	value := make(chan int, 5)
	defer close(value)
	go sendvalue(value)
	go sendvalue(value)
	printvalue := <-value
	fmt.Println("Channel value is: ", printvalue)
	time.Sleep(1 * time.Second)
}
