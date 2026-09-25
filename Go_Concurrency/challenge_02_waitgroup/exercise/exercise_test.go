package main

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestBatchProcessTickets_AllProcessed(t *testing.T) {
	tickets := []Ticket{
		{ID: "RO26-0001", Customer: "Cust 1", DieType: "Rotary"},
		{ID: "RO26-0002", Customer: "Cust 2", DieType: "Rotary"},
		{ID: "FL26-0010", Customer: "Cust 3", DieType: "Flatbed"},
		{ID: "FL26-0011", Customer: "Cust 4", DieType: "Flatbed"},
		{ID: "ST26-0005", Customer: "Cust 5", DieType: "Flexo"},
	}

	start := time.Now()
	results := BatchProcessTickets(tickets)
	elapsed := time.Since(start)

	if len(results) != len(tickets) {
		t.Fatalf("Expected %d processed tickets, got %d. Did wg.Wait() exit early or were goroutines skipped?", len(tickets), len(results))
	}

	// Should run concurrently (~50ms), well below sequential (~250ms)
	if elapsed >= 200*time.Millisecond {
		t.Errorf("Execution took %v. Expected concurrent execution under 200ms.", elapsed)
	}

	// Verify all ticket IDs exist in results
	expectedIDs := []string{"FL26-0010", "FL26-0011", "RO26-0001", "RO26-0002", "ST26-0005"}
	sort.Strings(results)
	sort.Strings(expectedIDs)

	for i := range expectedIDs {
		if results[i] != expectedIDs[i] {
			t.Errorf("Mismatch at index %d: expected %s, got %s", i, expectedIDs[i], results[i])
		}
	}
}

func TestBatchProcessTickets_Stress100(t *testing.T) {
	tickets := make([]Ticket, 100)
	for i := 0; i < 100; i++ {
		tickets[i] = Ticket{
			ID:       fmt.Sprintf("JOB-%04d", i),
			Customer: "Bulk Stress Test",
			DieType:  "Rotary",
		}
	}

	results := BatchProcessTickets(tickets)
	if len(results) != 100 {
		t.Fatalf("Stress test failed: expected 100 tickets, got %d", len(results))
	}
}
