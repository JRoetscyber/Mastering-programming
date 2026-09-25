// FILE: challenge_04_worker_pool/exercise/exercise.go
// TOPIC: Production Worker Pool & Graceful Closer Pattern
//
// PRODUCTION SCENARIO:
// If 500 tooling CAD jobs arrive overnight in DiePro, you cannot spawn
// 500 simultaneous goroutines without risking memory exhaustion or crashing
// external CAD rendering services.
//
// Instead, you configure a POOL of fixed workers (e.g., 3 or 4) that chew
// through a buffered jobs queue. When the queue is drained and closed,
// each worker cleanly terminates, and a closer goroutine shuts down the
// results stream so the consumer finishes cleanly.
//
// YOUR GOALS:
// 1. Complete `Worker`:
//    - Range over `jobs <-chan Job`.
//    - Process each job and push `Result` to `results chan<- Result`.
//    - Record the job count in `stats.Record(id)`.
//    - Guarantee `wg.Done()` runs when jobs channel is drained.
// 2. Complete `RunProductionWorkerPool`:
//    - Create buffered channels for `jobs` and `results`.
//    - Launch `numWorkers` worker goroutines with `wg.Add(1)`.
//    - Feed all jobs into `jobs` and `close(jobs)`.
//    - Spin up the closer goroutine (`wg.Wait()` followed by `close(results)`).
//    - Read all results from `results` using `for r := range results`.
//    - Return results and the distribution map.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a die tooling render request.
type Job struct {
	TicketID   string
	Complexity int // 1 to 5 (simulates rendering complexity)
}

// Result represents the completed tooling rendering output.
type Result struct {
	TicketID     string
	WorkerID     int
	RenderTimeMs int
}

// WorkerStats tracks how many jobs each worker processed.
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

// Worker processes jobs until the `jobs` channel is closed.
func Worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup, stats *WorkerStats) {
	// TODO: Task 1 - Setup worker:
	// 1. Ensure `wg.Done()` is called when this function exits.
	// 2. Use `for job := range jobs` to loop until channel is closed.
	// 3. For each job:
	//    - sleep for time.Duration(job.Complexity*10) * time.Millisecond
	//    - send Result{TicketID: job.TicketID, WorkerID: id, RenderTimeMs: job.Complexity * 10} to `results`
	//    - call stats.Record(id)
}

// RunProductionWorkerPool creates a worker pool of size `numWorkers`, feeds the jobs,
// and returns all results along with a map showing how many jobs each worker handled.
func RunProductionWorkerPool(numWorkers int, jobs []Job) ([]Result, map[int]int) {
	stats := NewWorkerStats()

	// TODO: Task 2 - Worker pool orchestration:
	// 1. Create buffered channels:
	//      jobsChan := make(chan Job, len(jobs))
	//      resultsChan := make(chan Result, len(jobs))
	// 2. Launch `numWorkers` workers, calling `wg.Add(1)` for each.
	// 3. Push all `jobs` into `jobsChan`, then call `close(jobsChan)`.
	// 4. Start the closer goroutine:
	//      go func() {
	//          wg.Wait()
	//          close(resultsChan)
	//      }()
	// 5. Collect results:
	//      var collected []Result
	//      for r := range resultsChan { collected = append(collected, r) }
	// 6. Return collected and stats.GetCounts()

	return nil, stats.GetCounts()
}

func main() {
	jobs := []Job{
		{TicketID: "RO26-0001", Complexity: 2},
		{TicketID: "RO26-0002", Complexity: 4},
		{TicketID: "FL26-0010", Complexity: 1},
		{TicketID: "RO26-0003", Complexity: 3},
		{TicketID: "FL26-0011", Complexity: 2},
		{TicketID: "ST26-0004", Complexity: 5},
		{TicketID: "RO26-0004", Complexity: 1},
		{TicketID: "FL26-0012", Complexity: 3},
		{TicketID: "RO26-0005", Complexity: 2},
	}

	const poolSize = 3
	fmt.Printf("=== Starting Worker Pool (Workers: %d, Jobs: %d) ===\n\n", poolSize, len(jobs))
	start := time.Now()

	results, distribution := RunProductionWorkerPool(poolSize, jobs)

	fmt.Println("--- Job Processing Log ---")
	for _, r := range results {
		fmt.Printf("Ticket %-10s processed by Worker #%d (%dms)\n", r.TicketID, r.WorkerID, r.RenderTimeMs)
	}

	fmt.Println("\n--- Worker Workload Distribution ---")
	for w := 1; w <= poolSize; w++ {
		fmt.Printf("Worker #%d completed: %d jobs\n", w, distribution[w])
	}

	fmt.Printf("\nTotal time elapsed: %v\n", time.Since(start))
}
