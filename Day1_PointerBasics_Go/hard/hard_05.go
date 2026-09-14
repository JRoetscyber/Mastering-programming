// FILE: hard_05.go
// TOPIC: Pointer Basics - Day 1 (Go)
// DIFFICULTY: Hard
//
// TASK: Swap what two pointers POINT TO using **int
// -------------------------------------------------
// This is different from swapping VALUES. Here you swap the POINTER TARGETS:
// after the operation, pa points to b and pb points to a.
//
// 1. Write a function:
//      func swapTargets(pp1, pp2 **int)
//    that swaps what *pp1 and *pp2 point to (i.e., *pp1, *pp2 = *pp2, *pp1).
//
// 2. In main():
//      a) Declare: a, b := 10, 20
//      b) Declare: pa, pb := &a, &b
//      c) Print what pa and pb point to BEFORE the swap.
//      d) Call swapTargets(&pa, &pb).
//      e) Print what pa and pb point to AFTER the swap.
//      f) Print a and b directly — they should be UNCHANGED (only the
//         pointers moved, not the values).
//
// EXPECTED OUTPUT:
//   Before: *pa=10, *pb=20
//   After : *pa=20, *pb=10
//   a=10, b=20  (values unchanged — only pointer targets swapped)

package main

import "fmt"

// TODO: Write swapTargets here.

func main() {
	// TODO: Write your solution here.

}
