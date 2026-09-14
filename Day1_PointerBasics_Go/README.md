# Day 1 — Pointer Basics in Go (`&` and `*`)

## Topic Overview

Go has pointers, but they work differently from C/C++ in three important ways:

| Feature                    | C++                   | Go                              |
|----------------------------|-----------------------|---------------------------------|
| Pointer declaration        | `int* p`              | `var p *int`  or  `p := &x`   |
| Address-of                 | `&x`                  | `&x`  (same)                   |
| Dereference                | `*p`                  | `*p`  (same)                   |
| Null pointer               | `nullptr`             | `nil`                           |
| Pointer arithmetic         | `p++`, `p+2`         | **NOT ALLOWED** (use `unsafe` for that) |
| Struct member via pointer  | `p->field`            | `p.field`  (auto-deref, no `->`) |
| Memory management          | manual (`new`/`delete`)| **automatic** (garbage collector)|

### Quick Mental Model

```go
x := 42        // int variable
p := &x        // p is *int — holds the address of x
fmt.Println(*p) // dereference: prints 42
*p = 99        // write through pointer: x is now 99
fmt.Println(x)  // prints 99
```

---

## Folder Structure

```
Day1_PointerBasics_Go/
|-- README.md          <- you are here
|-- SOLUTIONS.md       <- full worked solutions (check AFTER attempting!)
|-- easy/              <- 5 exercises  (declaration, address-of, dereferencing)
|-- medium/            <- 5 exercises  (modify via pointer, **T, swap, slice element)
|-- hard/              <- 5 exercises  (nil guard, pointer funcs, aliasing, heap escape, **swap)
|-- insane/            <- 5 exercises  (***T, unsafe arithmetic, closures, struct chains, puzzle)
```

---

## How to Use This Folder

1. **Pick a difficulty level** — start at `easy/`.
2. **Read the comment block** at the top of the file.
3. **Write your solution** inside the `// TODO` section.
4. **Run** with the Go toolchain:
   ```bash
   # From inside the easy/ folder (or whichever level)
   go run easy_01.go

   # Or build first:
   go build easy_01.go && ./easy_01
   ```
5. **Check your output** against what the comment block says.
6. **Stuck?** Open `SOLUTIONS.md` — solutions are labelled by filename.

---

## Key Go Pointer Syntax at a Glance

```go
var p *int          // declare a nil pointer to int
x := 100
p = &x              // p now holds the address of x
fmt.Println(*p)     // dereference: read the value (100)
*p = 200            // dereference: write a value (x is now 200)

var pp **int = &p   // pointer to pointer
fmt.Println(**pp)   // double dereference (200)
```

---

## Ground Rules for This Set

- No dynamic memory beyond the Go runtime (no `new(T)` until hard exercises).
- Each file is `package main` and runs with `go run <file>`.
- No goroutines, channels, or interfaces — pure pointer mechanics.
- `fmt.Printf("%p\n", p)` prints a pointer as a hex address.
