// FILE: hard_04.go
// TOPIC: Pointer Basics - Day 1 (Go)
// DIFFICULTY: Hard
//
// TASK: Go allows returning a pointer to a local variable (heap escape)
// ---------------------------------------------------------------------
// In C++, returning the address of a local variable is undefined behaviour
// because the local is destroyed when the function returns.
// In Go, the COMPILER detects this and automatically allocates the variable
// on the HEAP, keeping it alive as long as the pointer exists.
// This is called "escape to heap".
//
// 1. Write a function:
//      func newInt(val int) *int
//    Inside, declare  n := val  and return &n.
//    (In C++ this would be UB. In Go it is perfectly safe.)
//
// 2. In main():
//      a) Call newInt(10) and store the returned pointer in `p`.
//      b) Print *p.
//      c) Modify the value through p (*p = 99).
//      d) Print *p again.
//      e) Call newInt(55) and store in a second pointer `q`.
//      f) Print *p and *q — they are INDEPENDENT heap allocations.
//
// EXPECTED OUTPUT:
//   *p after alloc  : 10
//   *p after modify : 99
//   *p=99, *q=55  (independent)

package main

import "fmt"

// TODO: Write newInt here.

func main() {
	// TODO: Write your solution here.

}
