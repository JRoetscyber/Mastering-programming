// FILE: easy_02.go
// TOPIC: Control Flow Deep Dive - Day 2 (Go)
// DIFFICULTY: Easy
//
// TASK: switch in Go — Multi-case and break by default
// -----------------------------------------------------
// In Go:
// 1. You DO NOT need `break` at the end of each case! Go breaks automatically.
// 2. To handle multiple values in one case, use commas:  case 'E', 'e':
// 3. If you WANT fallthrough, you must explicitly write `fallthrough`.
//
// 1. Write a function `checkSeverity(severity rune)`:
//    - case 'E', 'e': print "Level: ERROR (Critical action required)"
//    - case 'W', 'w': print "Level: WARNING (Check system logs)"
//    - case 'I', 'i': print "Level: INFO (Normal status)"
//    - default: print "Level: UNKNOWN"
//
// 2. Test in main with 'W', 'e', and 'X'.
//
// EXPECTED OUTPUT:
//   Testing 'W': Level: WARNING (Check system logs)
//   Testing 'e': Level: ERROR (Critical action required)
//   Testing 'X': Level: UNKNOWN

package main

import "fmt"

// TODO: Write checkSeverity here.

func main() {
	// TODO: Write your solution here.

}
