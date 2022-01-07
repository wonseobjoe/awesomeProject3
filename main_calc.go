package main

import (
	add "awesomeProject3/add"
	sub "awesomeProject3/sub"
	"fmt"
)

func main() {
	fmt.Println("Addition = ", add.Add(10, 20))
	fmt.Println("Subtraction = ", sub.Sub(10, 20))
}
