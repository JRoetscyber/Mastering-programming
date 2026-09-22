# Day 2 — Control Flow Deep Dive

## Topic Overview

Control flow determines **which path** the CPU takes when executing instructions. 
Writing robust C++ requires mastering branching and looping beyond simple syntax.

### Key Concepts Covered

1. **`if / else` & Short-Circuiting**:
   - In `A && B`, if `A` is false, `B` is **never evaluated**.
   - In `A || B`, if `A` is true, `B` is **never evaluated**.
   - Guard clauses (early exit pattern) simplify complex logic.

2. **`switch` Statements**:
   - Compiles to a fast O(1) jump table under the hood when values are dense.
   - Intentional fallthrough (`[[fallthrough]];`) vs missing `break;` bugs.

3. **Loops (`for`, `while`, `do while`)**:
   - `for`: Best when iteration count or boundaries are known.
   - `while`: Best when waiting on a condition (event loop, input).
   - `do while`: Guaranteed to execute **at least once** before checking condition (great for menus/retry loops).
   - `break` vs `continue`.

4. **`goto`**:
   - Generally considered harmful for general control flow (spaghetti code).
   - **Legitimate real-world use**: Multi-level loop breakouts and single-point error cleanup (used extensively in the Linux kernel and low-level drivers).

---

## Folder Structure

```
Day2_ControlFlow/
|-- README.md          <- you are here
|-- SOLUTIONS.md       <- complete worked solutions (check after attempting)
|-- easy/              <- 3 exercises
|-- medium/            <- 3 exercises
|-- hard/              <- 3 exercises
```

---

## How to Compile & Run

```bash
# Using g++ / MinGW
g++ -std=c++17 -Wall -Wextra easy/easy_01.cpp -o easy_01.exe
./easy_01.exe

# Or MSVC
cl /std:c++17 /W4 easy/easy_01.cpp
easy_01.exe
```
