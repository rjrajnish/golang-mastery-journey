package main

import "fmt"

func changeValue(number int) {

	number = 500

}

func changePointer(number *int) {

	*number = 500

}

func main() {

	value := 100

	changeValue(value)

	fmt.Println("Value Function:", value)

	changePointer(&value)

	fmt.Println("Pointer Function:", value)

}