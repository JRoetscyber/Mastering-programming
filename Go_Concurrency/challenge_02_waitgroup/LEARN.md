# Challenge 2 - WaitGroup: Waiting for Workers to Finish

## The problem from Challenge 1

In Challenge 1 we used `time.Sleep` to stop main from exiting too early. That is a hack.

What if one goroutine takes 3 seconds but another takes 10 seconds? You do not know how long to sleep. You need a proper way to say:

**"Wait here until ALL the goroutines are done."**

That is exactly what a `WaitGroup` does.

---

## How a WaitGroup works

Think of it as a scoreboard with a number on it.

- Every time you start a goroutine, you add 1 to the scoreboard. `wg.Add(1)`
- Every time a goroutine finishes, it takes 1 off the scoreboard. `wg.Done()`
- `wg.Wait()` means: stand here until the scoreboard reaches zero.

```
Start:      scoreboard = 0
Add job 1:  scoreboard = 1
Add job 2:  scoreboard = 2
Add job 3:  scoreboard = 3
Add job 4:  scoreboard = 4

(jobs run in the background...)

Job 2 done: scoreboard = 3
Job 4 done: scoreboard = 2
Job 1 done: scoreboard = 1
Job 3 done: scoreboard = 0  <-- wg.Wait() unblocks here
```

---

## The `defer` keyword

In the code you will see this pattern:

```go
func processTicket(ticketID string, wg *sync.WaitGroup) {
    defer wg.Done()
    // ... do work
}
```

`defer` means: **"run this line when the function exits, no matter what."**

So even if the function panics or returns early, `wg.Done()` will always be called. It is a safety net.

Always use `defer wg.Done()` at the TOP of the goroutine function - not at the bottom.

---

## The & symbol - passing a pointer

You will notice we pass `&wg` not just `wg`. The `&` means: "pass the address of this WaitGroup, not a copy of it."

If you passed a copy, the goroutine would be marking done on its own private copy - the original scoreboard would never change and `wg.Wait()` would block forever.

(Think of `&` as sharing the same whiteboard instead of handing each worker their own notepad.)

---

## Run it

```bash
cd Go_Concurrency/challenge_02_waitgroup
go run main.go
```

You will see all four tickets printed in a random order (because they run at the same time), then after all of them finish you will see "All tickets processed."

No `time.Sleep` hack needed.

---

## Diagram

See `waitgroup_flow.drawio.svg` in this folder.

---

## Key things to remember

| Code | What it does |
|---|---|
| `var wg sync.WaitGroup` | Create the scoreboard |
| `wg.Add(1)` | Tell it: one more worker is starting |
| `defer wg.Done()` | Tell it: I am done (runs automatically when function exits) |
| `wg.Wait()` | Block here until scoreboard = zero |
| `&wg` | Pass the real scoreboard, not a copy |

---

## Hands-on Challenge & Test Suite

Put your synchronization skills to the test in `challenge_02_waitgroup/exercise/`!

### Scenario: DiePro Batch Ticket Processor
When dispatching a batch of production job tickets in DiePro, each ticket takes variable time to compile. You cannot sleep. You must coordinate them deterministically using `sync.WaitGroup`.

### Files
- [`exercise/exercise.go`](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_02_waitgroup/exercise/exercise.go): Contains the starter skeleton and `// TODO` tasks.
- [`exercise/exercise_test.go`](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_02_waitgroup/exercise/exercise_test.go): Automated test suite verifying 100% completion, parallel execution speed, and high-load stress (100 concurrent tickets).

### How to Run & Test
```powershell
# 1. Navigate to the exercise directory
cd .\Go_Concurrency\challenge_02_waitgroup\exercise

# 2. Run your code interactively
go run exercise.go

# 3. Test yourself (automated grading)
go test -v

# 4. Check for race conditions
go test -v -race
```

> [!WARNING]
> Always pass the `sync.WaitGroup` as a pointer (`*sync.WaitGroup`, passing `&wg`). If you pass it by value, each goroutine modifies a local copy and `wg.Wait()` deadlocks forever!

