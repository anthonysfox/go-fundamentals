package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5}

	// Loop through ids
	for i, id := range ids {
		// i is index and id is cooresponding id
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	// When not using the index use an underscore otherwis error will be thrown saying not being used
	for _, id := range ids {
		// i is index and id is cooresponding id
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	// Below is only accessing keys
	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
