# SOLUTIONS — Day 1: Pointer Basics in Go (`&` and `*`)

> **Check your work only AFTER you have attempted each exercise yourself.**
> Solutions are labelled by filename.

---

## easy_01.go — Declare and use a *int pointer

```go
package main

import "fmt"

func main() {
    score := 95
    p     := &score

    fmt.Println("Value of score via pointer:", *p)
    fmt.Printf("Address stored in p: %p\n", p)
}
```

**Explanation**
`p := &score` uses short declaration to create `p` as a `*int` and assign it the
address of `score` in one step. `%p` is the fmt verb for pointer (prints hex address).
`*p` dereferences the pointer to read the `int` value stored at that address.

---

## easy_02.go — Address-of operator on multiple types

```go
package main

import "fmt"

func main() {
    name  := "Alice"
    count := 7
    price := 3.99

    fmt.Printf("Address of name  : %p\n", &name)
    fmt.Printf("Address of count : %p\n", &count)
    fmt.Printf("Address of price : %p\n", &price)
}
```

**Explanation**
`%p` works on any pointer type in Go without a cast — a significant improvement
over C++ where you need `static_cast<void*>` for char pointers.
Notice that `&name` yields a `*string`, `&count` a `*int`, etc.
The addresses themselves are determined by the stack layout at runtime.

---

## easy_03.go — Pointer to float64 + type inspection

```go
package main

import "fmt"

func main() {
    temperature := 36.6
    pTemp       := &temperature

    fmt.Println("Direct value  :", temperature)
    fmt.Println("Via pointer   :", *pTemp)
    fmt.Printf("Type of pTemp : %T\n", pTemp)
}
```

**Explanation**
`%T` prints the concrete Go type — here `*float64`. This is a useful debugging
tool. Note that all pointers are the same size (8 bytes on a 64-bit platform)
regardless of what they point to, just like in C++.

---

## easy_04.go — Pointer to bool + address equality

```go
package main

import "fmt"

func main() {
    active := true
    p      := &active

    fmt.Println("Value via pointer:", *p)

    if p == &active {
        fmt.Println("Address check    : Match")
    } else {
        fmt.Println("Address check    : No match")
    }
}
```

**Explanation**
`p == &active` compares the two pointer VALUES (addresses), not the bools they
might point to. Since `p` was set to `&active`, they are identical. This is
pointer equality, analogous to `p == &active` in C++.

---

## easy_05.go — Re-pointing a pointer

```go
package main

import "fmt"

func main() {
    a, b := 10, 20
    p    := &a

    fmt.Println("p points to    :", *p)

    p = &b // p now holds the address of b

    fmt.Println("p now points to:", *p)
}
```

**Explanation**
Assigning `p = &b` overwrites the address stored in `p`. It does NOT change `a`
or `b`. `a` remains 10 throughout. The pointer is just a variable holding an
address — changing the variable doesn't affect the thing it used to point to.

---
---

## medium_01.go — Modify a variable through its pointer

```go
package main

import "fmt"

func main() {
    lives := 3
    p     := &lives

    fmt.Println("lives =", lives)

    *p -= 1 // subtract 1 via pointer
    fmt.Println("lives =", lives)

    *p = 0 // set to zero via pointer
    fmt.Println("lives =", lives)
}
```

**Explanation**
`*p -= 1` dereferences `p` to reach the `int` stored at `&lives`, then subtracts
1 from it. The change is visible through `lives` because `p` and `lives` share
the same memory address.
In Go, `*p--` is a syntax error — `--` is a statement (not an expression) and the
parser tries to decrement `p` (the pointer) which is illegal. Use `*p -= 1` or
`(*p)--`.

---

## medium_02.go — Pointer to a pointer (**int)

```go
package main

import "fmt"

func main() {
    x  := 42
    p  := &x
    pp := &p

    fmt.Printf("x = %d  |  *p = %d  |  **pp = %d\n", x, *p, **pp)
    fmt.Printf("&x = %p | *pp = %p\n", &x, *pp)
}
```

**Explanation**
Dereference layers:
- `pp` → address of `p`
- `*pp` → the value stored in `p`, which is `&x`
- `**pp` → the value at `&x`, which is `x` (42)

Both `&x` and `*pp` resolve to the same address — the location of `x` in memory.

---

## medium_03.go — Swap using raw pointers

```go
package main

import "fmt"

func main() {
    a, b := 100, 200
    pa   := &a
    pb   := &b

    fmt.Printf("Before: a = %d, b = %d\n", a, b)

    temp := *pa // save a's value
    *pa = *pb   // copy b's value into a
    *pb = temp  // copy saved value into b

    fmt.Printf("After : a = %d, b = %d\n", a, b)
}
```

