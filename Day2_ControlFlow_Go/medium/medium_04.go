// FILE: medium_04.go
// TOPIC: Control Flow & Bitwise - Day 2 (Go)
// DIFFICULTY: Medium
//
// TASK: Go's Special Bit-Clear Operator (&^)
// ------------------------------------------
//   - Bitwise OR (`|`) sets a bit.
//   - Bitwise XOR (`^`) toggles a bit.
//   - Bit Clear (`&^`) clears a bit. THIS IS UNIQUE TO GO!
//     (In C++, you have to do `& ~flag`. Go made a dedicated operator for it).
//
// 1. Declare: state := 0
// 2. Define a flag: const Alarm = 8 // Binary: 1000
//
// 3. SET the Alarm bit on `state` using `|=`.
// 4. Write an if-statement checking if Alarm is on (state & Alarm != 0). Print "Alarm is ON".
//
// 5. TOGGLE the Alarm bit using `^=`.
// 6. Write an if-statement checking if it's off. Print "Alarm is OFF".
//
// 7. SET the Alarm bit again (`|=`).
// 8. CLEAR the Alarm bit using Go's special operator: `state &^= Alarm`
// 9. Verify it is off by printing "Alarm successfully cleared".
//
// EXPECTED OUTPUT:
//   Alarm is ON
//   Alarm is OFF
//   Alarm successfully cleared

package main

import "fmt"

func main() {
	// TODO: Write your solution here.

}
