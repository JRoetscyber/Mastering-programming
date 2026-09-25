# Challenge 3 - Channels: Goroutines Talking Back to You

## The problem

WaitGroups tell you when the workers are done. But what if the workers calculated something and you need the result?

You cannot just set a variable inside a goroutine and read it from main - two goroutines touching the same variable at the same time causes a **data race** (two machines writing to the same bit of paper at the same time = mess).

You need a safe way for a goroutine to **send a value back** to the main routine.

That is what a **channel** is.

---

## What is a channel?

A channel is a pipe.

```
goroutine 1  ---[pipe]---> main
goroutine 2  ---[pipe]---> main
goroutine 3  ---[pipe]---> main
```

You create the pipe in main, hand it to each goroutine, and each goroutine drops its result into the pipe. Main reads results out of the pipe as they arrive.

```go
// Create a channel that carries PricingResult values
results := make(chan PricingResult)

// Inside a goroutine: SEND a value into the pipe
results <- myResult

// In main: RECEIVE (wait for) a value from the pipe
r := <-results
```

The arrow `<-` always points toward where the value is going.
- `results <- value`   means "push value INTO results"
- `value := <-results` means "pull value OUT OF results"

---

## Reading from a channel blocks

When you write `r := <-results`, main will **wait right there** until a goroutine drops something into the pipe. This replaces `time.Sleep` AND gives you the data back.

---

## Directional channels (write-only / read-only)

In the function signature you will see:

```go
func calculatePrice(results chan<- PricingResult)
//                              ^^^
//                         write-only - can only push in
```

`chan<- T` means this end of the pipe can only push IN.
`<-chan T` means this end can only read OUT.

This prevents a goroutine from accidentally reading its own pipe and getting confused. The compiler will catch the mistake for you.

---

## Run it

```bash
cd Go_Concurrency/challenge_03_channels
go run main.go
```

You will see the three prices printed as each goroutine finishes. Order may vary run to run.

---

## Diagram

See `channels_flow.drawio.svg` in this folder.

---

## Key things to remember

| Code | Plain English |
|---|---|
| `make(chan MyType)` | Create a pipe that carries MyType values |
| `pipe <- value` | Push a value into the pipe |
| `value := <-pipe` | Wait for and pull a value out of the pipe |
| `chan<- T` | Write-only end of a pipe (function parameter) |
| `<-chan T` | Read-only end of a pipe (function parameter) |

---

## Hands-on Challenge & Test Suite

Put your channel communication skills to the test in `challenge_03_channels/exercise/`!

### Scenario: Quoting Engine Material Pricing Pipeline
In DiePro's estimating engine, pricing calculations happen concurrently across customer tickets. Instead of modifying shared variables, workers stream results down directional channels. You must compute setup fees and material surcharges for high-value orders and collect them via channels.

### Files
- [`exercise/exercise.go`](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_03_channels/exercise/exercise.go): Contains the starter skeleton, pricing business logic requirements, and `// TODO` tasks.
- [`exercise/exercise_test.go`](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_03_channels/exercise/exercise_test.go): Automated test suite verifying channel communication, timeout safety, and math correctness.

### How to Run & Test
```powershell
# 1. Navigate to the exercise directory
cd .\Go_Concurrency\challenge_03_channels\exercise

# 2. Run your code interactively
go run exercise.go

# 3. Test yourself (automated grading)
go test -v

# 4. Check for race conditions
go test -v -race
```

> [!TIP]
> Use directional channel signatures (`chan<- T` for senders, `<-chan T` for receivers) to enforce compile-time safety so functions don't read from or write to pipes they shouldn't!

