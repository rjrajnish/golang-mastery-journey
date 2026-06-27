package main

import "fmt"
func printArray(arr [3]int) {

	arr[0] = 999

	fmt.Println("Inside Function:", arr)

}

func updateArray(arr *[3]int) {

	arr[0] = 999

}
func main() {

	fmt.Println("========================================")
	fmt.Println("Go Arrays - Complete Examples")
	fmt.Println("========================================")

	// ----------------------------------------
	// Example 1: Declare an Array
	// ----------------------------------------

	var numbers [5]int

	fmt.Println("\n1. Empty Array")
	fmt.Println(numbers)

	// ----------------------------------------
	// Example 2: Initialize Array
	// ----------------------------------------

	marks := [5]int{80, 75, 90, 88, 95}

	fmt.Println("\n2. Initialized Array")
	fmt.Println(marks)

	// ----------------------------------------
	// Example 3: Access Elements
	// ----------------------------------------

	fmt.Println("\n3. Access Elements")

	fmt.Println("First:", marks[0])
	fmt.Println("Last :", marks[4])

	// ----------------------------------------
	// Example 4: Update Elements
	// ----------------------------------------

	marks[1] = 100

	fmt.Println("\n4. Updated Array")

	fmt.Println(marks)

	// ----------------------------------------
	// Example 5: Array Length
	// ----------------------------------------

	fmt.Println("\n5. Array Length")

	fmt.Println(len(marks))

	// ----------------------------------------
	// Example 6: Loop using for
	// ----------------------------------------

	fmt.Println("\n6. Loop using for")

	for i := 0; i < len(marks); i++ {
		fmt.Println(marks[i])
	}

	// ----------------------------------------
	// Example 7: Loop using range
	// ----------------------------------------

	fmt.Println("\n7. Loop using range")

	for index, value := range marks {
		fmt.Println(index, value)
	}

	// ----------------------------------------
	// Example 8: Multidimensional Array
	// ----------------------------------------

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("\n8. Two-Dimensional Array")

	fmt.Println(matrix)

	// ----------------------------------------
	// Example 9: Arrays are Value Types
	// ----------------------------------------

	a := [3]int{10, 20, 30}

	b := a

	b[0] = 100

	fmt.Println("\n9. Array Copy")

	fmt.Println("Original :", a)

	fmt.Println("Copied   :", b)

	// ----------------------------------------
	// Example 10: Pass Array to Function
	// ----------------------------------------

	fmt.Println("\n10. Pass By Value")

	printArray(a)

	fmt.Println("After Function:", a)

	// ----------------------------------------
	// Example 11: Pass Pointer to Array
	// ----------------------------------------

	fmt.Println("\n11. Pass Pointer")

	updateArray(&a)

	fmt.Println(a)

}