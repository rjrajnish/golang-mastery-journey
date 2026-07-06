 package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// ----------------------------------------
// Custom Function (Simulating Package Logic)
// ----------------------------------------

func Add(a, b int) int {
	return a + b
}

func Welcome(name string) string {
	return fmt.Sprintf("Welcome %s to Go!", name)
}

func main() {

	fmt.Println("========================================")
	fmt.Println(" Go Packages & Modules - Complete Guide")
	fmt.Println("========================================")

	// ----------------------------------------
	// Example 1: fmt Package
	// ----------------------------------------

	fmt.Println("\n1. fmt Package")

	fmt.Println("Hello, Golang!")

	// ----------------------------------------
	// Example 2: math Package
	// ----------------------------------------

	fmt.Println("\n2. math Package")

	fmt.Println("Square Root of 64:", math.Sqrt(64))
	fmt.Println("Power 2^5:", math.Pow(2, 5))

	// ----------------------------------------
	// Example 3: strings Package
	// ----------------------------------------

	fmt.Println("\n3. strings Package")

	name := "golang"

	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.Title(name))
	fmt.Println(strings.Contains(name, "go"))

	// ----------------------------------------
	// Example 4: time Package
	// ----------------------------------------

	fmt.Println("\n4. time Package")

	fmt.Println("Current Time:", time.Now())

	// ----------------------------------------
	// Example 5: Custom Function
	// ----------------------------------------

	fmt.Println("\n5. Custom Function")

	fmt.Println("Sum:", Add(10, 20))

	// ----------------------------------------
	// Example 6: Another Custom Function
	// ----------------------------------------

	fmt.Println("\n6. Welcome Function")

	fmt.Println(Welcome("Rajnish"))

	// ----------------------------------------
	// Example 7: Multiple Imports
	// ----------------------------------------

	fmt.Println("\n7. Multiple Packages Working Together")

	text := strings.ToUpper("go programming")

	fmt.Println(text)
	fmt.Println("Length:", len(text))

	// ----------------------------------------
	// Example 8: Package Alias
	// ----------------------------------------

	fmt.Println("\n8. Package Usage")

	fmt.Printf("Pi Value : %.2f\n", math.Pi)

	// ----------------------------------------
	// Example 9: Constants from Packages
	// ----------------------------------------

	fmt.Println("\n9. Constants")

	fmt.Println("Max Float:", math.MaxFloat64)

	// ----------------------------------------
	// Example 10: Why Packages Matter
	// ----------------------------------------

	fmt.Println("\n10. Reusable Functions")

	fmt.Println(Add(100, 200))

}