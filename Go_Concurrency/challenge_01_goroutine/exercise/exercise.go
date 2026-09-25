// FILE: challenge_01_goroutine/exercise/exercise.go
// TOPIC: The `go` Keyword & Goroutine Basics
//
// SHOP FLOOR SCENARIO:
// You have multiple inspection stations checking die-cut rotary dies.
// Running them one-by-one in sequence slows down the factory shift.
// Your task is to launch each station inspection concurrently in its own goroutine.
//
// YOUR GOALS:
// 1. Complete `InspectStation` to simulate work and record completion.
// 2. Complete `RunInspectionsConcurrently` to spin up a goroutine for each station.
//    BE CAREFUL: Avoid the classic loop closure capture bug! Pass the station
//    variable as an argument into the goroutine or closure.
// 3. In `main()`, run the concurrent inspections and observe the total elapsed time.

package main

import (
	"fmt"
	"time"
)

// Station represents an inspection station on the shop floor.
type Station struct {
	ID        int
	ToolName  string
	CheckTime time.Duration
}

// InspectStation simulates the inspection work for a single station.
// It should:
//   - Sleep for s.CheckTime
//   - If onDone != nil, invoke onDone(s)
func InspectStation(s Station, onDone func(Station)) {
	// TODO: Task 1 - Simulate inspection work:
	// 1. Sleep for s.CheckTime
	// 2. If onDone != nil, invoke onDone(s)
	time.Sleep(s.CheckTime)
	if onDone != nil {
		onDone(s)
	}
}

// RunInspectionsConcurrently launches an inspection for each station
// concurrently using goroutines.
//
// HINT: When launching goroutines in a loop, make sure each goroutine
// receives its own copy of the station!
func RunInspectionsConcurrently(stations []Station, onDone func(Station)) {
	for s := range stations {
		stationPtr := &stations[s]

		go func(s *Station) {
			InspectStation(*s, onDone)
		}(stationPtr)
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

	// Launch inspections concurrently
	RunInspectionsConcurrently(stations, onComplete)

	// Since we haven't reached Challenge 2 (WaitGroup), we use a temporary sleep
	// so main() does not exit before our goroutines finish printing.
	time.Sleep(500 * time.Millisecond)

	elapsed := time.Since(start)
	fmt.Printf("\nFinished in: %v\n", elapsed)
	fmt.Printf("Notice: Sequential time would be ~800ms. Concurrent time is ~200ms!\n")
}
