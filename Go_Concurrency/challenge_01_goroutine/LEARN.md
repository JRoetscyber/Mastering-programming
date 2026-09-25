# Challenge 1 - Your First Goroutine

## What is a goroutine?

Imagine you run a shop floor with die-cutting machines.

Without goroutines: you have **one machine**. You feed it Job 1, wait for it to finish, then feed it Job 2, then Job 3. Everything is in a line.

With goroutines: you have **multiple machines**. You press start on all three at the same time and walk away. They all work at the same time.

A **goroutine** (a task Go runs in the background) is how you start that second machine. You just put the word `go` in front of a function call.

```go
// Without go: runs here, blocks everything else
printJob(1)

// With go: starts immediately, runs in the background, you move on
go printJob(1)
```

---

## The one catch

Your `main()` function is also a goroutine - the first one. If `main()` exits, **everything exits**, including all the goroutines you started.

So if you do this:

```go
go printJob(1)
// main exits immediately here - goroutine never gets to print!
```

That is why in this challenge we use `time.Sleep` as a dirty workaround. We tell main: "wait 2 seconds so the goroutines have time to finish."

This is not the real fix - Challenge 2 shows you the proper tool for that.

---

## Run it

```bash
cd Go_Concurrency/challenge_01_goroutine
go run main.go
```

**What you will see:**
- The "without goroutines" section prints jobs 1, 2, 3 in perfect order, one at a time.
- The "with goroutines" section prints them out of order - because all three machines started at the same time and whichever finishes first prints first.

The order is not a bug. It is the point.

---

## Diagram

See `goroutine_flow.drawio.svg` in this folder.

---

## Key things to remember

| Thing | What it means in plain English |
|---|---|
| `go myFunc()` | Start `myFunc` in the background right now |
| Goroutines are cheap | You can have thousands - Go manages them, not the OS |
| `main()` does not wait | If main exits, all goroutines are killed |

---

## Hands-on Challenge & Test Suite

Now put your knowledge to the test in `challenge_01_goroutine/exercise/`!

### Scenario: Concurrent Shop Floor Inspection
You have multiple inspection stations checking die-cut rotary dies. If inspected sequentially, the factory floor halts. You must inspect them in parallel using goroutines.

### Files
- [`exercise/exercise.go`](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_01_goroutine/exercise/exercise.go): Contains the starter skeleton and `// TODO` tasks.
- [`exercise/exercise_test.go`](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_01_goroutine/exercise/exercise_test.go): Automated test suite verifying concurrency speedup and closure capture safety.

### How to Run & Test
```powershell
# 1. Navigate to the exercise directory
cd .\Go_Concurrency\challenge_01_goroutine\exercise

# 2. Run your code interactively
go run exercise.go

# 3. Test yourself (automated grading)
go test -v

# 4. Check for race conditions
go test -v -race
```

> [!TIP]
> In Go, if you launch goroutines inside a `for` loop, ensure each goroutine receives its own distinct copy of the loop item to avoid the classic closure capture bug!

