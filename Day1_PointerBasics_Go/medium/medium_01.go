// FILE: medium_01.go
// TOPIC: Pointer Basics - Day 1 (Go)
// DIFFICULTY: Medium
//
// TASK: Modify a variable's value THROUGH its pointer
// ---------------------------------------------------
// 1. Declare an int `lives = 3`.
// 2. Print the original value of `lives`.
// 3. Declare a *int pointer `p` pointing to `lives`.
// 4. Use the pointer to subtract 1 from `lives` (write through *p).
// 5. Print the new value of `lives` directly (NOT via the pointer).
// 6. Use the pointer again to set `lives` to 0.
// 7. Print `lives` one final time.
//
// EXPECTED OUTPUT:
//   lives = 3
//   lives = 2
//   lives = 0
//
// NOTE: In Go, `*p -= 1` is the idiomatic way to decrement through a pointer.
//       `(*p)--` also works. `*p--` does NOT (Go's -- is a statement, not
//       an expression, and the parse would be ambiguous).

package main

import "fmt"

func main() {
	// TODO: Write your solution here.

}
