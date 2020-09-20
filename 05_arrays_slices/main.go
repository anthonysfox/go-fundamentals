package main

import "fmt"

func main() {
	// Arrays
	// var fruitArray [2]string

	// // Assign Values
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	// Declare and assign
	// Sets the length of the array as well
	// fruitArray := [2]string{"Apple", "Orange"}

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
	// Slice the array
	fmt.Println(fruitSlice[1:2])
}
