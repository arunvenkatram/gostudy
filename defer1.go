package main
import "fmt"

func display(a int) {  
    fmt.Println(a)
}
func main() {  
    defer display(1)
    fmt.Println("testing print")
    defer display(2)
    defer display(3)
    fmt.Println(4)
}