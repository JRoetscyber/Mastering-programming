package main

import (
	"fmt"
	"testing"
)

func TestWorkerPool_AllJobsCompleted(t *testing.T) {
	const jobCount = 20
	const numWorkers = 3

	jobs := make([]Job, jobCount)
	for i := 0; i < jobCount; i++ {
		jobs[i] = Job{
			TicketID:   fmt.Sprintf("JOB-%03d", i+1),
			Complexity: (i % 3) + 1,
		}
	}

	results, distribution := RunProductionWorkerPool(numWorkers, jobs)

	if len(results) != jobCount {
		t.Fatalf("Expected %d completed results, got %d", jobCount, len(results))
	}

	// Verify all jobs are present in the results
	seen := make(map[string]bool)
	for _, r := range results {
		seen[r.TicketID] = true
	}

	for _, j := range jobs {
		if !seen[j.TicketID] {
			t.Errorf("Missing result for %s", j.TicketID)
		}
	}

	// Verify distribution: each of the 3 workers should have processed at least 1 job
	totalJobsDistributed := 0
	for w := 1; w <= numWorkers; w++ {
		count := distribution[w]
		if count == 0 {
			t.Errorf("Worker %d received 0 jobs. Work was not distributed across the pool!", w)
		}
		totalJobsDistributed += count
	}

	if totalJobsDistributed != jobCount {
		t.Errorf("Expected total recorded jobs in stats to be %d, got %d", jobCount, totalJobsDistributed)
	}
}

func TestWorkerPool_SingleWorker(t *testing.T) {
	jobs := []Job{
		{TicketID: "SOLO-1", Complexity: 1},
		{TicketID: "SOLO-2", Complexity: 1},
	}

	results, distribution := RunProductionWorkerPool(1, jobs)
	if len(results) != 2 {
		t.Fatalf("Expected 2 results, got %d", len(results))
	}
	if distribution[1] != 2 {
		t.Errorf("Expected Worker 1 to handle 2 jobs, got %d", distribution[1])
	}
}
