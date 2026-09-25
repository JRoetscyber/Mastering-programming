# Challenge 4 - Worker Pool: The Real Production Pattern

## What is a worker pool?

Challenges 1-3 created one goroutine per job. If you have 5000 tickets to price, you do NOT want 5000 goroutines all running at the same time - that would flood memory.

A **worker pool** solves this. You create a **fixed number of workers** (say, 3), and they all share a **queue of jobs**. Each worker grabs a job, does it, sends back the result, then grabs the next one. Workers are never idle when there is work to do.

**Shop floor example:** you have 3 machines. When a machine finishes a job it picks the next ticket off the stack. The stack (queue) does not grow based on how many jobs there are - it is always 3 machines.

---

## The pattern in three steps

### Step 1: Create the channels

```go
jobs    := make(chan Job, 10)    // the stack of work
results := make(chan Result, 10) // where workers drop finished work
```

The `10` is a buffer - the channel can hold up to 10 items without blocking. This makes feeding jobs in smoother.

### Step 2: Start the workers FIRST

```go
for i := 1; i <= 3; i++ {
    wg.Add(1)
    go worker(i, jobs, results, &wg)
}
```

The workers sit and wait on the `jobs` channel. They are ready but idle - nothing in the queue yet.

### Step 3: Feed the jobs in, then close the channel

```go
for _, t := range tickets {
    jobs <- t
}
close(jobs) // tell workers: no more work coming
```

`close(jobs)` is the signal. When the workers drain the channel and see it is closed, the `for job := range jobs` loop exits naturally and the worker calls `wg.Done()`.

---

## The closer goroutine trick

```go
go func() {
    wg.Wait()
    close(results)
}()
```

This small anonymous goroutine waits until ALL workers are done, then closes the results channel. This lets main use `for r := range results` - it loops until the results channel is closed, then stops. Clean, no counting needed.

---

## Run it

```bash
cd Go_Concurrency/challenge_04_worker_pool
go run main.go
```

Watch the output. You will see which worker picked which ticket. With 3 workers and 8 tickets, each worker handles 2-3 tickets. Run it a few times - the assignment of tickets to workers changes each time.

---

## Diagram

See `worker_pool_flow.drawio.svg` in this folder.

---

## How this connects to DiePro

When DiePro Edge needs to:
- Price 200 quotes overnight
- Generate PDFs for 50 finished tickets
- Send dispatch notifications to 30 customers

...it uses this exact pattern. A handful of workers chew through the queue without spawning thousands of goroutines.

---

## Key things to remember

| Concept | Plain English |
|---|---|
| Buffered channel `make(chan T, N)` | The pipe can hold N items before it blocks |
| `close(jobs)` | Signal workers: no more work, stop looping |
| `for job := range jobs` | Keep reading until the channel is closed |
| Closer goroutine | Wait for all workers, then close results so main can range over it |

---

## Hands-on Challenge & Test Suite

Put your worker pool skills to the test in `challenge_04_worker_pool/exercise/`!

### Scenario: Production DiePro Batch CAD Tooling Worker Pool
When hundreds of die tooling CAD jobs arrive overnight in DiePro, you cannot spawn hundreds of goroutines at once. You must configure a worker pool of fixed size that consumes a buffered queue, coordinates with a closer goroutine, and records per-worker distribution statistics.

### Files
- [`exercise/exercise.go`](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_04_worker_pool/exercise/exercise.go): Contains the starter skeleton, worker structs, workload distribution tracker, and `// TODO` tasks.
- [`exercise/exercise_test.go`](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_04_worker_pool/exercise/exercise_test.go): Automated test suite verifying full job completion, multi-worker load distribution, and clean channel closure.

### How to Run & Test
```powershell
# 1. Navigate to the exercise directory
cd .\Go_Concurrency\challenge_04_worker_pool\exercise

# 2. Run your code interactively
go run exercise.go

# 3. Test yourself (automated grading)
go test -v

# 4. Check for race conditions
go test -v -race
```

> [!IMPORTANT]
> The closer goroutine pattern (`go func() { wg.Wait(); close(results) }()`) is essential! If you close the results channel too early, workers will panic on write. If you forget to close it, `for r := range results` will deadlock forever.

