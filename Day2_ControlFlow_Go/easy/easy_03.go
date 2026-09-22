// FILE: easy_03.go
// TOPIC: Control Flow Deep Dive - Day 2 (Go)
// DIFFICULTY: Easy
//
// TASK: While and Do-While loops in Go (using the `for` keyword)
// --------------------------------------------------------------
// Go has NO `while` or `do while` keyword! Everything is a `for` loop:
//   - C++ while (cond)   -> Go: for cond {}
//   - C++ do-while       -> Go: for { ... if !cond { break } }
//
// 1. Declare: count := 10
// 2. Write a while-style loop in Go (`for count < 5`) that prints:
//      "While running: " + count
// 3. Write a do-while style loop in Go that is guaranteed to run at least once:
//      for {
//          // body
//          if !(count < 5) { break }
//      }
//    Printing: "Do-While running: " + count
//
// EXPECTED OUTPUT:
//   Do-While running: 10
//   Explanation: do-while executes the body at least once before checking the condition.

package main

import "fmt"

func main() {
	// TODO: Write your solution here.

}
