package main

import "fmt"

type shape interface {
	area() float64
	volume() float64
}
type rectangle struct {
	width  float64
	height float64
	length float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) volume() float64 {
	return r.width * r.height * r.length
}

func main() {
	var s shape
	s = rectangle{5.0, 4.0, 3.0}
	r := rectangle{7.0, 8.0, 9.0}
	fmt.Println("Area of rectangle is :", s.area())
	fmt.Println("Volume of rectangle is :", s.volume())
	fmt.Println("Area of another rectangle is :", r.area())
	fmt.Println("Area of another rectangl is :", r.volume())
	fmt.Println("Are both the recangles same? ", r == s)
}
