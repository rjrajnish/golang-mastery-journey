 package main

import "fmt"

// ----------------------------------------
// Struct Definition
// ----------------------------------------

type Student struct {
	ID     int
	Name   string
	Age    int
	Course string
}

// Nested Struct
type Address struct {
	City    string
	State   string
	Country string
}

type Employee struct {
	ID      int
	Name    string
	Salary  float64
	Address Address
}

func main() {

	fmt.Println("========================================")
	fmt.Println("      Go Structs - Complete Guide")
	fmt.Println("========================================")

	// ----------------------------------------
	// Example 1: Create Struct
	// ----------------------------------------

	student := Student{
		ID:     101,
		Name:   "Rajnish",
		Age:    25,
		Course: "Go Programming",
	}

	fmt.Println("\n1. Create Struct")
	fmt.Println(student)

	// ----------------------------------------
	// Example 2: Access Fields
	// ----------------------------------------

	fmt.Println("\n2. Access Fields")
	fmt.Println("Name:", student.Name)
	fmt.Println("Course:", student.Course)

	// ----------------------------------------
	// Example 3: Update Fields
	// ----------------------------------------

	student.Course = "Advanced Go"

	fmt.Println("\n3. Update Field")
	fmt.Println(student)

	// ----------------------------------------
	// Example 4: Anonymous Struct
	// ----------------------------------------

	user := struct {
		Name string
		Role string
	}{
		Name: "Rahul",
		Role: "Admin",
	}

	fmt.Println("\n4. Anonymous Struct")
	fmt.Println(user)

	// ----------------------------------------
	// Example 5: Nested Struct
	// ----------------------------------------

	employee := Employee{
		ID:     1,
		Name:   "Amit",
		Salary: 75000,
		Address: Address{
			City:    "Delhi",
			State:   "Delhi",
			Country: "India",
		},
	}

	fmt.Println("\n5. Nested Struct")
	fmt.Println(employee)

	// ----------------------------------------
	// Example 6: Slice of Structs
	// ----------------------------------------

	students := []Student{
		{1, "Raj", 24, "Go"},
		{2, "Rahul", 23, "React"},
		{3, "Amit", 26, "Node.js"},
	}

	fmt.Println("\n6. Slice of Structs")

	for _, s := range students {
		fmt.Println(s)
	}

	// ----------------------------------------
	// Example 7: Struct Comparison
	// ----------------------------------------

	s1 := Student{1, "Raj", 24, "Go"}
	s2 := Student{1, "Raj", 24, "Go"}

	fmt.Println("\n7. Struct Comparison")
	fmt.Println(s1 == s2)

	// ----------------------------------------
	// Example 8: Pointer to Struct
	// ----------------------------------------

	ptr := &student

	ptr.Name = "Rajnish Pandey"

	fmt.Println("\n8. Pointer to Struct")
	fmt.Println(student)

	// ----------------------------------------
	// Example 9: Zero Value Struct
	// ----------------------------------------

	var newStudent Student

	fmt.Println("\n9. Zero Value Struct")
	fmt.Println(newStudent)

	// ----------------------------------------
	// Example 10: Function with Struct
	// ----------------------------------------

	displayStudent(student)

}

func displayStudent(s Student) {

	fmt.Println("\n10. Function with Struct")
	fmt.Printf("ID      : %d\n", s.ID)
	fmt.Printf("Name    : %s\n", s.Name)
	fmt.Printf("Age     : %d\n", s.Age)
	fmt.Printf("Course  : %s\n", s.Course)

}