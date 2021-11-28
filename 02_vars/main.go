package main

import "fmt"

func main() {
	//var name = "Brad"
	// Shorthand
	//name := "Brad" //Can only use := inside a function
	//email := "brad@gmail.com"

	// Declare and assign two variables
	name, email := "Brad", "brad@gmail.com"

	size := 1.3
	var age = 37
	const isCool = true

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)
}
