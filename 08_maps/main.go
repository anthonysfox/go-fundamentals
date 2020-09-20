package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign key valyes
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare map and key values
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	fmt.Println(emails)

	// delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
