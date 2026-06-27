package main

import "fmt"

func main() {

	var ptr *int

	fmt.Println(ptr)

	if ptr == nil {
		fmt.Println("Pointer is nil")
	}

}