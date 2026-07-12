 package main

import (
	"fmt"
	"sync"
	"time"
)

// Shared Resource
var counter int

// Mutex
var mutex sync.Mutex

// RWMutex
var rwMutex sync.RWMutex

func increment(wg *sync.WaitGroup) {

	defer wg.Done()

	mutex.Lock()

	counter++

	mutex.Unlock()

}

func readCounter(wg *sync.WaitGroup) {

	defer wg.Done()

	rwMutex.RLock()

	fmt.Println("Current Counter:", counter)

	rwMutex.RUnlock()

}

func writeCounter(value int, wg *sync.WaitGroup) {

	defer wg.Done()

	rwMutex.Lock()

	counter = value

	rwMutex.Unlock()

}

func main() {

	fmt.Println("========================================")
	fmt.Println("  Go Mutex & Race Condition Examples")
	fmt.Println("========================================")

	// ----------------------------------------
	// Example 1
	// Race Condition (Simulation)
	// ----------------------------------------

	fmt.Println("\n1. Race Condition Example")

	counter = 0

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {

		wg.Add(1)

		go increment(&wg)

	}

	wg.Wait()

	fmt.Println("Final Counter:", counter)

	// ----------------------------------------
	// Example 2
	// Mutex Lock
	// ----------------------------------------

	fmt.Println("\n2. Mutex Example")

	counter = 0

	wg = sync.WaitGroup{}

	for i := 0; i < 1000; i++ {

		wg.Add(1)

		go increment(&wg)

	}

	wg.Wait()

	fmt.Println("Counter:", counter)

	// ----------------------------------------
	// Example 3
	// RWMutex Read
	// ----------------------------------------

	fmt.Println("\n3. RWMutex Read")

	wg = sync.WaitGroup{}

	for i := 0; i < 3; i++ {

		wg.Add(1)

		go readCounter(&wg)

	}

	wg.Wait()

	// ----------------------------------------
	// Example 4
	// RWMutex Write
	// ----------------------------------------

	fmt.Println("\n4. RWMutex Write")

	wg = sync.WaitGroup{}

	wg.Add(1)

	go writeCounter(500, &wg)

	wg.Wait()

	fmt.Println("Updated Counter:", counter)

	// ----------------------------------------
	// Example 5
	// Multiple Readers
	// ----------------------------------------

	fmt.Println("\n5. Multiple Readers")

	wg = sync.WaitGroup{}

	for i := 0; i < 5; i++ {

		wg.Add(1)

		go readCounter(&wg)

	}

	wg.Wait()

	// ----------------------------------------
	// Example 6
	// Simulate Shared Resource
	// ----------------------------------------

	fmt.Println("\n6. Shared Resource")

	counter = 100

	wg = sync.WaitGroup{}

	wg.Add(2)

	go func() {

		defer wg.Done()

		mutex.Lock()

		counter += 50

		time.Sleep(time.Second)

		mutex.Unlock()

	}()

	go func() {

		defer wg.Done()

		mutex.Lock()

		counter -= 20

		mutex.Unlock()

	}()

	wg.Wait()

	fmt.Println("Final Value:", counter)

}