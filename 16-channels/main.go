 package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("========================================")
	fmt.Println("      Go Channels - Complete Guide")
	fmt.Println("========================================")

	// ----------------------------------------
	// Example 1: Create Channel
	// ----------------------------------------

	fmt.Println("\n1. Create Channel")

	ch := make(chan string)

	go func() {
		ch <- "Hello from Goroutine"
	}()

	message := <-ch

	fmt.Println(message)

	// ----------------------------------------
	// Example 2: Integer Channel
	// ----------------------------------------

	fmt.Println("\n2. Integer Channel")

	numberChannel := make(chan int)

	go func() {
		numberChannel <- 100
	}()

	number := <-numberChannel

	fmt.Println(number)

	// ----------------------------------------
	// Example 3: Buffered Channel
	// ----------------------------------------

	fmt.Println("\n3. Buffered Channel")

	buffer := make(chan string, 2)

	buffer <- "Go"
	buffer <- "Programming"

	fmt.Println(<-buffer)
	fmt.Println(<-buffer)

	// ----------------------------------------
	// Example 4: Range Channel
	// ----------------------------------------

	fmt.Println("\n4. Range Over Channel")

	values := make(chan int)

	go func() {

		for i := 1; i <= 5; i++ {
			values <- i
		}

		close(values)

	}()

	for value := range values {

		fmt.Println(value)

	}

	// ----------------------------------------
	// Example 5: Directional Channel
	// ----------------------------------------

	fmt.Println("\n5. Directional Channel")

	directional := make(chan string)

	go sendMessage(directional)

	receiveMessage(directional)

	// ----------------------------------------
	// Example 6: WaitGroup + Channel
	// ----------------------------------------

	fmt.Println("\n6. WaitGroup")

	var wg sync.WaitGroup

	results := make(chan string)

	files := []string{
		"image.png",
		"video.mp4",
		"report.pdf",
	}

	for _, file := range files {

		wg.Add(1)

		go download(file, results, &wg)

	}

	go func() {

		wg.Wait()

		close(results)

	}()

	for result := range results {

		fmt.Println(result)

	}

	// ----------------------------------------
	// Example 7: Channel Blocking
	// ----------------------------------------

	fmt.Println("\n7. Channel Synchronization")

	done := make(chan bool)

	go func() {

		time.Sleep(2 * time.Second)

		fmt.Println("Task Completed")

		done <- true

	}()

	<-done

	fmt.Println("Main Function Finished")

}

func sendMessage(ch chan<- string) {

	ch <- "Directional Channel Example"

}

func receiveMessage(ch <-chan string) {

	fmt.Println(<-ch)

}

func download(file string, result chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()

	time.Sleep(time.Second)

	result <- file + " Downloaded"

}