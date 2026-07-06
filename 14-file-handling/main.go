 package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("========================================")
	fmt.Println("      Go File Handling - Complete Guide")
	fmt.Println("========================================")

	fileName := "example.txt"

	// ----------------------------------------
	// Example 1: Create File
	// ----------------------------------------

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n1. File Created Successfully")

	file.Close()

	// ----------------------------------------
	// Example 2: Write to File
	// ----------------------------------------

	err = os.WriteFile(fileName, []byte("Hello, Go File Handling!\n"), 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n2. Data Written Successfully")

	// ----------------------------------------
	// Example 3: Read Entire File
	// ----------------------------------------

	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n3. File Content")
	fmt.Println(string(data))

	// ----------------------------------------
	// Example 4: Append Data
	// ----------------------------------------

	file, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	file.WriteString("Learning Go is fun!\n")

	file.Close()

	fmt.Println("\n4. Data Appended")

	// ----------------------------------------
	// Example 5: Read Line by Line
	// ----------------------------------------

	file, err = os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n5. Read Line By Line")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	file.Close()

	// ----------------------------------------
	// Example 6: File Information
	// ----------------------------------------

	info, err := os.Stat(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n6. File Info")

	fmt.Println("Name :", info.Name())
	fmt.Println("Size :", info.Size(), "bytes")
	fmt.Println("Mode :", info.Mode())

	// ----------------------------------------
	// Example 7: Check File Exists
	// ----------------------------------------

	fmt.Println("\n7. File Exists")

	_, err = os.Stat(fileName)

	if err == nil {
		fmt.Println("Yes")
	}

	// ----------------------------------------
	// Example 8: Rename File
	// ----------------------------------------

	newName := "notes.txt"

	err = os.Rename(fileName, newName)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n8. File Renamed")

	// ----------------------------------------
	// Example 9: Delete File
	// ----------------------------------------

	err = os.Remove(newName)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n9. File Deleted Successfully")
}