package main

import "fmt"

func sum(numbers ...int) int {

	total := 0

	for _, num := range numbers {
		total += num
	}

	return total
}

func main() {

	fmt.Println(sum(10, 20))
	fmt.Println(sum(10, 20, 30))
	fmt.Println(sum(10, 20, 30, 40))
}