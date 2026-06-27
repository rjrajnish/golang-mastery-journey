package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func updateUser(user *User) {

	user.Name = "Rajnish Pandey"

	user.Age = 30

}

func main() {

	user := User{
		Name: "Raj",
		Age: 25,
	}

	updateUser(&user)

	fmt.Println(user)

}