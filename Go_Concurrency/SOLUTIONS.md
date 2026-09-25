# Go Concurrency — Reference Solutions & Mentor Guide

> **Attempt each challenge on your own first!**
> Real concurrency intuition only builds when you write the code, hit a deadlock or race condition, diagnose it, and see your tests turn green.

---

## Challenge 1 — Goroutines & Concurrency Fundamentals

📁 Location: `challenge_01_goroutine/exercise/`

### Complete Solution

```go
package main

import (
	"fmt"
	"time"
)

type Station struct {
	ID        int
	ToolName  string
	CheckTime time.Duration
}

func InspectStation(s Station, onDone func(Station)) {
	// Task 1: Simulate inspection work
	time.Sleep(s.CheckTime)
	if onDone != nil {
		onDone(s)
	}
}

func RunInspectionsConcurrently(stations []Station, onDone func(Station)) {
	// Task 2: Launch each station in its own goroutine
	for _, s := range stations {
		// Passing `s` as a parameter to the goroutine function (or passing `s`
		// into InspectStation) ensures that each goroutine captures a unique copy
		// of the loop variable.
		go InspectStation(s, onDone)
	}
}

func main() {
	stations := []Station{
		{ID: 1, ToolName: "Rotary Cylinder 01", CheckTime: 200 * time.Millisecond},
		{ID: 2, ToolName: "Flatbed Plate 04", CheckTime: 200 * time.Millisecond},
		{ID: 3, ToolName: "Magnetic Cylinder 02", CheckTime: 200 * time.Millisecond},
		{ID: 4, ToolName: "Flexo Stereo Die", CheckTime: 200 * time.Millisecond},
	}

	fmt.Println("=== Starting Concurrent Shop Floor Inspection ===")
	start := time.Now()

	onComplete := func(s Station) {
		fmt.Printf("[DONE] Station #%d (%s) finished inspection.\n", s.ID, s.ToolName)
	}

	RunInspectionsConcurrently(stations, onComplete)

	// Temporary sleep until Challenge 2 teaches WaitGroup
	time.Sleep(500 * time.Millisecond)

	elapsed := time.Since(start)
	fmt.Printf("\nFinished in: %v\n", elapsed)
	fmt.Printf("Notice: Sequential time would be ~800ms. Concurrent time is ~200ms!\n")
}
```

### Senior Dev Takeaways
1. **The Closure Capture Bug**: In older Go versions (pre-1.22), referencing the loop variable `s` directly inside an anonymous `go func() { ... }` closure would cause all goroutines to point to the exact same variable address, meaning every goroutine might end up processing the last item in the slice! Passing `s` as a function parameter (`go InspectStation(s, ...)`) or `go func(st Station) { ... }(s)` avoids this completely.
2. **Main Exiting Early**: Remember that when `main()` exits, the entire Go runtime process tears down, killing any active background goroutines immediately.

---

## Challenge 2 — Deterministic Synchronization with `sync.WaitGroup`

📁 Location: `challenge_02_waitgroup/exercise/`

### Complete Solution

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Ticket struct {
	ID       string
	Customer string
	DieType  string
}

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

func ProcessTicket(t Ticket, wg *sync.WaitGroup, list *SafeTicketList) {
	// Task 1: Always decrement counter when exiting
	defer wg.Done()

	// Simulated CAD generation
	time.Sleep(50 * time.Millisecond)

	list.Add(t.ID)
}

func BatchProcessTickets(tickets []Ticket) []string {
	var list SafeTicketList
	var wg sync.WaitGroup

	// Task 2: Coordinate all tickets with the WaitGroup
	for _, t := range tickets {
		wg.Add(1)
		go ProcessTicket(t, &wg, &list) // MUST pass &wg by pointer!
	}

	// Block here until every worker calls wg.Done()
	wg.Wait()

	return list.GetAll()
}
```

### Senior Dev Takeaways
1. **Always Pass `*sync.WaitGroup` by Pointer**: A `sync.WaitGroup` contains internal counter state and mutexes. If you pass `wg` by value (without `&`), Go makes a bitwise copy. Each goroutine will decrement its own copy, leaving the original counter unchanged and causing `wg.Wait()` to hang forever (deadlock).
2. **`defer wg.Done()` at the Top**: Never put `wg.Done()` at the very bottom of a function without `defer`. If an early return, error check, or panic occurs, `wg.Done()` will be skipped and `wg.Wait()` will hang.

---

## Challenge 3 — Safe Streaming with Directional Channels

📁 Location: `challenge_03_channels/exercise/`

### Complete Solution

```go
package main

