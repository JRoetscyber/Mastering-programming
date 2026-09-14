# Day 1 — Pointer Basics (`&` and `*`)

## Topic Overview

Pointers are variables that store **memory addresses**. Mastering them is essential
for understanding how C++ (and C) manage memory, pass data efficiently, and interact
with hardware.

This exercise set covers three core concepts:

| Concept               | Syntax     | What it does                                     |
|-----------------------|------------|--------------------------------------------------|
| **Pointer declaration**| `int* p;` | Declares `p` as a pointer to an `int`            |
| **Address-of operator**| `&var`    | Produces the memory address of `var`             |
| **Dereference operator**| `*p`     | Accesses the value *at* the address stored in `p`|

### Quick Mental Model

```
int x = 42;
int* p = &x;   // p holds the address of x (e.g. 0x7ffd...)
*p = 99;       // writes 99 into whatever memory p points at -> x is now 99
```

---

## Folder Structure

```
Day1_PointerBasics/
|-- README.md          <- you are here
|-- SOLUTIONS.md       <- full worked solutions (check AFTER attempting!)
|-- easy/              <- 5 exercises  (pointer declaration, reading addresses, dereferencing)
|-- medium/            <- 5 exercises  (modifying via pointer, ptr-to-ptr, swap, array element)
|-- hard/              <- 5 exercises  (pointer arithmetic, null/dangling, functions, const)
|-- insane/            <- 5 exercises  (multi-level indirection, 2D arrays, returning pointers,
                                        structs, aliasing puzzles)
```

---

## How to Use This Folder

1. **Pick a difficulty level** - start at `easy/` unless you already know the basics.
2. **Open the `.cpp` file** for the exercise you want to attempt.
3. **Read the comment block** at the top - it describes the task precisely.
4. **Write your solution** inside the `// TODO` section in `main()`.
5. **Compile and run**:
   ```bash
   # g++ (MinGW / WSL)
   g++ -std=c++17 -Wall -Wextra -o out easy_01.cpp && ./out

   # MSVC (Developer Command Prompt)
   cl /std:c++17 /W4 easy_01.cpp && easy_01.exe
   ```
6. **Check your output** against what the comment block says should print.
7. **Stuck?** Open `SOLUTIONS.md` - solutions are labelled by filename.

---

## Ground Rules for This Set

- No `new` / `delete` - dynamic memory comes later.
- No classes or linked lists.
- No `std::vector` - work with raw arrays where needed.
- Every file is self-contained and compiles on its own.

---

> "Pointers are the most powerful feature of C++... and the most dangerous.
>  Master the basics here before you touch heap memory."
