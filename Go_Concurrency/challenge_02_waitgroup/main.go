package main

import (
	"fmt"
	"sync"
	"time"
)

// processTicket simulates reading a job ticket off the shop floor queue.
func processTicket(ticketID string, wg *sync.WaitGroup) {
	// defer means "run this when this function exits, no matter what".
	// It tells the WaitGroup: "I'm done."
	defer wg.Done()

	fmt.Printf("[START] Processing ticket %s\n", ticketID)
	time.Sleep(600 * time.Millisecond) // simulated work
	fmt.Printf("[DONE]  Ticket %s complete\n", ticketID)
}

func main() {
	// A WaitGroup is like a counter.
	// Add(1) = "one more worker starting"
	// Done() = "one worker finished"
	// Wait() = "block here until the counter hits zero"
	var wg sync.WaitGroup

	tickets := []string{"RO26-0001", "RO26-0002", "RO26-0003", "FL26-0010"}

	for _, ticket := range tickets {
		wg.Add(1) // tell the WaitGroup: one more is starting
		go processTicket(ticket, &wg)
	}

	// Wait here until every goroutine has called wg.Done()
	wg.Wait()
	fmt.Println("\nAll tickets processed.")
}
