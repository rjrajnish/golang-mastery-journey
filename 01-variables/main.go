package main

import "fmt"

func main() {
// basic variables
	var name string ="Rajnish"
	var age int=25
    var salary float64=9000.56
	var isDeveloper bool= true
	fmt.Println("Name : ",name)
	fmt.Println("Age : ", age)
	fmt.Println("Salary : ",salary)
	fmt.Println("Is This Developer :", isDeveloper)
	// short declarations
	city := "New Delhi"
	country := "India"

	fmt.Print("Location:",city,country)
	// multiple declarations
	new_name,new_age,new_salary,new_city,new_country := "Rajnish",25,9000.56,"noida","India"
	fmt.Print("multiple ",new_name,new_age,new_salary,new_city,new_country)
}