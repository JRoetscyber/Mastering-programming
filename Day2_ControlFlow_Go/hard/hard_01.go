// FILE: hard_01.go
// TOPIC: Control Flow Deep Dive - Day 2 (Go)
// DIFFICULTY: Hard
//
// TASK: Breaking Out of Nested Loops: Labeled Break vs goto in Go
// ---------------------------------------------------------------
// In C++, breaking out of nested loops requires `goto` or flag variables.
// Go supports `goto`, but Go ALSO gives you a cleaner native feature:
// **LABELED BREAK** (e.g. `break SearchLoop`)!
//
// 1. You have a 3x3 grid:
//      grid := [3][3]int{
//          {11, 12, 13},
//          {21, 99, 23},
//          {31, 32, 33},
//      }
//
// 2. Search for target `99` using nested loops.
// 3. Label the outer loop:  SearchLoop: for r := 0; r < 3; r++ { ... }
// 4. Inspect elements: print "Inspecting [r][c]: <val>"
// 5. When `99` is found, use `break SearchLoop` to instantly exit BOTH loops!
// 6. Print: "Found target 99 at [r][c]! Exited immediately."
//
// EXPECTED OUTPUT:
//   Inspecting [0][0]: 11
//   Inspecting [0][1]: 12
//   Inspecting [0][2]: 13
//   Inspecting [1][0]: 21
//   Inspecting [1][1]: 99
//   Found target 99 at [1][1]! Exited immediately.

package main

import "fmt"

func main() {
	// TODO: Write your solution here.

}
