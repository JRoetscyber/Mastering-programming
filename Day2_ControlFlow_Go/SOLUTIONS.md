# SOLUTIONS — Day 2: Control Flow Deep Dive in Go

> **Check these ONLY after attempting each exercise yourself!**

---

## easy_01.go — Short-Circuit Evaluation

```go
package main

import "fmt"

func main() {
    var ptr *int // nil by default

    if ptr != nil && *ptr > 50 {
        fmt.Println("Valid and greater than 50")
    } else {
        fmt.Println("Safety triggered: pointer is nil or value <= 50")
    }

    val := 85
    ptr = &val

    if ptr != nil && *ptr > 50 {
        fmt.Println("Valid and greater than 50")
    } else {
        fmt.Println("Safety triggered: pointer is nil or value <= 50")
    }
}
```

---

## easy_02.go — switch in Go

```go
package main

import "fmt"

func checkSeverity(severity rune) {
    fmt.Printf("Testing '%c': ", severity)
    switch severity {
    case 'E', 'e':
        fmt.Println("Level: ERROR (Critical action required)")
    case 'W', 'w':
        fmt.Println("Level: WARNING (Check system logs)")
    case 'I', 'i':
        fmt.Println("Level: INFO (Normal status)")
    default:
        fmt.Println("Level: UNKNOWN")
    }
}

func main() {
    checkSeverity('W')
    checkSeverity('e')
    checkSeverity('X')
}
```

---

## easy_03.go — While & Do-While in Go

```go
package main

import "fmt"

func main() {
    count := 10

    // While simulation
    for count < 5 {
        fmt.Println("While running:", count)
    }

    // Do-while simulation (executes body at least once)
    for {
        fmt.Println("Do-While running:", count)
        if !(count < 5) {
            break
        }
    }

    fmt.Println("Explanation: do-while executes the body at least once before checking the condition.")
}
```

---

## medium_01.go — Loop Control with continue & break

```go
package main

import "fmt"

func main() {
    readings := []int{12, -1, 45, 0, 999, 28, -5, 60, 999, 10}

    validTotal := 0
    validCount := 0

    for _, r := range readings {
        if r < 0 {
            fmt.Println("Skipping corrupted sensor reading:", r)
            continue
        }

        if r == 999 {
            fmt.Println("CRITICAL FAILURE (999) detected! Aborting loop.")
            break
        }

        fmt.Println("Valid reading:", r)
        validTotal += r
        validCount++
    }

    fmt.Println("Processed valid count:", validCount)
    fmt.Println("Valid total:", validTotal)
    if validCount > 0 {
        fmt.Println("Valid average:", validTotal/validCount)
    }
}
```

---

## medium_02.go — State Machine with iota & switch

```go
package main

import "fmt"

type State int

const (
    StateIdle State = iota
    StateConnecting
    StateConnected
    StateDisconnected
)

func main() {
    currentState := StateIdle
    attempts := 0

    for currentState != StateDisconnected {
        switch currentState {
        case StateIdle:
            fmt.Println("System IDLE. Initiating connection...")
            currentState = StateConnecting

        case StateConnecting:
            attempts++
            fmt.Printf("Attempting connection (attempt %d)...\n", attempts)
            if attempts == 3 {
                currentState = StateConnected
            }

        case StateConnected:
            fmt.Println("Connection established! Performing work...")
            fmt.Println("Work finished. Disconnecting...")
            currentState = StateDisconnected
        }
    }

    fmt.Println("Session closed gracefully.")
}
```

---

## medium_03.go — Guard Clauses

```go
package main

import "fmt"

func canProcessTransaction(balance, amount int, isActive, isFraud bool) bool {
    if !isActive {
        fmt.Println("Declined: Account inactive")
        return false
    }
    if isFraud {
        fmt.Println("Declined: Fraud alert triggered")
        return false
    }
    if amount <= 0 {
        fmt.Println("Declined: Invalid transaction amount")
        return false
    }
    if amount > balance {
        fmt.Println("Declined: Insufficient funds")
        return false
    }

    fmt.Println("Approved: Transaction processed!")
    return true
}

func main() {
    fmt.Print("Test A: ")
    canProcessTransaction(500, 100, false, false)

    fmt.Print("Test B: ")
    canProcessTransaction(500, 800, true, false)

    fmt.Print("Test C: ")
    canProcessTransaction(500, 250, true, false)
}
```

