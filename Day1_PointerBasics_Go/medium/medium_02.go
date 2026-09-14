// FILE: medium_02.go
// TOPIC: Pointer Basics - Day 1 (Go)
// DIFFICULTY: Medium
//
// TASK: Pointer to a pointer (double indirection — **int)
// -------------------------------------------------------
// 1. Declare:  x  := 42           (an int)
// 2. Declare:  p  := &x           (*int)
// 3. Declare:  pp := &p           (**int  — a pointer to the pointer)
// 4. Print x in THREE ways:
//      a) directly: x
//      b) single dereference: *p
//      c) double dereference: **pp
// 5. Print the ADDRESS of x in TWO ways:
//      a) &x
//      b) *pp  (dereference pp once — you get the value stored in p, which is &x)
//
// EXPECTED OUTPUT (addresses will differ):
//   x = 42  |  *p = 42  |  **pp = 42
//   &x = 0x...  |  *pp = 0x...  (both should be identical)

package main

import "fmt"

func main() {
	// TODO: Write your solution here.

}
