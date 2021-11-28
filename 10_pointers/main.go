//open a terminal and then: go run main.go

package main

import "fmt"

func main() {
    a := 5
    b := &a // assign address of a to b

    fmt.Println(a, b) //prints an int and an address of an int
    fmt.Printf("%T\n", b) //print the type of b

    // Use * to read val from address
    fmt.Println(*b)
    fmt.Println(*&a)

    // Change val with pointer
    *b = 10 //This will change the value of a
    fmt.Println(a)

}