---

## hard_01.go — Labeled Break in Go

```go
package main

import "fmt"

func main() {
    grid := [3][3]int{
        {11, 12, 13},
        {21, 99, 23},
        {31, 32, 33},
    }

    target := 99
    foundRow, foundCol := -1, -1

SearchLoop:
    for r := 0; r < 3; r++ {
        for c := 0; c < 3; c++ {
            fmt.Printf("Inspecting [%d][%d]: %d\n", r, c, grid[r][c])
            if grid[r][c] == target {
                foundRow = r
                foundCol = c
                break SearchLoop // Instant exit out of BOTH loops!
            }
        }
    }

    if foundRow != -1 {
        fmt.Printf("Found target 99 at [%d][%d]! Exited immediately.\n", foundRow, foundCol)
    }
}
```

---

## hard_02.go — Cleanup with goto in Go

```go
package main

import "fmt"

func initializeServiceGoto(fail1, fail2, fail3 bool) bool {
    if fail1 {
        fmt.Println("Failed Stage 1")
        goto cleanup_none
    }
    fmt.Println("Stage 1 Acquired (Network Buffer)")

    if fail2 {
        fmt.Println("Failed Stage 2")
        goto cleanup_stage1
    }
    fmt.Println("Stage 2 Acquired (DB Connection)")

    if fail3 {
        fmt.Println("Failed Stage 3")
        goto cleanup_stage2
    }
    fmt.Println("Stage 3 Acquired (User Auth)")

    fmt.Println("Service Started Successfully!")
    return true

cleanup_stage2:
    fmt.Println("Releasing Stage 2 (DB Connection)")
cleanup_stage1:
    fmt.Println("Releasing Stage 1 (Network Buffer)")
cleanup_none:
    fmt.Println("Initialization aborted.")
    return false
}

func main() {
    initializeServiceGoto(false, true, false)
}
```

---

## hard_03.go — Bubble Sort in Go

```go
package main

import "fmt"

func main() {
    arr := []int{4, 1, 8, 3, 9, 2, 7}
    n := len(arr)
    passes := 0

    for i := 0; i < n-1; i++ {
        passes++
        swapped := false

        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j] // Clean Go swap
                swapped = true
            }
        }

        if !swapped {
            break
        }
    }

    fmt.Println("Passes taken:", passes)
    fmt.Println("Sorted slice:", arr)
}
```

---

## easy_04.go — Bitmasking with `if` and `iota`

```go
package main

import "fmt"

func main() {
    const (
        Exec = 1 << iota // 1
        Write            // 2
        Read             // 4
    )

    myPermission := 5

    // In Go, conditionals must strictly evaluate to a boolean!
    // So you must write `!= 0`.
    if (myPermission & Read) != 0 {
        fmt.Println("Has Read permission")
    }
    if (myPermission & Write) != 0 {
        fmt.Println("Has Write permission")
    }
    if (myPermission & Exec) != 0 {
        fmt.Println("Has Exec permission")
    }
}
```

---

## medium_04.go — Go's Special Bit-Clear Operator (&^)

```go
package main

import "fmt"

func main() {
    state := 0
    const Alarm = 8

    // SET bit
    state |= Alarm
    if (state & Alarm) != 0 {
        fmt.Println("Alarm is ON")
    }

    // TOGGLE bit
    state ^= Alarm
    if (state & Alarm) == 0 {
        fmt.Println("Alarm is OFF")
    }

    // SET again, then CLEAR using Go's dedicated bit-clear operator
    state |= Alarm
    state &^= Alarm // This does exactly what `&= ~Alarm` does in C++!
    
    if (state & Alarm) == 0 {
        fmt.Println("Alarm successfully cleared")
    }
}
```

---

## hard_04.go — Shift Operators and for loops (Hamming Weight)

```go
package main

import "fmt"

func main() {
    num := 29
    count := 0

    for num > 0 {
        if (num & 1) == 1 {
            count++
        }
        num >>= 1
    }

    fmt.Println("Number of set bits:", count)
}
```