**Explanation**
The temporary `temp` holds an `int` value (not a pointer). `*pa = *pb` copies the
integer value at `pb`'s address into the memory at `pa`'s address. This is the
manual pointer-based swap — Go's built-in `a, b = b, a` does the same thing more
elegantly, but doing it by hand proves you understand what's happening under the hood.

---

## medium_04.go — Pointer to a slice element

```go
package main

import "fmt"

func main() {
    scores := []int{10, 20, 30, 40, 50}
    p      := &scores[2] // address of the third element

    fmt.Println("Pointed-to value:", *p)

    *p = 99 // modify through pointer

    fmt.Println("Slice after edit:", scores)
}
```

**Explanation**
`&scores[2]` yields a `*int` pointing directly into the slice's backing array.
Writing `*p = 99` mutates that array element in place.
**Danger:** if `append` is called on `scores` and the slice needs to grow, Go
allocates a new backing array and copies the data. The old pointer `p` would then
point at the OLD (now abandoned) array — a dangling pointer equivalent in Go.
Never hold element pointers across potential appends.

---

## medium_05.go — Read/write through **int chain

```go
package main

import "fmt"

func main() {
    value := 5
    p1    := &value
    p2    := &p1

    fmt.Println("Via p2 before:", **p2)

    **p2 *= 3 // multiply value through two levels of indirection

    fmt.Println("Via p2 after :", **p2)
    fmt.Println("value direct :", value)
}
```

**Explanation**
`**p2` walks two levels: `p2` → `p1` → `value`. The compound assignment `**p2 *= 3`
is identical to `value *= 3` but exercises both levels of the pointer chain.
All three expressions (`value`, `*p1`, `**p2`) are aliases for the exact same
memory cell.

---
---

## hard_01.go — nil pointer guard

```go
package main

import (
    "fmt"
    "strconv"
)

func safeRead(p *int) string {
    if p == nil {
        return "nil pointer - cannot read"
    }
    return "value = " + strconv.Itoa(*p)
}

func main() {
    var p *int // zero value is nil

    fmt.Println(safeRead(p))

    x := 77
    p = &x

    fmt.Println(safeRead(p))
}
```

**Explanation**
The zero value for any pointer type in Go is `nil`. Unlike C++, where an
uninitialised pointer holds a garbage address (making accidental dereferences
silently corrupt random memory), Go's nil pointer produces a predictable runtime
panic when dereferenced — which is at least debuggable.
The guard `if p == nil` is the canonical Go defensive pattern before
dereferencing any pointer you did not just assign.

**Bonus answer:** In C++, `int* p;` in a local scope holds a garbage address —
any dereference is undefined behaviour that may silently corrupt memory.
Go always zero-initialises: `var p *int` is guaranteed to be `nil`.
`nil` dereference panics immediately and predictably; garbage dereference in C++
may corrupt data and not crash until much later.

---

## hard_02.go — Pass pointer to function

```go
package main

import "fmt"

func doubleIt(ptr *int) {
    *ptr *= 2
}

func addFive(ptr *int) {
    *ptr += 5
}

func main() {
    n := 10

    doubleIt(&n)
    fmt.Println("After doubleIt:", n)

    addFive(&n)
    fmt.Println("After addFive :", n)

    doubleIt(&n)
    fmt.Println("After doubleIt:", n)
}
```

**Explanation**
Go is strictly pass-by-value. Without a pointer, `doubleIt` would receive a copy
of `n` and the original would be unchanged. Passing `&n` gives the function the
address of `n`. Writing `*ptr *= 2` dereferences that address and modifies the
original variable. This is the Go way to achieve what C++ call-by-reference does.

---

## hard_03.go — Pointer aliasing

```go
package main

import "fmt"

func main() {
    x := 42
    p := &x
    q := &x // both point to the same memory cell

    fmt.Printf("Initial : *p=%d, *q=%d\n", *p, *q)

    *p = 100
    fmt.Printf("After *p=100: *q=%d\n", *q)

    *q = 200
    fmt.Printf("After *q=200: x=%d\n", x)

    fmt.Println("p == q:", p == q)
}
```

**Explanation**
`p` and `q` are independent pointer variables, but they both hold the same
address (`&x`). A write through either one is therefore visible through the
other AND through the original variable name `x`. This is aliasing.
Aliasing is a common source of hard-to-find bugs — if two pointers unexpectedly
share an address, modifying through one causes surprising changes through the
other.

---

## hard_04.go — Go heap escape (returning pointer to local)

