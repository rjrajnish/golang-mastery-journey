package main

import "fmt"

func increaseAge(age *int) {

	*age = *age + 1

}

func main() {

	age := 25

	increaseAge(&age)

	fmt.Println("Age:", age)

}