package main

import (
	"fmt"
	"runtime"
)

func multiplication(ch chan int) {
	for i := 0; i < 4; i++ { // if we make the loop execute more times than channel size (by setting i < 6 or so, then the numtilicationfunction actually does not stop.)
		num := <-ch
		fmt.Println(num * num)
	}
}
func main() {
	fmt.Println("main function started")
	c := make(chan int, 3)
	go multiplication(c)
	fmt.Println("Currently active goroutines: ", runtime.NumGoroutine())
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	fmt.Println("Currently active goroutines after 4 buffers: ", runtime.NumGoroutine())
	go multiplication(c)
	fmt.Println("Number of goroutines after calling go routing again: ", runtime.NumGoroutine())
	c <- 5
	c <- 6
	c <- 7
	c <- 8
	fmt.Println("Number of goroutines after passing another 4 values to channel: ", runtime.NumGoroutine())
	fmt.Println("main function stopped")
}
