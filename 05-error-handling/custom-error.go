package main

import "fmt"

type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}

func validateAge(age int) error {

	if age < 18 {
		return ValidationError{
			Message: "age must be at least 18",
		}
	}

	return nil
}

func main() {

	err := validateAge(15)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Validation Passed")
}