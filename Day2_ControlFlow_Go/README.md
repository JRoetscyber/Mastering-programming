# Day 2 — Control Flow Deep Dive in Go

## Topic Overview

Go simplifies control flow compared to C++, but it has unique mechanics:

| Feature | C++ | Go |
|---|---|---|
| **While loop** | `while (cond) {}` | `for cond {}` (no `while` keyword) |
| **Do-while loop** | `do {} while (cond);` | `for { ... if !cond { break } }` |
| **Infinite loop** | `for (;;) {}` or `while (true) {}` | `for {}` |
| **Switch fallthrough** | Fallthrough by default (needs `break;`) | **Breaks by default** (needs `fallthrough` keyword) |
| **Multi-case switch** | `case 'A': case 'a':` | `case 'A', 'a':` (comma-separated!) |
| **Nested loop breakout** | `goto label;` | Labeled break: `break MyLoop` or `goto` |
| **Defer cleanup** | RAII or `goto cleanup` | `defer` statement (idiomatic Go cleanup) |

---

## Folder Structure

```
Day2_ControlFlow_Go/
|-- README.md          <- you are here
|-- SOLUTIONS.md       <- complete worked solutions
|-- easy/              <- 3 exercises
|-- medium/            <- 3 exercises
|-- hard/              <- 3 exercises
```

---

## How to Run

```bash
# In any directory
go run easy_01.go
```
