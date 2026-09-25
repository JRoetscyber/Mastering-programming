// FILE: challenge_02_waitgroup/exercise/exercise.go
// TOPIC: Deterministic Synchronization with sync.WaitGroup
//
// SHOP FLOOR SCENARIO:
// In DiePro, when an operator dispatches a batch of production job tickets
// (e.g., RO26-0001, FL26-0010), we cannot guess how long each ticket takes
// to render and generate tooling CAD specs. Some take 10ms, others take 200ms.
// Using `time.Sleep` is unacceptable. We need `sync.WaitGroup` to wait until
// ALL tickets are finished, and then immediately continue.
//
// YOUR GOALS:
// 1. Complete `ProcessTicket` using `defer wg.Done()` at the top.
// 2. Complete `BatchProcessTickets`:
//    - Initialize a `sync.WaitGroup`.
//    - For each ticket, call `wg.Add(1)`.
//    - Launch a goroutine passing `&wg` (by pointer!).
//    - Call `wg.Wait()` to block until all workers complete.
// 3. Return the list of completed ticket IDs.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Ticket represents a production die job ticket.
type Ticket struct {
	ID       string
	Customer string
	DieType  string // "Rotary", "Flatbed", "Flexo"
}

// SafeTicketList provides a thread-safe slice to collect processed ticket IDs.
type SafeTicketList struct {
	mu      sync.Mutex
	tickets []string
}

func (s *SafeTicketList) Add(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tickets = append(s.tickets, id)
}

func (s *SafeTicketList) GetAll() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	result := make([]string, len(s.tickets))
	copy(result, s.tickets)
	return result
}

// ProcessTicket simulates background processing for a ticket.
func ProcessTicket(t Ticket, wg *sync.WaitGroup, list *SafeTicketList) {
	// TODO: Task 1 - Ensure the WaitGroup counter decrements when this function exits.
	// HINT: Use `defer` so it always executes!

	// Simulated CAD generation time
	time.Sleep(50 * time.Millisecond)

	list.Add(t.ID)
}

// BatchProcessTickets coordinates processing a batch of tickets concurrently.
// It must NOT use time.Sleep to wait for completion!
func BatchProcessTickets(tickets []Ticket) []string {
	var list SafeTicketList

	// TODO: Task 2 - Initialize the WaitGroup, loop over each ticket:
	// 1. Call wg.Add(1)
	// 2. Launch `ProcessTicket(t, &wg, &list)` in a goroutine
	// 3. Call wg.Wait() after the loop so main blocks until all are finished

	return list.GetAll()
}

func main() {
	tickets := []Ticket{
		{ID: "RO26-0001", Customer: "Apex Packaging", DieType: "Rotary"},
		{ID: "RO26-0002", Customer: "Golden Label Co", DieType: "Rotary"},
		{ID: "FL26-0010", Customer: "CartonCraft", DieType: "Flatbed"},
		{ID: "ST26-0044", Customer: "Speedy Print", DieType: "Flexo"},
		{ID: "RO26-0005", Customer: "Matrix Die", DieType: "Rotary"},
	}

	fmt.Println("=== Batch Processing Production Tickets ===")
	start := time.Now()

	processed := BatchProcessTickets(tickets)

	fmt.Printf("Successfully processed %d tickets in %v:\n", len(processed), time.Since(start))
	for _, id := range processed {
		fmt.Printf(" - %s\n", id)
	}
}
