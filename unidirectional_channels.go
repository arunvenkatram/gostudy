package main

import "fmt"

//Unidirectional channels are mainly for safety concerns.
func hello(readch <-chan string) {
	fmt.Println("Hello ", <-readch, "!")
}
func main() {
	ch := make(chan string)
	go hello(ch)
	ch <- "Arun"
	fmt.Println("Main function ended")
}
