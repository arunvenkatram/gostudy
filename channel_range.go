package main

import "fmt"

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	// Using range, we can read values from closed channels. In closed channel, value still remains in buffer.
	for numbers := range c {
		fmt.Println(numbers)
	}
}
