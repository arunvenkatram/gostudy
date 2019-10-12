package main

import "fmt"

//declared the structure named emp
type emp struct {
	name    string
	address string
	age     int
}

//function which accepts variable of emp type and prints name property
func display(e emp) {
	fmt.Println(e.name)
}

func main() {
	// declares a variable, empdata1, of the type emp
	var empdata1 emp
	//assign values to members of empdata1
	empdata1.name = "Arun"
	empdata1.address = "HBR layout, Bangalore"
	empdata1.age = 27

	//declares and assign values to variable empdata2 of type emp
	empdata2 := emp{"Venkat", "Ponicherry", 27}

	//prints the member name of empdata1 and empdata2 using display function
	display(empdata1)
	display(empdata2)
	fmt.Println(empdata1)
	fmt.Println(empdata2)
}
