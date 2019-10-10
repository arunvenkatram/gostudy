package main

import (
	"fmt"
)

// in a buffered channel, during send data, current Go routine will not be blocked untill channel is full.
// Also, data will start flowing out only when channel limit is reached. This is because, main is not getting blocked and it continues execution.
func multiplication(c chan int) {
	for i := 0; i <= 10; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main started")
	c := make(chan int, 3)
	go multiplication(c)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	// When we specify Sleep time, multiplication goroutine gets enough time to get executed.
	// If we pump in one more vaulue, then channel buffer gets filled up and main go routing gets blocked and all the value exists at multiplication goroutine.
	// time.Sleep(2 * time.Second)
	fmt.Println("main completed")
}
