package main

import (
	"errors"
	"fmt"
)

func readConfig() error {
	return errors.New("config file not found")
}

func startServer() error {

	err := readConfig()

	if err != nil {
		return fmt.Errorf("server startup failed: %w", err)
	}

	return nil
}

func main() {

	err := startServer()

	if err != nil {
		fmt.Println(err)
	}
}