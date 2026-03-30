package main

import (
	"fmt"
	"time"
)

func main() {

	day := "Tuesday"
	fmt.Println("Today is", day)

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("It's the weekend!")
	case "Monday", "Tuesday":
		fmt.Println("Work days. Lots of meetings.")
	default:
		fmt.Println("Midweek days. Keep going!")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good Morning!")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good Evening")
	}

	//checking Variable types with switch case
	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Integer: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case bool:
			fmt.Printf("Bool: %t\n", v)
		default:
			fmt.Printf("Unknown Type: %T\n", v)
		}
	}

	checkType(21)
	checkType("Test")
	checkType(true)
	checkType(12.12)

}
