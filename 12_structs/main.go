package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday(age int) {
	p.age = age
}

// getMarried method (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Anthony", lastName: "Fox", city: "Emerson", gender: "m", age: 25}

	// Alt
	person2 := Person{"Stephanie", "Rose", "Emerson", "f", 25}

	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday(34)
	person2.getMarried("Fox")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
