// FILE: easy_01.go
// TOPIC: Control Flow Deep Dive - Day 2 (Go)
// DIFFICULTY: Easy
//
// TASK: Short-Circuit Evaluation & Guard Clause with if/else
// -----------------------------------------------------------
// In Go, logical AND (&&) stops immediately if the left condition is false.
// Logical OR (||) stops immediately if the left condition is true.
//
// This prevents nil pointer panics!
// For example:
//   if ptr != nil && *ptr > 50  <-- *ptr is NEVER dereferenced if ptr is nil!
//
// 1. Declare a *int pointer `var ptr *int` (zero value is nil).
// 2. Write an if-statement checking:
//      if ptr != nil && *ptr > 50:
//          print "Valid and greater than 50"
//      else:
//          print "Safety triggered: pointer is nil or value <= 50"
// 3. Verify that running this does NOT panic.
// 4. Next, assign `val := 85` and `ptr = &val`.
// 5. Run the same check again.
//
// EXPECTED OUTPUT:
//   Safety triggered: pointer is nil or value <= 50
//   Valid and greater than 50

package main

import "fmt"

func main() {
	// TODO: Write your solution here.

}
