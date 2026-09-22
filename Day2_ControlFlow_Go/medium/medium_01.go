// FILE: medium_01.go
// TOPIC: Control Flow Deep Dive - Day 2 (Go)
// DIFFICULTY: Medium
//
// TASK: Loop Control with continue & break over a slice
// -----------------------------------------------------
// You have a slice of sensor readings:
//   readings := []int{12, -1, 45, 0, 999, 28, -5, 60, 999, 10}
//
// Loop through them using `for _, r := range readings`:
// 1. If r < 0:
//    - Print: "Skipping corrupted sensor reading: <r>"
//    - Use `continue`
// 2. If r == 999:
//    - Print: "CRITICAL FAILURE (999) detected! Aborting loop."
//    - Use `break`
// 3. Otherwise:
//    - Accumulate validTotal and validCount
//    - Print: "Valid reading: <r>"
// 4. After the loop, print total count, sum, and average.
//
// EXPECTED OUTPUT:
//   Valid reading: 12
//   Skipping corrupted sensor reading: -1
//   Valid reading: 45
//   Valid reading: 0
//   CRITICAL FAILURE (999) detected! Aborting loop.
//   Processed valid count: 3
//   Valid total: 57
//   Valid average: 19

package main

import "fmt"

func main() {
	// TODO: Write your solution here.

}
