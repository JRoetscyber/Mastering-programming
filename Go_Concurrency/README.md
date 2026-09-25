# Go Concurrency - Hands-On Learning & Challenge Series

Work through these challenges in order. Each module includes a concept walkthrough, a live demonstration script, an interactive coding challenge (`// TODO`), and an automated unit test suite.

---

## 🗂️ Module Directory

| Module | Core Concept | Exercise & Testing |
|---|---|---|
| [**Challenge 1 - Goroutines**](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_01_goroutine/LEARN.md) | `go` keyword, non-determinism, closure capture safety | [`exercise/exercise.go`](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_01_goroutine/exercise/exercise.go) |
| [**Challenge 2 - WaitGroup**](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_02_waitgroup/LEARN.md) | `sync.WaitGroup`, `defer wg.Done()`, pointer passing | [`exercise/exercise.go`](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_02_waitgroup/exercise/exercise.go) |
| [**Challenge 3 - Channels**](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_03_channels/LEARN.md) | Channels, directional constraints (`chan<-`), data streaming | [`exercise/exercise.go`](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_03_channels/exercise/exercise.go) |
| [**Challenge 4 - Worker Pool**](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_04_worker_pool/LEARN.md) | Bounded worker pools, buffered queues, closer goroutines | [`exercise/exercise.go`](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/challenge_04_worker_pool/exercise/exercise.go) |

---

## 🛠️ How to Learn and Test Yourself

### Step 1: Read the Concept & Inspect the Diagram
Read `LEARN.md` and check the `.drawio.svg` diagram inside the challenge folder.

### Step 2: Run the Demo
Run the included demonstration script to observe the behavior:
```powershell
go run .\Go_Concurrency\challenge_01_goroutine\main.go
go run .\Go_Concurrency\challenge_02_waitgroup\main.go
go run .\Go_Concurrency\challenge_03_channels\main.go
go run .\Go_Concurrency\challenge_04_worker_pool\main.go
```

### Step 3: Solve the Exercise
Navigate to the module's `exercise/` folder and implement the tasks marked with `// TODO` in `exercise.go`.

```powershell
# Example for Challenge 1:
cd .\Go_Concurrency\challenge_01_goroutine\exercise
go run exercise.go
```

### Step 4: Run the Automated Test Suite (Self-Grading)
Run `go test -v` to test your implementation against concurrency speedup, edge cases, and timeouts:
```powershell
# Run tests for current exercise
go test -v

# Run with Go's race detector to ensure thread-safety
go test -v -race
```

### Step 5: Test All Modules at Once
From the `Go_Concurrency/` root directory, you can run tests across all challenges:
```powershell
cd .\Go_Concurrency
go test -v ./...
```

---

## 💡 Stuck? Check the Solutions
Worked solutions with in-depth senior developer mentor notes are available in [**`SOLUTIONS.md`**](file:///c:/Users/User/source/repos/Mastering-programming/Go_Concurrency/SOLUTIONS.md). Try to attempt the exercises first before checking!

---

## 🧠 The Big Picture: Progression

```
Challenge 1: go keyword          --> Spawn a background worker
Challenge 2: WaitGroup           --> Wait deterministically until all workers finish
Challenge 3: Channels            --> Safely pass data back and forth without race conditions
Challenge 4: Worker Pool         --> Industrial-grade pattern: fixed workers processing dynamic queues
```
