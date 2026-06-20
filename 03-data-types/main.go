package main

import "fmt"

func main() {

	// String
	var name string = "Rajnish Pandey"

	// Integer
	var age int = 25

	// Float
	var salary float64 = 50000.75

	// Boolean
	var isDeveloper bool = true

	// Byte (ASCII value)
	var grade byte = 'B'

	// Rune (Unicode character)
	var heart rune = '❤'

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Developer:", isDeveloper)
	fmt.Println("Grade:", string(grade))
	fmt.Println("Heart Symbol:", string(heart))

	fmt.Println("\n---- Type Information ----")

	fmt.Printf("name -> %T\n", name)
	fmt.Printf("age -> %T\n", age)
	fmt.Printf("salary -> %T\n", salary)
	fmt.Printf("isDeveloper -> %T\n", isDeveloper)
	fmt.Printf("grade -> %T\n", grade)
	fmt.Printf("heart -> %T\n", heart)
}