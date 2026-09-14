# SOLUTIONS — Day 1: Pointer Basics (`&` and `*`)

> **Check your work only AFTER you have attempted each exercise yourself.**
> Solutions are labelled by filename. Scroll to the one you need.

---

## easy_01.cpp — Declare and use a pointer to int

```cpp
#include <iostream>

int main() {
    int  score = 95;
    int* p     = &score;

    std::cout << "Value of score via pointer: " << *p       << "\n";
    std::cout << "Address stored in p: "         << (void*)p << "\n";

    return 0;
}
```

**Explanation**  
`int* p = &score;` stores the address of `score` inside `p`.  
`*p` dereferences the pointer — it reads the `int` value at that address (95).  
We cast `p` to `void*` before printing to guarantee `cout` shows a hex address
(without the cast some implementations might behave unexpectedly).

---

## easy_02.cpp — Address-of operator on multiple types

```cpp
#include <iostream>

int main() {
    char   letter = 'A';
    int    count  = 7;
    double price  = 3.99;

    std::cout << "Address of letter : " << static_cast<void*>(&letter) << "\n";
    std::cout << "Address of count  : " << &count                      << "\n";
    std::cout << "Address of price  : " << &price                      << "\n";

    return 0;
}
```

**Explanation**  
`&` yields the address of any object. The cast to `void*` for `char` is
necessary because `operator<<` has a special overload for `char*` that treats
it as a C-string — without the cast you would see the character, not the address.
`int*` and `double*` print as hex addresses by default.

---

## easy_03.cpp — Pointer to double, sizeof a pointer

```cpp
#include <iostream>

int main() {
    double  temperature = 36.6;
    double* pTemp       = &temperature;

    std::cout << "Direct value   : " << temperature  << "\n";
    std::cout << "Via pointer    : " << *pTemp        << "\n";
    std::cout << "Size of pointer: " << sizeof(pTemp) << "\n";

    return 0;
}
```

**Explanation**  
`sizeof(pTemp)` gives the size of the POINTER variable, NOT the size of a
`double`. On a 64-bit build, all pointers are 8 bytes regardless of what type
they point to.

---

## easy_04.cpp — Pointer to char, address comparison

```cpp
#include <iostream>

int main() {
    char  grade  = 'B';
    char* pGrade = &grade;

    std::cout << "Grade via pointer: " << *pGrade << "\n";
    std::cout << "Address check    : "
              << (pGrade == &grade ? "Match" : "No match") << "\n";

    return 0;
}
```

**Explanation**  
`pGrade == &grade` compares two pointer VALUES (both hold the same address, so
the condition is true). This is pointer equality, not value equality.

---

## easy_05.cpp — Re-pointing a pointer

```cpp
#include <iostream>

int main() {
    int a = 10, b = 20;
    int* p = &a;

    std::cout << "p points to    : " << *p << "\n";

    p = &b;   // p now holds the address of b instead

    std::cout << "p now points to: " << *p << "\n";

    return 0;
}
```

**Explanation**  
Re-assigning `p = &b` changes what address `p` holds. It does NOT change
`a` or `b`. Think of `p` as a signpost: moving the signpost does not move the
destination.

---
---

## medium_01.cpp — Modify a variable through its pointer

```cpp
#include <iostream>

int main() {
    int  lives = 3;
    int* p     = &lives;

    std::cout << "lives = " << lives << "\n";

    (*p)--;                // subtract 1 via pointer
    std::cout << "lives = " << lives << "\n";

    *p = 0;               // set to 0 via pointer
    std::cout << "lives = " << lives << "\n";

    return 0;
}
```

**Explanation**  
`(*p)--` dereferences `p` to get the `int`, then decrements it. The parentheses
matter: `*p--` would decrement the pointer address itself (undefined here since
`p` was not allocated as part of an array).

---

## medium_02.cpp — Pointer to a pointer (double indirection)

```cpp
#include <iostream>

int main() {
    int   x  = 42;
    int*  p  = &x;
    int** pp = &p;

    std::cout << "x = "   << x    << "  |  "
              << "*p = "  << *p   << "  |  "
              << "**pp = " << **pp << "\n";

    std::cout << "&x = "  << &x   << "  |  "
              << "*pp = " << *pp  << "\n";  // *pp == p == &x

    return 0;
}
```

**Explanation**  
Dereference layers:  
- `pp` → address of `p`  
- `*pp` → the value stored in `p`, which is `&x` (the address of `x`)  
- `**pp` → the value at `&x`, which is `x` itself (42)

---

## medium_03.cpp — Swap two variables using pointers

