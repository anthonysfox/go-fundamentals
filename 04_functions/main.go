package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

// Dont need first int since both are ints
func getSum(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(getSum(1, 2))
}
