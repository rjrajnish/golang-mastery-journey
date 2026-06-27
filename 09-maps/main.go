 package main

import "fmt"

func main() {

	fmt.Println("===================================")
	fmt.Println("      Go Maps - Complete Guide")
	fmt.Println("===================================")

	// ----------------------------------------
	// Example 1: Create a Map
	// ----------------------------------------

	student := map[string]string{
		"name":   "Rajnish",
		"course": "Go Programming",
	}

	fmt.Println("\n1. Create Map")
	fmt.Println(student)

	// ----------------------------------------
	// Example 2: Access Value
	// ----------------------------------------

	fmt.Println("\n2. Access Value")
	fmt.Println("Name:", student["name"])

	// ----------------------------------------
	// Example 3: Add New Key
	// ----------------------------------------

	student["city"] = "Ghaziabad"

	fmt.Println("\n3. Add New Key")
	fmt.Println(student)

	// ----------------------------------------
	// Example 4: Update Existing Value
	// ----------------------------------------

	student["course"] = "Advanced Go"

	fmt.Println("\n4. Update Value")
	fmt.Println(student)

	// ----------------------------------------
	// Example 5: Delete Key
	// ----------------------------------------

	delete(student, "city")

	fmt.Println("\n5. Delete Key")
	fmt.Println(student)

	// ----------------------------------------
	// Example 6: Check Key Exists
	// ----------------------------------------

	value, exists := student["course"]

	fmt.Println("\n6. Check Key Exists")

	if exists {
		fmt.Println("Course:", value)
	} else {
		fmt.Println("Course Not Found")
	}

	// ----------------------------------------
	// Example 7: Loop Through Map
	// ----------------------------------------

	fmt.Println("\n7. Loop Through Map")

	for key, value := range student {
		fmt.Printf("%s : %s\n", key, value)
	}

	// ----------------------------------------
	// Example 8: Length
	// ----------------------------------------

	fmt.Println("\n8. Map Length")
	fmt.Println(len(student))

	// ----------------------------------------
	// Example 9: Empty Map using make()
	// ----------------------------------------

	users := make(map[int]string)

	users[101] = "Rahul"
	users[102] = "Amit"

	fmt.Println("\n9. make()")
	fmt.Println(users)

	// ----------------------------------------
	// Example 10: Nested Map
	// ----------------------------------------

	employee := map[string]map[string]string{
		"emp1": {
			"name": "Ravi",
			"city": "Delhi",
		},
	}

	fmt.Println("\n10. Nested Map")
	fmt.Println(employee)

	// ----------------------------------------
	// Example 11: Nil Map
	// ----------------------------------------

	var products map[int]string

	fmt.Println("\n11. Nil Map")
	fmt.Println(products)

	if products == nil {
		fmt.Println("Map is nil")
	}

	// ----------------------------------------
	// Example 12: Iterate Keys Only
	// ----------------------------------------

	fmt.Println("\n12. Keys")

	for key := range student {
		fmt.Println(key)
	}
}