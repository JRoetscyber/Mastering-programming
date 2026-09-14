// FILE: hard_02.go
// TOPIC: Pointer Basics - Day 1 (Go)
// DIFFICULTY: Hard
//
// TASK: Pass a pointer to a function to modify the caller's variable
// ------------------------------------------------------------------
// Go is pass-by-value. To let a function modify the caller's variable,
// pass a POINTER to it.
//
// 1. Write:  func doubleIt(ptr *int)  — multiplies *ptr by 2.
// 2. Write:  func addFive(ptr *int)   — adds 5 to *ptr.
// 3. In main():
//      a) Declare n := 10
//      b) Call doubleIt(&n). Print n.
//      c) Call addFive(&n).  Print n.
//      d) Call doubleIt(&n). Print n.
//
// EXPECTED OUTPUT:
//   After doubleIt: 20
//   After addFive : 25
//   After doubleIt: 50

package main

import "fmt"

// TODO: Write doubleIt and addFive here.

func main() {
	// TODO: Write your solution here.

}