```cpp
#include <iostream>

int main() {
    int a = 100, b = 200;
    int* pa = &a;
    int* pb = &b;

    std::cout << "Before: a = " << a << ", b = " << b << "\n";

    int temp = *pa;   // save a's value
    *pa = *pb;        // write b's value into a
    *pb = temp;       // write saved value into b

    std::cout << "After : a = " << a << ", b = " << b << "\n";

    return 0;
}
```

**Explanation**  
The swap uses a temporary `int` (not a pointer — we only need to hold a value
temporarily). Writing `*pa = *pb` copies the INT VALUE at `pb`'s address into
the memory at `pa`'s address. This is the foundational swap algorithm that
`std::swap` uses internally (via move semantics in modern C++).

---

## medium_04.cpp — Pointer to a specific array element

```cpp
#include <iostream>

int main() {
    int scores[] = {10, 20, 30, 40, 50};
    int* p       = &scores[2];   // point at the third element

    std::cout << "Pointed-to value: " << *p << "\n";

    *p = 99;   // modify through pointer

    std::cout << "Array after edit:";
    for (int i = 0; i < 5; ++i)
        std::cout << " " << scores[i];
    std::cout << "\n";

    return 0;
}
```

**Explanation**  
`&scores[2]` is equivalent to `scores + 2` — both give the address of the
third element. Modifying `*p` directly mutates `scores[2]` because `p` holds
its exact address.

---

## medium_05.cpp — Read and write through a pointer-to-pointer chain

```cpp
#include <iostream>

int main() {
    int   value = 5;
    int*  p1    = &value;
    int** p2    = &p1;

    std::cout << "Via p2 before: " << **p2 << "\n";

    **p2 *= 3;   // multiply value by 3 using only p2

    std::cout << "Via p2 after : " << **p2  << "\n";
    std::cout << "value direct : " << value << "\n";

    return 0;
}
```

**Explanation**  
`**p2` → `*p1` → `value`. The compound assignment `**p2 *= 3` is equivalent
to `value = value * 3` but travels through two layers of indirection. All three
names (`value`, `*p1`, `**p2`) refer to the same memory cell.

---
---

## hard_01.cpp — Walk an array with pointer arithmetic only

```cpp
#include <iostream>

int main() {
    int nums[] = {2, 4, 6, 8, 10};
    int* p     = nums;           // same as &nums[0]
    int* start = p;

    std::cout << "Elements:";
    for (int i = 0; i < 5; ++i, ++p)
        std::cout << " " << *p;
    std::cout << "\n";

    // p now points one-past-the-end; move back to last element
    int* last = p - 1;
    std::cout << "Element span: " << (last - start) << "\n";

    return 0;
}
```

**Explanation**  
Pointer arithmetic on a typed pointer moves in units of the pointed-to type.
`++p` on an `int*` advances the address by `sizeof(int)` (typically 4 bytes),
so it lands on the next element. `last - start` performs pointer subtraction,
which yields the number of ELEMENTS between the two pointers (not bytes).

---

## hard_02.cpp — Null pointer guard

```cpp
#include <iostream>

int main() {
    int* p = nullptr;

    if (p == nullptr)
        std::cout << "Pointer is null - cannot dereference safely.\n";
    else
        std::cout << "Pointer is valid. Value = " << *p << "\n";

    int x = 77;
    p = &x;

    if (p == nullptr)
        std::cout << "Pointer is null - cannot dereference safely.\n";
    else
        std::cout << "Pointer is valid. Value = " << *p << "\n";

    return 0;
}
```

**Explanation**  
`nullptr` (C++11+) is the null pointer constant. An uninitialised pointer
holds a GARBAGE address — dereferencing it is undefined behaviour that may not
crash immediately, making it far more dangerous than a null pointer (which at
least crashes predictably). Always initialise to `nullptr`.

**Bonus answer:**  
An uninitialised pointer holds a random garbage address. If you accidentally
dereference it, you might read or corrupt some random memory location silently
instead of crashing — the bug becomes almost impossible to track down.
`nullptr` ensures any accidental dereference produces a predictable crash.

---

## hard_03.cpp — Pass pointers into functions to modify originals

```cpp
#include <iostream>

void doubleIt(int* ptr) {
    *ptr *= 2;
}

void addFive(int* ptr) {
    *ptr += 5;
}

int main() {
    int n = 10;

    doubleIt(&n);
    std::cout << "After doubleIt: " << n << "\n";

    addFive(&n);
    std::cout << "After addFive : " << n << "\n";

    doubleIt(&n);
    std::cout << "After doubleIt: " << n << "\n";

    return 0;
}
```

