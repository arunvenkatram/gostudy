package main

import (
	"fmt"
	"runtime"
)

func square(c chan int) {
	fmt.Println("Square function started")
	sqnum := <-c
	c <- sqnum * sqnum
}
func cube(c chan int) {
	fmt.Println("Cube function strted")
	cubenum := <-c
	//Here, c is actually a bidirectional channel.  Actually, all channels in this code are bidirectional
	c <- cubenum * cubenum * cubenum
}
func main() {
	fmt.Println("Main function started")
	squarech := make(chan int)
	cubech := make(chan int)
	go square(squarech)
	go cube(cubech)
	numbertopass := 3
	fmt.Println("Sending value to square routine")
	squarech <- numbertopass
	fmt.Println("Main funcion resuming")
	fmt.Println("Sending value to cube routing")
	cubech <- numbertopass
	fmt.Println("Resuming main")
	fmt.Println("Fetching values from channels")
	fmt.Println("Number of routines before fetching values: ", runtime.NumGoroutine())
	sqval, cubeval := <-squarech, <-cubech
	fmt.Println("Number of routines after clearing channels: ", runtime.NumGoroutine())
	fmt.Println("square and cube values are: ", sqval, cubeval)
}
