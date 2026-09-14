// FILE: hard_03.go
// TOPIC: Pointer Basics - Day 1 (Go)
// DIFFICULTY: Hard
//
// TASK: Pointer aliasing — two pointers to the same memory cell
// -------------------------------------------------------------
// When two pointers hold the same address, a write through either one
// is visible through the other. This is called aliasing.
//
// 1. Declare: x := 42
// 2. Declare: p := &x  and  q := &x   (both point to x)
// 3. Print both *p and *q — they should both be 42.
// 4. Write 100 through p (*p = 100).
//    Print *q — it should also be 100 (same memory!).
// 5. Write 200 through q (*q = 200).
//    Print x directly — it should be 200.
// 6. Print whether p == q (they should hold the same address).
//
// EXPECTED OUTPUT:
//   Initial : *p=42, *q=42
//   After *p=100: *q=100
//   After *q=200: x=200
//   p == q: true

package main

import "fmt"

func main() {
	// TODO: Write your solution here.

}