**Explanation**  
By default C++ is pass-by-value: functions receive a COPY of the argument.
Passing `&n` gives the function the ADDRESS of `n`. Inside, `*ptr` dereferences
that address to reach the original variable. Any write through `*ptr` persists
after the function returns — this is pass-by-pointer (the C way to achieve what
references do in modern C++).

---

## hard_04.cpp — const pointer vs pointer to const

```cpp
#include <iostream>

int main() {
    int a = 1, b = 2;

    // Case A: constant pointer — p cannot be re-assigned, *p CAN change
    int* const pA = &a;
    *pA = 10;
    std::cout << "Case A (*pA after write)   : " << *pA << "\n";

    // Case B: pointer to const — p CAN be re-assigned, *p cannot change
    const int* pB = &a;
    pB = &b;          // re-point to b — legal
    // *pB = 99;      // <- COMPILE ERROR — *pB is read-only
    std::cout << "Case B (*pB after re-point): " << *pB << "\n";

    // Case C: const pointer to const — nothing can change
    const int* const pC = &b;
    std::cout << "Case C (*pC)              : " << *pC << "\n";

    return 0;
}
```

**Explanation**  
Read the declaration RIGHT-TO-LEFT from the variable name:  
- `int* const pA` → `pA` is a `const` pointer to `int` (the pointer is fixed)  
- `const int* pB` → `pB` is a pointer to `const int` (the pointed-to int is fixed)  
- `const int* const pC` → `pC` is a `const` pointer to `const int` (both fixed)

---

## hard_05.cpp — Find maximum using pointer arithmetic only

```cpp
#include <iostream>

int main() {
    int data[]   = {15, 3, 72, 8, 44, 61, 5};
    const int SIZE = 7;

    int* p      = data;
    int* maxPtr = data;

    for (int i = 0; i < SIZE; ++i, ++p) {
        if (*p > *maxPtr)
            maxPtr = p;
    }

    std::cout << "Maximum value : " << *maxPtr          << "\n";
    std::cout << "Found at index: " << (maxPtr - data)  << "\n";

    return 0;
}
```

**Explanation**  
`maxPtr` tracks the address of the current maximum element. Pointer subtraction
`maxPtr - data` yields the zero-based index (element offset from the start of
the array). This technique is the foundation of many standard algorithms like
`std::max_element`, which returns an iterator (effectively a pointer) rather
than a value.

---
---

## insane_01.cpp — Triple pointer indirection (int***)

```cpp
#include <iostream>

int main() {
    int    x  = 7;
    int*   p1 = &x;
    int**  p2 = &p1;
    int*** p3 = &p2;

    std::cout << "***p3 = " << ***p3 << "\n";

    ***p3 = 21;   // modify x through three levels of indirection

    std::cout << "x after change = " << x << "\n";

    // Four ways to express the address of x:
    std::cout << "Address of x (4 ways): "
              << &x          << "  "
              << *p2         << "  "   // *p2 == p1 == &x
              << **p3        << "  "   // **p3 == p1 == &x
              << &(***p3)    << "\n";  // address of what ***p3 denotes == &x

    return 0;
}
```

**Explanation**  
Each dereference level peels away one pointer:  
`p3` → address of p2  
`*p3` → p2 (address of p1)  
`**p3` → p1 (address of x)  
`***p3` → x (the int value)  

`&(***p3)` re-applies the address-of operator to `x`, so it gives `&x` again.
Triple indirection is rare in real code (usually a sign of design smell), but
understanding it cements your mental model of the dereference chain.

---

## insane_02.cpp — Pointer arithmetic across a 2D array

```cpp
#include <iostream>

int main() {
    int grid[2][3] = { {1, 2, 3}, {4, 5, 6} };
    const int ROWS = 2, COLS = 3;

    int* p = (int*)grid;   // flat view of the 2D array

    // (a) Print all 6 values in row-major order
    std::cout << "Flat traversal  :";
    for (int i = 0; i < ROWS * COLS; ++i)
        std::cout << " " << *(p + i);
    std::cout << "\n";

    // (b) grid[1][2] via formula
    std::cout << "grid[1][2] via p: " << *(p + 1*COLS + 2) << "\n";

    // (c) Change grid[0][1] to 99
    *(p + 0*COLS + 1) = 99;

    // (d) Print full grid with normal indexing
    std::cout << "Grid after edit :\n";
    for (int r = 0; r < ROWS; ++r) {
        for (int c = 0; c < COLS; ++c)
            std::cout << grid[r][c] << " ";
        std::cout << "\n";
    }

    return 0;
}
```

