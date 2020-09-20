package main

import "fmt"

func main() {
	// Using var
	// var name string = "Anthony"
	var age int = 37
	const isCool = true

	// Shorthand for variable
	// name := "Anthony"
	// size := 1.3

	// More shorthand for above
	name, size := "Anthony", 1.9

	fmt.Println(name, age, isCool)

	// %T is the format representing getting the type of a the value
	// Checkout https://golang.org/pkg/fmt/
	fmt.Printf("%T\n", size)
}
