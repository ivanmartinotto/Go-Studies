package main

import "fmt"

func main() {
	var t any

	fmt.Printf("Hello " + "World!\n")

	fmt.Println(1 + 1)

	fmt.Println(3.14)

	fmt.Println(true, false)

	fmt.Printf("%+v\n", []int{1, 2, 3})

	fmt.Printf("%+v\n", t)
}
