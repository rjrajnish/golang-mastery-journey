package main

import (
	"errors"
	"fmt"
)

func login(username string) error {

	if username == "" {
		return errors.New("username cannot be empty")
	}

	return nil
}

func main() {

	err := login("")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Login Successful")
}