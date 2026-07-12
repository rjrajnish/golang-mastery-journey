 package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {

	for {

		select {

		case <-ctx.Done():

			fmt.Println("Worker Stopped")
			fmt.Println("Reason:", ctx.Err())

			return

		default:

			fmt.Println("Working...")

			time.Sleep(time.Second)

		}

	}

}

func downloadFile(ctx context.Context) {

	select {

	case <-time.After(5 * time.Second):

		fmt.Println("Download Completed")

	case <-ctx.Done():

		fmt.Println("Download Cancelled")

	}

}

func main() {

	fmt.Println("========================================")
	fmt.Println("      Go Context - Complete Guide")
	fmt.Println("========================================")

	// ----------------------------------------
	// Example 1
	// Background Context
	// ----------------------------------------

	fmt.Println("\n1. Background Context")

	ctx := context.Background()

	fmt.Println(ctx)

	// ----------------------------------------
	// Example 2
	// WithCancel
	// ----------------------------------------

	fmt.Println("\n2. WithCancel")

	cancelCtx, cancel := context.WithCancel(context.Background())

	go worker(cancelCtx)

	time.Sleep(3 * time.Second)

	cancel()

	time.Sleep(time.Second)

	// ----------------------------------------
	// Example 3
	// WithTimeout
	// ----------------------------------------

	fmt.Println("\n3. WithTimeout")

	timeoutCtx, timeoutCancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)

	defer timeoutCancel()

	downloadFile(timeoutCtx)

	// ----------------------------------------
	// Example 4
	// WithDeadline
	// ----------------------------------------

	fmt.Println("\n4. WithDeadline")

	deadline := time.Now().Add(3 * time.Second)

	deadlineCtx, deadlineCancel := context.WithDeadline(
		context.Background(),
		deadline,
	)

	defer deadlineCancel()

	select {

	case <-time.After(5 * time.Second):

		fmt.Println("Completed")

	case <-deadlineCtx.Done():

		fmt.Println("Deadline Reached")

	}

	// ----------------------------------------
	// Example 5
	// Context Value
	// ----------------------------------------

	fmt.Println("\n5. Context Value")

	type contextKey string

	const UserIDKey contextKey = "userID"

	valueCtx := context.WithValue(
		context.Background(),
		UserIDKey,
		101,
	)

	fmt.Println("User ID:", valueCtx.Value(UserIDKey))

}