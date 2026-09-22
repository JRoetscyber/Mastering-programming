// FILE: easy_04.go
// TOPIC: Control Flow & Bitwise - Day 2 (Go)
// DIFFICULTY: Easy
//
// TASK: Bitmasking with `if` and `iota`
// -------------------------------------
// In Go, `iota` combined with the left-shift operator (`<<`) is the standard
// way to create binary flags!
//
// 1. Define constants using iota:
//      const (
//          Exec = 1 << iota  // 1 << 0 = 1  (001)
//          Write             // 1 << 1 = 2  (010)
//          Read              // 1 << 2 = 4  (100)
//      )
//
// 2. A file has permissions set to 5 (Binary: 101).
//      myPermission := 5
//
// 3. Write three separate `if` statements using the bitwise AND operator (`&`).
//      - If (myPermission & Read) != 0, print "Has Read permission"
//      - etc.
//    (Note: Go requires explicit comparison `!= 0`, unlike C++)
//
// EXPECTED OUTPUT:
//   Has Read permission
//   Has Exec permission

package main

import "fmt"

func main() {
	// TODO: Write your solution here.

}
