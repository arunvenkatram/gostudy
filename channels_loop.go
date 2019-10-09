package main

import "fmt"

func square(c chan int) {
	for i := 0; i <= 10; i++ {
		c <- i * i
	}
	close(c)
}
func main() {
	fmt.Println("Starting main function")
	c := make(chan int)
	go square(c)
	for {
		value, ok := <-c
		if ok == false {
			// Channel return 0 if channel is closed. Hence there will be 0 in output.
			fmt.Println(value, ok, "Channel ends here")
			break
		} else {
			fmt.Println(value, ok)
		}

	}
}
