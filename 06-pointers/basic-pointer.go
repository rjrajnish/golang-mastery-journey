package main

import "fmt"

func main() {

	number := 100

	pointer := &number

	fmt.Println("Value:", number)

	fmt.Println("Memory Address:", pointer)

	fmt.Println("Value Using Pointer:", *pointer)

}