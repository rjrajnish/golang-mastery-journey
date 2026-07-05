package main

import (
	"fmt"
	"sync"
	"time"
)

// ----------------------------------------
// Example Function
// ----------------------------------------

func printNumbers() {

	for i := 1; i <= 5; i++ {

		fmt.Println("Number:", i)

		time.Sleep(300 * time.Millisecond)

	}

}

func printLetters() {

	letters := []string{"A", "B", "C", "D", "E"}

	for _, letter := range letters {

		fmt.Println("Letter:", letter)

		time.Sleep(300 * time.Millisecond)

	}

}

func downloadFile(file string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Downloading:", file)

	time.Sleep(2 * time.Second)

	fmt.Println("Completed:", file)

}

func main() {

	fmt.Println("========================================")
	fmt.Println("     Go Goroutines - Complete Guide")
	fmt.Println("========================================")

	// ----------------------------------------
	// Example 1
	// Normal Function
	// ----------------------------------------

	fmt.Println("\n1. Normal Function")

	printNumbers()

	// ----------------------------------------
	// Example 2
	// First Goroutine
	// ----------------------------------------

	fmt.Println("\n2. Goroutine")

	go printLetters()

	time.Sleep(2 * time.Second)

	// ----------------------------------------
	// Example 3
	// Multiple Goroutines
	// ----------------------------------------

	fmt.Println("\n3. Multiple Goroutines")

	go printNumbers()

	go printLetters()

	time.Sleep(3 * time.Second)

	// ----------------------------------------
	// Example 4
	// WaitGroup
	// ----------------------------------------

	fmt.Println("\n4. WaitGroup")

	var wg sync.WaitGroup

	files := []string{
		"image.png",
		"video.mp4",
		"document.pdf",
	}

	for _, file := range files {

		wg.Add(1)

		go downloadFile(file, &wg)

	}

	wg.Wait()

	fmt.Println("All Downloads Completed")

	// ----------------------------------------
	// Example 5
	// Anonymous Goroutine
	// ----------------------------------------

	fmt.Println("\n5. Anonymous Goroutine")

	var wait sync.WaitGroup

	wait.Add(1)

	go func() {

		defer wait.Done()

		fmt.Println("Hello from Anonymous Goroutine")

	}()

	wait.Wait()

	fmt.Println("\nProgram Finished")

}