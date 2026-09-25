package main

import (
	"fmt"
	"sync"
	"time"
)

// Job is the work we want done — price a single ticket.
type Job struct {
	TicketID  string
	BasePrice float64
}

// Result is what comes back when the job is done.
type Result struct {
	TicketID   string
	FinalPrice float64
}

// worker is ONE machine on the floor.
// It reads jobs from the `jobs` channel and writes results to the `results` channel.
// When the jobs channel is closed (no more work), it stops and calls wg.Done().
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs { // loop until the jobs channel is closed
		fmt.Printf("Worker %d picked up ticket %s\n", id, job.TicketID)
		time.Sleep(300 * time.Millisecond) // simulate work
		results <- Result{
			TicketID:   job.TicketID,
			FinalPrice: job.BasePrice * 1.15,
		}
	}
	fmt.Printf("Worker %d is done — no more jobs.\n", id)
}

func main() {
	const numWorkers = 3 // three machines running in parallel

	jobs := make(chan Job, 10)    // a buffered pipe that can hold up to 10 jobs
	results := make(chan Result, 10)

	// Start the workers BEFORE you send any jobs
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Feed 8 tickets into the jobs channel
	tickets := []Job{
		{"RO26-0001", 1200},
		{"RO26-0002", 2100},
		{"FL26-0010", 850},
		{"RO26-0003", 3300},
		{"FL26-0011", 920},
		{"RO26-0004", 1750},
		{"FL26-0012", 680},
		{"RO26-0005", 2600},
	}

	for _, t := range tickets {
		jobs <- t
	}
	close(jobs) // signal: no more work coming — workers will finish and exit

	// Wait for all workers to finish, then close the results channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print every result as it arrives
	fmt.Println("\n--- Results ---")
	for r := range results {
		fmt.Printf("Ticket %-10s  Final price: R%.2f\n", r.TicketID, r.FinalPrice)
	}

	fmt.Println("\nAll done.")
}
