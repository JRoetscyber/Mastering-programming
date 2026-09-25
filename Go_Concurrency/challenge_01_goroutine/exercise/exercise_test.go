package main

import (
	"sync"
	"testing"
	"time"
)

// TestRunInspectionsConcurrently_Timing verifies that stations run in parallel,
// completing in significantly less time than sequential execution.
func TestRunInspectionsConcurrently_Timing(t *testing.T) {
	const checkDuration = 60 * time.Millisecond
	stations := []Station{
		{ID: 1, ToolName: "Station 1", CheckTime: checkDuration},
		{ID: 2, ToolName: "Station 2", CheckTime: checkDuration},
		{ID: 3, ToolName: "Station 3", CheckTime: checkDuration},
		{ID: 4, ToolName: "Station 4", CheckTime: checkDuration},
	}

	var mu sync.Mutex
	completed := make(map[int]bool)

	start := time.Now()
	RunInspectionsConcurrently(stations, func(s Station) {
		mu.Lock()
		completed[s.ID] = true
		mu.Unlock()
	})

	// Wait up to 180ms for all 4 to finish (sequential would be 240ms)
	deadline := time.After(180 * time.Millisecond)
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	done := false
	for !done {
		select {
		case <-deadline:
			t.Fatalf("Timed out: inspections did not complete concurrently within 180ms")
		case <-ticker.C:
			mu.Lock()
			count := len(completed)
			mu.Unlock()
			if count == len(stations) {
				done = true
			}
		}
	}

	elapsed := time.Since(start)
	if elapsed >= 240*time.Millisecond {
		t.Errorf("Expected concurrent execution (< 240ms), but took %v (likely ran sequentially)", elapsed)
	}
}

// TestRunInspectionsConcurrently_AllCompleted verifies that each unique station
// is received and processed, preventing closure capture bugs where the same station is repeated.
func TestRunInspectionsConcurrently_AllCompleted(t *testing.T) {
	stations := []Station{
		{ID: 101, ToolName: "Die A", CheckTime: 20 * time.Millisecond},
		{ID: 102, ToolName: "Die B", CheckTime: 20 * time.Millisecond},
		{ID: 103, ToolName: "Die C", CheckTime: 20 * time.Millisecond},
	}

	var mu sync.Mutex
	seen := make(map[int]int)

	RunInspectionsConcurrently(stations, func(s Station) {
		mu.Lock()
		seen[s.ID]++
		mu.Unlock()
	})

	time.Sleep(100 * time.Millisecond)

	mu.Lock()
	defer mu.Unlock()

	for _, s := range stations {
		if count, ok := seen[s.ID]; !ok || count != 1 {
			t.Errorf("Station ID %d was not processed exactly once (got %d times). Check for closure capture!", s.ID, count)
		}
	}
}
