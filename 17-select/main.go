package main

import (
	"fmt"
	"time"
)

func worker(name string, delay time.Duration, ch chan string) {

	time.Sleep(delay)

	ch <- fmt.Sprintf("%s completed", name)

}

func main() {

	fmt.Println("========================================")
	fmt.Println("     Go Select Statement Examples")
	fmt.Println("========================================")

	// ----------------------------------------
	// Example 1: Basic Select
	// ----------------------------------------

	fmt.Println("\n1. Basic Select")

	ch1 := make(chan string)

	go func() {

		time.Sleep(time.Second)

		ch1 <- "Hello from Channel"

	}()

	select {

	case message := <-ch1:
		fmt.Println(message)

	}

	// ----------------------------------------
	// Example 2: Multiple Channels
	// ----------------------------------------

	fmt.Println("\n2. Multiple Channels")

	emailChannel := make(chan string)
	smsChannel := make(chan string)

	go worker("Email", 2*time.Second, emailChannel)
	go worker("SMS", time.Second, smsChannel)

	select {

	case result := <-emailChannel:
		fmt.Println(result)

	case result := <-smsChannel:
		fmt.Println(result)

	}

	// ----------------------------------------
	// Example 3: Timeout
	// ----------------------------------------

	fmt.Println("\n3. Timeout")

	api := make(chan string)

	go func() {

		time.Sleep(3 * time.Second)

		api <- "API Response"

	}()

	select {

	case response := <-api:

		fmt.Println(response)

	case <-time.After(2 * time.Second):

		fmt.Println("Request Timed Out")

	}

	// ----------------------------------------
	// Example 4: Default Case
	// ----------------------------------------

	fmt.Println("\n4. Default Case")

	messageChannel := make(chan string)

	select {

	case msg := <-messageChannel:

		fmt.Println(msg)

	default:

		fmt.Println("No Messages Available")

	}

	// ----------------------------------------
	// Example 5: Continuous Select
	// ----------------------------------------

	fmt.Println("\n5. Continuous Select")

	ticker := time.NewTicker(time.Second)

	stop := time.After(5 * time.Second)

	for {

		select {

		case time := <-ticker.C:

			fmt.Println("Tick:", time.Format("15:04:05"))

		case <-stop:

			fmt.Println("Stopping...")

			ticker.Stop()

			return

		}

	}

}