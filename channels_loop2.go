package main

import "fmt"

func square(c chan int) {
	for i := 0; i <= 10; i++ {
		c <- i * i
	}
	close(c)
}
func main() {
	c := make(chan int)
	fmt.Println("Starting main function")
	go square(c)
	for x := range c {
		fmt.Println(x)
	}
	fmt.Println("Stopping main function")
}