import (
	"fmt"
	"time"
)

type JobTicket struct {
	TicketID  string
	BasePrice float64
}

type PricingResult struct {
	TicketID   string
	BasePrice  float64
	SetupFee   float64
	Surcharge  float64
	FinalPrice float64
}

func CalculatePricing(ticket JobTicket, results chan<- PricingResult) {
	time.Sleep(30 * time.Millisecond)

	// Task 1: Calculate fees and surcharge
	setupFee := ticket.BasePrice * 0.15
	surcharge := 0.0
	if ticket.BasePrice > 2000.0 {
		surcharge = ticket.BasePrice * 0.10
	}
	finalPrice := ticket.BasePrice + setupFee + surcharge

	// Send result down the pipe
	results <- PricingResult{
		TicketID:   ticket.TicketID,
		BasePrice:  ticket.BasePrice,
		SetupFee:   setupFee,
		Surcharge:  surcharge,
		FinalPrice: finalPrice,
	}
}

func BatchCalculatePricing(tickets []JobTicket) []PricingResult {
	// Task 2: Channel pipeline
	results := make(chan PricingResult, len(tickets))

	// Launch a goroutine for each ticket
	for _, t := range tickets {
		go CalculatePricing(t, results)
	}

	// Collect exactly len(tickets) results from the channel
	var collected []PricingResult
	for i := 0; i < len(tickets); i++ {
		r := <-results
		collected = append(collected, r)
	}

	return collected
}
```

### Senior Dev Takeaways
1. **Directional Channels**: Notice `chan<- PricingResult` in `CalculatePricing`. By constraining the parameter to write-only, the compiler prevents accidental reads or close calls inside the worker.
2. **Channel Deadlocks**: If you try to read more items than are ever sent to an unclosed channel (`<-results`), the Go runtime detects that all goroutines are asleep and throws `fatal error: all goroutines are asleep - deadlock!`.

---

## Challenge 4 — Production Worker Pool & Graceful Closer Pattern

📁 Location: `challenge_04_worker_pool/exercise/`

### Complete Solution

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	TicketID   string
	Complexity int
}

type Result struct {
	TicketID     string
	WorkerID     int
	RenderTimeMs int
}

type WorkerStats struct {
	mu     sync.Mutex
	counts map[int]int
}

func NewWorkerStats() *WorkerStats {
	return &WorkerStats{counts: make(map[int]int)}
}

func (s *WorkerStats) Record(workerID int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counts[workerID]++
}

func (s *WorkerStats) GetCounts() map[int]int {
	s.mu.Lock()
	defer s.mu.Unlock()
	copied := make(map[int]int)
	for k, v := range s.counts {
		copied[k] = v
	}
	return copied
}

func Worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup, stats *WorkerStats) {
	defer wg.Done()

	// Task 1: Drain jobs until closed
	for job := range jobs {
		renderTime := job.Complexity * 10
		time.Sleep(time.Duration(renderTime) * time.Millisecond)

		results <- Result{
			TicketID:     job.TicketID,
			WorkerID:     id,
			RenderTimeMs: renderTime,
		}

		stats.Record(id)
	}
}

func RunProductionWorkerPool(numWorkers int, jobs []Job) ([]Result, map[int]int) {
	stats := NewWorkerStats()

	// Task 2: Worker pool orchestration
	jobsChan := make(chan Job, len(jobs))
	resultsChan := make(chan Result, len(jobs))

	var wg sync.WaitGroup

	// 1. Start workers first
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(i, jobsChan, resultsChan, &wg, stats)
	}

	// 2. Feed all jobs into the jobs channel
	for _, j := range jobs {
		jobsChan <- j
	}
	close(jobsChan) // Signal workers: no more work coming!

	// 3. Closer goroutine: wait for workers to drain jobs, then close results
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// 4. Collect results until resultsChan is closed
	var collected []Result
	for r := range resultsChan {
		collected = append(collected, r)
	}

	return collected, stats.GetCounts()
}
```

### Senior Dev Takeaways
1. **Why Worker Pools?**: Spawning an unbounded number of goroutines (e.g. 10,000) can exhaust memory and overload databases. A worker pool bounds active concurrency to a safe, constant number while processing an arbitrary backlog of work.
2. **The Closer Goroutine**:
   ```go
   go func() {
       wg.Wait()
       close(resultsChan)
   }()
   ```
   This background goroutine is one of the most elegant idioms in Go. It allows the consumer loop `for r := range resultsChan` to read items as they finish in real time, and cleanly terminate as soon as the last worker calls `wg.Done()`.
