// FILE: hard_03.go
// TOPIC: Control Flow Deep Dive - Day 2 (Go)
// DIFFICULTY: Hard
//
// TASK: Bubble Sort Invariants & Multi-variable Loop in Go
// --------------------------------------------------------
// Go supports tuple assignments (`a, b = b, a`) which makes swapping
// values inside a loop clean without temp variables!
//
// 1. Declare a slice:
//      arr := []int{4, 1, 8, 3, 9, 2, 7}
//
// 2. Implement Bubble Sort:
//      - Outer loop: i from 0 to len(arr) - 2
//      - Inner loop: j from 0 to len(arr) - i - 2
//      - Swap elements using Go syntax: arr[j], arr[j+1] = arr[j+1], arr[j]
//      - Early termination optimization: if no swaps occurred, break!
//
// 3. Print the number of passes taken and the sorted slice.
//
// EXPECTED OUTPUT:
//   Passes taken: 5
//   Sorted slice: [1 2 3 4 7 8 9]

package main

import "fmt"

func main() {
	// TODO: Write your solution here.

}