**Explanation**  
A 2D array `int grid[R][C]` is stored contiguously as `R*C` ints in memory.
The formula for element `[row][col]` as a flat offset is `row * COLS + col`.
The cast `(int*)grid` tells the compiler to treat the 2D array as a flat `int*`.
This is well-defined for standard-layout arrays in C++.

---

## insane_03.cpp — Function returning a pointer to a static local

```cpp
#include <iostream>

int* getCounter() {
    static int count = 0;   // persists between calls
    ++count;
    return &count;
}

int main() {
    int* p;

    p = getCounter();
    std::cout << "Call 1: " << *p << "\n";

    p = getCounter();
    std::cout << "Call 2: " << *p << "\n";

    p = getCounter();
    std::cout << "Call 3: " << *p << "\n";

    *p = 100;   // set counter to 100 via pointer
    std::cout << "Set to 100 via pointer.\n";

    p = getCounter();
    std::cout << "Call 4: " << *p << "\n";   // 101 (incremented from 100)

    return 0;
}
```

**Explanation**  
`static` local variables are allocated in the data segment (not the stack), so
their address remains valid for the lifetime of the program. Returning `&count`
is therefore safe — unlike returning the address of a non-static local.
Writing `*p = 100` mutates the hidden static `count` directly.

---

## insane_04.cpp — Pointers mixed with structs

```cpp
#include <iostream>

struct Point { int x; int y; };

int main() {
    Point  pt   = {3, 7};
    Point* pp   = &pt;

    // (c) print via arrow operator
    std::cout << "Original : x=" << pp->x << ", y=" << pp->y << "\n";

    // (d) modify using both syntaxes
    pp->x    = 10;    // arrow syntax
    (*pp).y  = 20;    // dereference-dot syntax (equivalent)

    std::cout << "Modified : x=" << pt.x << ", y=" << pt.y << "\n";

    // (e)-(h) pointer-to-pointer
    Point  pt2  = {100, 200};
    Point** ppp = &pp;

    *ppp = &pt2;    // re-point pp through ppp

    std::cout << "Via ppp  : x=" << (**ppp).x << ", y=" << (**ppp).y << "\n";

    return 0;
}
```

**Explanation**  
`pp->x` is syntactic sugar for `(*pp).x` — both dereference `pp` and then
access member `x`. The arrow `->` is the standard way to access members through
a pointer and is always preferred for readability.  
`*ppp = &pt2` writes a new address into `pp` (it changes what `pp` points to),
so subsequent uses of `pp` or `*ppp` reach `pt2`.

---

## insane_05.cpp — Predict the output: aliasing puzzle

```cpp
#include <iostream>

int main() {
    int a = 1, b = 2, c = 3;
    int* p = &a;
    int* q = p;          // q aliases a

    *p = 10;             // (1)  a=10, b=2,  c=3
    p  = &b;             // (2)  p now points to b
    *p = 20;             // (3)  a=10, b=20, c=3
    q  = p;              // (4)  q now aliases b (same as p)
    *q = 30;             // (5)  a=10, b=30, c=3
    p  = &c;             // (6)  p now points to c
    *p = *q + *p;        // (7)  RHS: *q=30, *p=3 -> c = 30+3 = 33

    std::cout << "a=" << a << "\n";   // 10
    std::cout << "b=" << b << "\n";   // 30
    std::cout << "c=" << c << "\n";   // 33

    return 0;
}
```

**Correct output:**
```
a=10
b=30
c=33
```

**Step-by-step explanation:**

| Step | Operation       | Effect                          | a  | b  | c  |
|------|-----------------|----------------------------------|----|----|-----|
| (1)  | `*p = 10`       | `p` → `a`, so a=10              | 10 | 2  | 3   |
| (2)  | `p = &b`        | `p` now holds address of `b`    | 10 | 2  | 3   |
| (3)  | `*p = 20`       | `p` → `b`, so b=20              | 10 | 20 | 3   |
| (4)  | `q = p`         | `q` now holds address of `b`    | 10 | 20 | 3   |
| (5)  | `*q = 30`       | `q` → `b`, so b=30              | 10 | 30 | 3   |
| (6)  | `p = &c`        | `p` now holds address of `c`    | 10 | 30 | 3   |
| (7)  | `*p = *q + *p`  | RHS evaluates first: 30 + 3 = 33; c=33 | 10 | 30 | 33 |

**Key insight:** Re-assigning a pointer (`p = &b`) does NOT affect other
pointers that previously held the same address. At step (2), `q` still points
to `a` even though `p` moved. At step (4), `q` is explicitly updated to `p`.

---

*End of SOLUTIONS.md*
