// FILE: challenge_03_channels/exercise/exercise.go
// TOPIC: Safe Data Streaming with Channels & Directional Constraints
//
// SHOP FLOOR / QUOTING SCENARIO:
// In DiePro's estimating engine, each die tooling quotation requires
// calculating:
// 1. Setup Fee: 15% of BasePrice.
// 2. Material Surcharge: 10% of BasePrice ONLY if BasePrice > R2000.00 (otherwise 0.00).
// 3. FinalPrice = BasePrice + SetupFee + Surcharge.
//
// Because quotes are calculated concurrently across customer orders,
// we cannot write to a shared variable. Instead, workers SEND their results
// through a channel pipe to the collector.
//
// YOUR GOALS:
// 1. Complete `CalculatePricing`:
//    - Perform the calculations.
//    - Push the `PricingResult` into the send-only channel `results chan<- PricingResult`.
// 2. Complete `BatchCalculatePricing`:
//    - Create the channel using `make(chan PricingResult)`.
//    - Spawn a goroutine for each ticket.
//    - Receive all results from the channel and return them as a slice.

package main

import (
	"fmt"
	"time"
)

// JobTicket represents an incoming quote request.
type JobTicket struct {
	TicketID  string
	BasePrice float64
}

// PricingResult represents the calculated pricing breakdown.
type PricingResult struct {
	TicketID   string
	BasePrice  float64
	SetupFee   float64
	Surcharge  float64
	FinalPrice float64
}

// CalculatePricing calculates fees and pushes the result into the channel.
// Notice the channel type `chan<- PricingResult`: this function can ONLY send to it!
func CalculatePricing(ticket JobTicket, results chan<- PricingResult) {
	// Simulate calculation latency
	time.Sleep(30 * time.Millisecond)

	// TODO: Task 1 - Compute:
	// 1. setupFee = ticket.BasePrice * 0.15
	// 2. surcharge = 0.0
	//    if ticket.BasePrice > 2000.0 { surcharge = ticket.BasePrice * 0.10 }
	// 3. finalPrice = ticket.BasePrice + setupFee + surcharge
	// 4. Send `PricingResult{...}` into `results` channel pipe using `<-`
}

// BatchCalculatePricing runs pricing for all tickets concurrently and collects the results.
func BatchCalculatePricing(tickets []JobTicket) []PricingResult {
	// TODO: Task 2 - Channel pipeline:
	// 1. Create a channel for `PricingResult`
	// 2. Launch a goroutine calling `CalculatePricing` for each ticket
	// 3. Loop `len(tickets)` times to receive `<-results` and append to a slice
	// 4. Return the collected slice

	return nil
}

func main() {
	tickets := []JobTicket{
		{TicketID: "RO26-0001", BasePrice: 1200.00},
		{TicketID: "FL26-0010", BasePrice: 850.00},
		{TicketID: "RO26-0002", BasePrice: 2400.00}, // > 2000 (qualifies for surcharge)
		{TicketID: "ST26-0005", BasePrice: 3100.00}, // > 2000 (qualifies for surcharge)
		{TicketID: "FL26-0012", BasePrice: 1500.00},
	}

	fmt.Println("=== Estimating Engine: Concurrent Pricing Pipeline ===")
	start := time.Now()

	results := BatchCalculatePricing(tickets)

	fmt.Printf("%-12s | %-10s | %-10s | %-10s | %-10s\n", "Ticket", "Base", "Setup(15%)", "Surch(10%)", "Final Price")
	fmt.Println("----------------------------------------------------------------")
	for _, r := range results {
		fmt.Printf("%-12s | R%-9.2f | R%-9.2f | R%-9.2f | R%-9.2f\n",
			r.TicketID, r.BasePrice, r.SetupFee, r.Surcharge, r.FinalPrice)
	}
	fmt.Printf("\nCalculated %d quotes in %v\n", len(results), time.Since(start))
}
