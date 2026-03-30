package main

import "fmt"

func main() {

	temp := 25

	if temp > 30 {
		fmt.Println("greater than 30")
	} else {
		fmt.Println("not greater than 30")
	}

	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D")
	}

	// Example of using a map to check access
	userAcess := map[string]bool{
		"jane": true,
		"john": false,
	}

	hasAcess := userAcess["jane"]

	if hasAcess {
		fmt.Println("Jane has access")
	} else {
		fmt.Println("Jane does not have access")
	}

}
