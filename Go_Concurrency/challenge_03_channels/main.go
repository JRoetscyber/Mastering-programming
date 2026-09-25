package main

import (
	"fmt"
	"time"
)

// PricingResult holds what a worker sends back.
type PricingResult struct {
	TicketID string
	Price    float64
}

// calculatePrice simulates the quoting engine pricing a ticket.
// Instead of returning a value normally, it SENDS the result down a channel.
func calculatePrice(ticketID string, basePrice float64, results chan<- PricingResult) {
	time.Sleep(400 * time.Millisecond) // simulate database + math work

	// Add a 15% setup fee
	finalPrice := basePrice * 1.15

	// Send the result into the channel.
	// This is like putting a note in a pipe for the main function to read.
	results <- PricingResult{TicketID: ticketID, Price: finalPrice}
}

func main() {
	tickets := map[string]float64{
		"RO26-0001": 1200.00,
		"FL26-0010": 850.00,
		"RO26-0002": 2100.00,
	}

	// Make a channel that can carry PricingResult values.
	// Think of it as a pipe between goroutines.
	results := make(chan PricingResult)

	// Launch a goroutine for each ticket
	for id, price := range tickets {
		go calculatePrice(id, price, results)
	}

	// Read one result from the channel for each ticket we sent.
	// <-results means "wait here until something arrives in the pipe".
	for range tickets {
		r := <-results
		fmt.Printf("Ticket %-10s  Final price: R%.2f\n", r.TicketID, r.Price)
	}

	fmt.Println("All prices calculated.")
}
