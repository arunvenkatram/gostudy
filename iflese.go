package main

import (
	"fmt"
)

func main() {
	arun := 75
	venkat := 75
	if arun < venkat {
		fmt.Println("arun is smaller than venkat")
	} else if venkat > arun {
		fmt.Println("Venkat is smaller than arun")
	} else {
		fmt.Println("Both are equal")
	}
}
