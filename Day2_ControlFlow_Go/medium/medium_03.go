// FILE: medium_03.go
// TOPIC: Control Flow Deep Dive - Day 2 (Go)
// DIFFICULTY: Medium
//
// TASK: Guard Clauses & Early Return (Idiomatic Go Pattern)
// ---------------------------------------------------------
// In Go, the "Guard Clause" pattern is the foundational style for the
// entire standard library! (Always "handle errors early and return").
//
// 1. Write a function:
//      func canProcessTransaction(balance, amount int, isActive, isFraud bool) bool
//
// 2. Use ONLY guard clauses (fail fast):
//      - if !isActive -> print "Declined: Account inactive", return false
//      - if isFraud   -> print "Declined: Fraud alert triggered", return false
//      - if amount <= 0 -> print "Declined: Invalid transaction amount", return false
//      - if amount > balance -> print "Declined: Insufficient funds", return false
//      - print "Approved: Transaction processed!", return true
//
// 3. Test with the 3 scenarios in main().
//
// EXPECTED OUTPUT:
//   Test A: Declined: Account inactive
//   Test B: Declined: Insufficient funds
//   Test C: Approved: Transaction processed!

package main

import "fmt"

// TODO: Write canProcessTransaction here.

func main() {
	// TODO: Write your solution here.

}
