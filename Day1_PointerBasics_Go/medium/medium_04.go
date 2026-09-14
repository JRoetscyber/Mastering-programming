// FILE: medium_04.go
// TOPIC: Pointer Basics - Day 1 (Go)
// DIFFICULTY: Medium
//
// TASK: Pointer to a specific element inside a slice
// --------------------------------------------------
// In Go you can take the address of a slice element: &slice[i]
//
// 1. Declare: scores := []int{10, 20, 30, 40, 50}
// 2. Declare a *int pointer `p` pointing to the THIRD element (index 2).
// 3. Print the value at `p` (should be 30).
// 4. Use `p` to change the third element to 99.
// 5. Print the ENTIRE slice using a range loop to confirm the change.
//
// EXPECTED OUTPUT:
//   Pointed-to value: 30
//   Slice after edit: [10 20 99 40 50]
//
// WARNING: If the slice is re-allocated (e.g. via append), the pointer
//          becomes invalid. This is a real Go gotcha to be aware of.

package main

import "fmt"

func main() {
	// TODO: Write your solution here.

}
