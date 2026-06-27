package main

import "fmt"

func main() {

	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}

	}()

	panic("unexpected error occurred")
}