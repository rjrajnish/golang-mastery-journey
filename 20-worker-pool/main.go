 package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID int
}

func worker(id int, jobs <-chan Job, results chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()

	for job := range jobs {

		fmt.Printf("Worker %d processing Job %d\n", id, job.ID)

		time.Sleep(time.Second)

		results <- fmt.Sprintf("Job %d Completed by Worker %d", job.ID, id)

	}

}

func main() {

	fmt.Println("========================================")
	fmt.Println("        Go Worker Pool Example")
	fmt.Println("========================================")

	const totalWorkers = 3

	const totalJobs = 10

	jobs := make(chan Job)

	results := make(chan string)

	var wg sync.WaitGroup

	// ----------------------------------------
	// Start Workers
	// ----------------------------------------

	for i := 1; i <= totalWorkers; i++ {

		wg.Add(1)

		go worker(i, jobs, results, &wg)

	}

	// ----------------------------------------
	// Send Jobs
	// ----------------------------------------

	go func() {

		for i := 1; i <= totalJobs; i++ {

			jobs <- Job{
				ID: i,
			}

		}

		close(jobs)

	}()

	// ----------------------------------------
	// Close Result Channel
	// ----------------------------------------

	go func() {

		wg.Wait()

		close(results)

	}()

	// ----------------------------------------
	// Receive Results
	// ----------------------------------------

	fmt.Println("\nProcessing Jobs...\n")

	for result := range results {

		fmt.Println(result)

	}

	fmt.Println("\nAll Jobs Completed")

}