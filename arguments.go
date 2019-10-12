package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args        //  All the arguments
	argument := os.Args[1] // only the first argument
	fmt.Println(argument)
	fmt.Println(args)
}
