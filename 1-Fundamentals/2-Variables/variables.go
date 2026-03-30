package main

import "fmt"

func main() {

	// Variable declaration and assignment (long version)
	var greeting string // zero-value is an empty string
	greeting = "Hello, World!"

	fmt.Println(greeting)

	var count int
	count = 10
	fmt.Println(count)

	var isRunning bool
	isRunning = true
	fmt.Println(isRunning)

	var firstName, lastName string
	firstName = "John"
	lastName = "Doe"
	fmt.Printf("Full Name: %s %s\n", firstName, lastName)

	// Variable declaration and assignment (short version)
	email := "test@test.com" // The type of 'email' is inferred to be string
	fmt.Println(email)

	age := 30 // The type of 'age' is inferred to be int
	fmt.Println(age)
}
