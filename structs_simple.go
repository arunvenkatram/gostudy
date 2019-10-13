package main

import "fmt"

type fruits struct {
	colour    string
	taste     string
	available bool
}

func main() {
	apple := fruits{
		colour:    "red",
		taste:     "sweet",
		available: false,
	}
	fmt.Println(apple)
	fmt.Println(apple.colour)
}
