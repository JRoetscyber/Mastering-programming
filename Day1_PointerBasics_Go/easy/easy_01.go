// FILE: easy_01.go
// TOPIC: Pointer Basics - Day 1 (Go)
// DIFFICULTY: Easy
//
// TASK: Declare a pointer to an int and use it
// ---------------------------------------------
// 1. Declare an int variable called `score` and assign it the value 95.
// 2. Declare an *int pointer called `p` and make it point to `score`.
// 3. Print the VALUE of `score` through the pointer (dereference *p).
// 4. Print the MEMORY ADDRESS stored in `p` using fmt.Printf and the %p verb.
//
// EXPECTED OUTPUT (address will differ on your machine):
//   Value of score via pointer: 95
//   Address stored in p: 0x...
//
// NOTE: In Go, the short declaration  p := &score  both declares p as *int
//       and assigns it the address of score in one step.

package main

import "fmt"

func main() {
	// TODO: Write your solution here.
	var score int = 95;
	var p *int = &score;
	fmt.Println(*p);
	fmt.Println(p);
}
