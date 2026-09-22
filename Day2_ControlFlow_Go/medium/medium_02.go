// FILE: medium_02.go
// TOPIC: Control Flow Deep Dive - Day 2 (Go)
// DIFFICULTY: Medium
//
// TASK: State Machine using iota enum and switch in Go
// ---------------------------------------------------
// In Go, enums are created using constants with `iota`.
//
// 1. Define states:
//      type State int
//      const (
//          StateIdle State = iota
//          StateConnecting
//          StateConnected
//          StateDisconnected
//      )
//
// 2. In main(), start with:
//      currentState := StateIdle
//      attempts := 0
//
// 3. Run a `for currentState != StateDisconnected` loop:
//      Use a `switch currentState`:
//        - case StateIdle:
//            print "System IDLE. Initiating connection..."
//            transition to StateConnecting
//
//        - case StateConnecting:
//            attempts++
//            print "Attempting connection (attempt " << attempts << ")..."
//            if attempts == 3 { transition to StateConnected }
//
//        - case StateConnected:
//            print "Connection established! Performing work..."
//            print "Work finished. Disconnecting..."
//            transition to StateDisconnected
//
// 4. Print "Session closed gracefully." after exiting.
//
// EXPECTED OUTPUT:
//   System IDLE. Initiating connection...
//   Attempting connection (attempt 1)...
//   Attempting connection (attempt 2)...
//   Attempting connection (attempt 3)...
//   Connection established! Performing work...
//   Work finished. Disconnecting...
//   Session closed gracefully.

package main

import "fmt"

// TODO: Define State enum here.

func main() {
	// TODO: Write your solution here.

}
