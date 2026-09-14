# 🎯 Mastering Programming

> A dual-track journey to systems & architecture mastery in **C++** and **Go** — one phase, one milestone, one hard-won lesson at a time.

This repo documents my progress following the official [roadmap.sh C++ Roadmap](https://roadmap.sh/cpp) and backend engineering paths. No shortcuts, no copy-pasted solutions — every exercise is written from scratch, broken, diagnosed, and implemented side-by-side in both C++ and Go.

---

## 📍 Current Progress & Roadmap Tracking

### 🟢 Fundamentals Track
- [x] **Introduction to Language** (`What is C++?`, `C vs C++`)
- [x] **Setting up your Environment** (`MinGW / GCC`, `Compilers`, `Build Flags`)
- [x] **Basic Operations** (`Arithmetic`, `Logical`, `Bitwise Operators`)
- [x] **Pointers Basics** (`Day1_PointerBasics/` & `Day1_PointerBasics_Go/`)
  - [x] Pointer declarations, address-of (`&`), and dereferencing (`*`)
  - [x] Pointer mutation, double indirection (`**`), pointer swaps
  - [x] Array traversal & pointer arithmetic (`p++`, pointer subtraction)
  - [x] Null pointer guards (`nullptr` / `nil`)
  - [x] `const` pointer vs pointer-to-`const` mechanics
- [x] **Control Flow & Statements** (`Day2_ControlFlow/` & `Day2_ControlFlow_Go/`)
  - [x] `if / else`, short-circuit evaluation (`&&` / `||`), and guard clauses
  - [x] `switch` statements: fallthrough vs break, jump table behavior
  - [x] `for`, `while`, and `do-while` loops (and Go's unified `for` model)
  - [x] `break` / `continue`, labeled break, and system cleanup patterns (`goto`)

---

### 🟡 In Progress / Next Up

- [ ] **Functions & Signatures**
  - [ ] Pass-by-value vs Pass-by-reference (`&`) vs Pass-by-pointer (`*`)
  - [ ] Return types & return value optimization (RVO)
  - [ ] Function Overloading & Default Arguments
  - [ ] Inline functions & call stack frames
- [ ] **References & Memory Model**
  - [ ] Deep dive: References (`T&`) vs Pointers (`T*`)
  - [ ] Object Lifetimes: Stack vs Heap allocation
  - [ ] Manual Dynamic Memory (`new` / `delete`, memory leaks)
- [ ] **Data Structures & Algorithms**
  - [ ] Big O Notation Analysis
  - [ ] Custom Linked Lists, Stacks, Queues, Hash Maps
  - [ ] Sorting (Bubble, Quick, Merge) & Searching (Binary Search)

---

## 📂 Repository Structure

```
mastering-programming/
├── Day1_PointerBasics/          <- C++ Pointer Mechanics (Easy, Medium, Hard)
├── Day1_PointerBasics_Go/       <- Go Pointer Mechanics & Heap Escape
├── Day2_ControlFlow/            <- C++ Control Flow, State Machines, Jump Tables
├── Day2_ControlFlow_Go/         <- Go Control Flow, Labeled Loops, Enums
├── phase-1-algorithmic-foundation/
│   ├── 01-big-o/
│   ├── 02-data-structures-cpp/
│   ├── 03-search-sort-algorithms/
│   └── 04-graph-traversal/
```

---

## 🧠 The Dual-Track Learning Philosophy

- **C++ (The Engine):** Teaches raw mechanics — memory layout, CPU cache lines, pointers, compile-time type safety, and direct hardware interaction.
- **Go (The Modern Counterpart):** Teaches systems architecture — clean concurrency, automatic memory management (garbage collector & escape analysis), and pragmatic simplicity.
- **Rosetta Stone Approach:** Every concept mastered in C++ is ported to Go to understand trade-offs in systems design.
- **Write to understand:** Leave personal insights, gotchas, and mental models directly in the code comments.

---

*Checkboxes get ticked as milestones are truly understood and proven in code, not just read about.*
