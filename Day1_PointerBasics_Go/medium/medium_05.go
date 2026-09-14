// FILE: medium_05.go
// TOPIC: Pointer Basics - Day 1 (Go)
// DIFFICULTY: Medium
//
// TASK: Read and write through a pointer-to-pointer chain (**int)
// ---------------------------------------------------------------
// 1. Declare:  value := 5
// 2. Declare:  p1 := &value      (*int)
// 3. Declare:  p2 := &p1         (**int)
// 4. Using ONLY `p2` (no direct use of `value` or `p1` after this point):
//      a) Print the current value (double-dereference **p2).
//      b) Multiply `value` by 3 through the pointer chain (**p2 *= 3).
//      c) Print the new value again through p2.
// 5. Finally, print `value` directly to confirm it changed.
//
// EXPECTED OUTPUT:
//   Via p2 before: 5
//   Via p2 after : 15
//   value direct : 15

package main

import "fmt"

func main() {
	// TODO: Write your solution here.

}
