package main

import "fmt"

func greet(c chan string) {
	var value string = <-c
	fmt.Println("Vaule in channel is ", value)
}
func main() {
	fmt.Println("Main function started.")
	c := make(chan string)
	go greet(c)
	c <- "Arun"
	// go greet(c)
	fmt.Println("Main function ended")
}