```go
package main

import "fmt"

func newInt(val int) *int {
    n := val   // n lives on the stack... until Go decides to escape it
    return &n  // compiler detects this and allocates n on the HEAP
}

func main() {
    p := newInt(10)
    fmt.Println("*p after alloc  :", *p)

    *p = 99
    fmt.Println("*p after modify :", *p)

    q := newInt(55)
    fmt.Printf("*p=%d, *q=%d  (independent)\n", *p, *q)
}
```

**Explanation**
Go's escape analysis runs at compile time. When it sees that `&n` outlives the
function (because it's returned), it allocates `n` on the heap automatically.
You can verify with: `go build -gcflags="-m" hard_04.go` — the compiler will
print `n escapes to heap`.
In C++, doing the equivalent (`return &localVar;`) is undefined behaviour because
the stack frame is torn down on return. Go's GC makes this safe.

---

## hard_05.go — Swap pointer targets using **int

```go
package main

import "fmt"

func swapTargets(pp1, pp2 **int) {
    *pp1, *pp2 = *pp2, *pp1
}

func main() {
    a, b   := 10, 20
    pa, pb := &a, &b

    fmt.Printf("Before: *pa=%d, *pb=%d\n", *pa, *pb)

    swapTargets(&pa, &pb)

    fmt.Printf("After : *pa=%d, *pb=%d\n", *pa, *pb)
    fmt.Printf("a=%d, b=%d  (values unchanged - only pointer targets swapped)\n", a, b)
}
```

**Explanation**
`swapTargets` receives the ADDRESSES of `pa` and `pb` (i.e., `**int`). Inside,
`*pp1` and `*pp2` are the pointer variables themselves. Swapping them makes `pa`
point to `b` and `pb` point to `a`. The integers `a` and `b` are never touched —
only which pointer target which variable changes. This demonstrates the difference
between swapping values and swapping indirection levels.

---
---

## insane_01.go — Triple pointer indirection (***int)

```go
package main

import "fmt"

func main() {
    x  := 7
    p1 := &x
    p2 := &p1
    p3 := &p2

    fmt.Println("***p3 =", ***p3)

    ***p3 = 21 // modify x through three levels

    fmt.Println("x after change =", x)

    fmt.Printf("Address of x (3 ways): %p  %p  %p\n", &x, *p2, **p3)
}
```

**Explanation**
Dereference chain:
```
p3  → address of p2
*p3 → p2 (which holds the address of p1)
**p3 → p1 (which holds the address of x)   i.e. same as *p2
***p3 → x  (the int value)
```
All three address expressions resolve to `&x`:
- `&x` — direct
- `*p2` → value of p2 is p1 which is `&x`
- `**p3` → dereference p3 (gives p2), dereference again (gives p1 = `&x`)

Triple indirection is essentially never needed in real Go code, but tracing
through it cements the mental model for any depth of pointer chain.

---

## insane_02.go — Pointer arithmetic via `unsafe`

```go
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    nums := [5]int{10, 20, 30, 40, 50}
    p    := &nums[0]

    // Read element at index 2 via unsafe pointer arithmetic
    elem2 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 2*unsafe.Sizeof(nums[0])))
    fmt.Println("Element at index 2 via unsafe pointer:", *elem2)

    // Write to element at index 3
    elem3 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 3*unsafe.Sizeof(nums[0])))
    *elem3 = 999

    fmt.Println("Array after edit:", nums)
}
```

**Explanation**
The recipe for Go pointer arithmetic:
1. `unsafe.Pointer(p)` — strip the type, get a raw pointer
2. `uintptr(...)` — convert to an integer so you can add byte offsets
3. `+ N * unsafe.Sizeof(element)` — advance N elements
4. `unsafe.Pointer(...)` — convert back to raw pointer
5. `(*int)(...)` — re-type to the correct pointer type

**Why not store intermediate uintptr?** The GC can move objects between GC cycles.
If you store a `uintptr` in a variable and a GC runs before you convert it back to
`unsafe.Pointer`, the address is stale. The entire conversion must happen in a
single expression. This is why the Go spec says `uintptr` must not be stored.

---

## insane_03.go — Loop-pointer gotcha (pre Go 1.22) and the fix

```go
package main

import "fmt"

func main() {
    vals := []int{1, 2, 3, 4, 5}
    ptrs := make([]*int, 5)

    // PART A: Go < 1.22 bug — all pointers alias the loop variable
    // In Go 1.22+, the loop variable is per-iteration so this is fine.
    // In Go < 1.22, every ptrs[i] would hold the address of the SAME
    // variable `v`, which ends up holding the last value (5).
    // Result on old Go: 5 5 5 5 5
    for i, v := range vals {
        ptrs[i] = &v
    }
    fmt.Print("Part A:")
    for _, ptr := range ptrs {
        fmt.Print(" ", *ptr)
    }
    fmt.Println()

    // PART B: Fix that works on ALL Go versions — shadow v with a new variable
    for i, v := range vals {
        v := v         // new variable scoped to this iteration
        ptrs[i] = &v  // address of the shadowed copy, not the loop var
    }
    fmt.Print("Part B:")
    for _, ptr := range ptrs {
        fmt.Print(" ", *ptr)
    }
    fmt.Println()
}
```

**Explanation — the pre-Go 1.22 behaviour:**
In Go versions before 1.22, the `range` loop reuses a SINGLE loop variable `v`
on each iteration, overwriting it. Taking `&v` inside the loop always gives the
address of that one variable. After the loop, `v` holds 5 (the last element),
so all five pointers dereference to 5.

**The fix:** `v := v` creates a NEW variable in the current block scope, shadowing
the loop variable. `&v` now captures the address of this NEW variable, which is
unique per iteration.

**Go 1.22+:** The language spec was changed so each iteration gets its own copy of
the loop variable. Both Part A and Part B produce `1 2 3 4 5` on modern Go.
The fix is still good practice for portability.

---

## insane_04.go — Pointers mixed with structs

```go
package main

import "fmt"

type Point struct{ X, Y int }

func main() {
    pt  := Point{3, 7}
    pp  := &pt

    // auto-deref vs explicit dereference
    fmt.Printf("pp.X=%d, (*pp).Y=%d\n", pp.X, (*pp).Y)

    // modify via both forms
    pp.X   = 10   // auto-deref (idiomatic Go)
    (*pp).Y = 20  // explicit dereference-dot

    fmt.Printf("Modified: X=%d, Y=%d\n", pt.X, pt.Y)

    // pointer-to-pointer to struct
    pt2 := Point{100, 200}
    ppp := &pp

    *ppp = &pt2 // re-point pp to pt2

    fmt.Printf("Via ppp  : X=%d, Y=%d\n", (**ppp).X, (**ppp).Y)

    // address of an individual field
    fieldPtr := &pt2.X
    *fieldPtr = 999
    fmt.Println("pt2 after field ptr write:", pt2)
}
```

**Explanation**
Go automatically dereferences pointers to structs: `pp.X` is identical to
`(*pp).X`. The `->` operator from C/C++ does not exist in Go — the dot operator
is always used regardless of whether the receiver is a value or a pointer.

`&pt2.X` yields a `*int` pointing directly to the `X` field inside the `Point`
struct. Writes through it mutate the struct field in place.

`*ppp = &pt2` changes what `pp` points to. Subsequent reads via `pp` or `*ppp`
now reach `pt2`.

---

## insane_05.go — Aliasing puzzle: worked solution

```go
package main

import "fmt"

func main() {
    a, b, c := 1, 2, 3
    p := &a
    q := p       // q aliases a (same address as p)

    *p = 10      // (1)
    p = &b       // (2)
    *p = 20      // (3)
    q = p        // (4)
    *q = 30      // (5)
    p = &c       // (6)
    *p = *q + *p // (7)

    fmt.Println("a =", a) // 10
    fmt.Println("b =", b) // 30
    fmt.Println("c =", c) // 33
}
```

**Correct output:**
```
a = 10
b = 30
c = 33
```

**Step-by-step trace:**

| Step | Operation        | What changes                       | a  | b  | c  |
|------|------------------|-------------------------------------|----|----|----|
| init | `a,b,c := 1,2,3` | three separate ints                 | 1  | 2  | 3  |
| init | `p := &a`        | p holds address of a                | 1  | 2  | 3  |
| init | `q := p`         | q also holds address of a (aliases) | 1  | 2  | 3  |
| (1)  | `*p = 10`        | p→a, so a=10                        | 10 | 2  | 3  |
| (2)  | `p = &b`         | p now holds address of b            | 10 | 2  | 3  |
| (3)  | `*p = 20`        | p→b, so b=20                        | 10 | 20 | 3  |
| (4)  | `q = p`          | q now holds address of b            | 10 | 20 | 3  |
| (5)  | `*q = 30`        | q→b, so b=30                        | 10 | 30 | 3  |
| (6)  | `p = &c`         | p now holds address of c            | 10 | 30 | 3  |
| (7)  | `*p = *q + *p`   | RHS: *q=30, *p=3 → c=33            | 10 | 30 | 33 |

**Key insight:** At step (2), `p` is re-assigned to `&b`. But `q` still holds
the old address (`&a`). The two pointers are now independent. At step (4), `q`
is explicitly updated to match `p` (both now point to `b`). Step (7) evaluates
the RHS fully before the assignment: `*q` is 30 (b), `*p` is 3 (c), sum is 33,
then `c` is written.

---

*End of SOLUTIONS.md*
