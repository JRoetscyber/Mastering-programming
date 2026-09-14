// FILE: hard_01.go
// TOPIC: Pointer Basics - Day 1 (Go)
// DIFFICULTY: Hard
//
// TASK: nil pointer guard — detect and handle a nil pointer safely
// ----------------------------------------------------------------
// In Go, a pointer that has not been assigned is nil (the zero value for
// all pointer types). Dereferencing nil causes a runtime panic.
//
// 1. Declare a *int pointer `p` without assigning it (it will be nil).
// 2. Write a helper function:
//      func safeRead(p *int) string
//    that returns "nil pointer — cannot read" if p is nil,
//    or returns the stringified value ("value = N") if p is not nil.
// 3. In main():
//      a) Call safeRead(p) and print the result (p is nil here).
//      b) Assign p to point to a local int x = 77.
//      c) Call safeRead(p) and print the result again.
//
// EXPECTED OUTPUT:
//   nil pointer - cannot read
//   value = 77
//
// BONUS QUESTION (no code needed):
//   What is the difference between a nil pointer and an uninitialised
//   pointer in C++? Why is Go's nil safer?

package main

import "fmt"

// TODO: Write safeRead here.

func main() {
	// TODO: Write your solution here.

}
