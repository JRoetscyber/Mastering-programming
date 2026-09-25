package main

import (
	"fmt"
	"time"
)

// printJob simulates a machine doing work on a die job.
// In real life this could be fetching a ticket from the database.
func printJob(jobNumber int) {
	fmt.Printf("Machine started on job #%d\n", jobNumber)
	time.Sleep(500 * time.Millisecond) // pretend it takes 500ms to process
	fmt.Printf("Machine finished job #%d\n", jobNumber)
}

func main() {
	fmt.Println("--- Without goroutines (one at a time) ---")

	// This runs jobs one after the other. Job 2 waits for job 1 to finish.
	for i := 1; i <= 3; i++ {
		printJob(i)
	}

	fmt.Println("\n--- With goroutines (all at once) ---")

	// The keyword `go` in front of a function call says:
	// "Start this, but don't wait for it — keep moving."
	for i := 1; i <= 3; i++ {
		go printJob(i)
	}

	// PROBLEM: main() exits too fast and kills the goroutines before they print.
	// The Sleep below is a dirty fix just to see output. We fix this properly in Challenge 2.
	time.Sleep(2 * time.Second)
	fmt.Println("Done.")
}
