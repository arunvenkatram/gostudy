package main

import "fmt"

func main() {
	fmt.Println("Starting main")
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	fmt.Println("length of the channel is ", len(ch), "and capacity of channel is ", cap(ch))
}
