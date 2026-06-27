package main

import "fmt"

func main() {

	fmt.Println("===================================")
	fmt.Println("        Go Slices Examples")
	fmt.Println("===================================")

	// ----------------------------------------
	// Example 1: Creating a Slice
	// ----------------------------------------

	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println("\n1. Creating a Slice")
	fmt.Println(numbers)

	// ----------------------------------------
	// Example 2: Length
	// ----------------------------------------

	fmt.Println("\n2. Length")
	fmt.Println("Length:", len(numbers))

	// ----------------------------------------
	// Example 3: Capacity
	// ----------------------------------------

	fmt.Println("\n3. Capacity")
	fmt.Println("Capacity:", cap(numbers))

	// ----------------------------------------
	// Example 4: Append
	// ----------------------------------------

	numbers = append(numbers, 60)

	fmt.Println("\n4. Append")
	fmt.Println(numbers)

	// ----------------------------------------
	// Example 5: Slice Expression
	// ----------------------------------------

	sub := numbers[1:4]

	fmt.Println("\n5. Slice Expression")
	fmt.Println(sub)

	// ----------------------------------------
	// Example 6: Copy
	// ----------------------------------------

	copySlice := make([]int, len(numbers))

	copy(copySlice, numbers)

	fmt.Println("\n6. Copy Slice")
	fmt.Println(copySlice)

	// ----------------------------------------
	// Example 7: Modify Original Slice
	// ----------------------------------------

	numbers[0] = 999

	fmt.Println("\n7. Original Slice")
	fmt.Println(numbers)

	fmt.Println("Copied Slice")
	fmt.Println(copySlice)

	// ----------------------------------------
	// Example 8: Range Loop
	// ----------------------------------------

	fmt.Println("\n8. Range Loop")

	for index, value := range numbers {

		fmt.Printf("Index: %d Value: %d\n", index, value)

	}

	// ----------------------------------------
	// Example 9: Nil Slice
	// ----------------------------------------

	var empty []int

	fmt.Println("\n9. Nil Slice")

	fmt.Println(empty)

	fmt.Println("Length:", len(empty))

	fmt.Println("Capacity:", cap(empty))

	fmt.Println("Is Nil:", empty == nil)

	// ----------------------------------------
	// Example 10: make()
	// ----------------------------------------

	users := make([]string, 3)

	users[0] = "Rajnish"

	users[1] = "Rahul"

	users[2] = "Amit"

	fmt.Println("\n10. make()")
	fmt.Println(users)

	// ----------------------------------------
	// Example 11: Remove Element
	// ----------------------------------------

	index := 2

	numbers = append(numbers[:index], numbers[index+1:]...)

	fmt.Println("\n11. Remove Element")
	fmt.Println(numbers)

}