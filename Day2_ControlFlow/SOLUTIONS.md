# SOLUTIONS — Day 2: Control Flow Deep Dive

> **Check these ONLY after attempting each exercise yourself!**

---

## easy_01.cpp — Short-Circuit Evaluation & Guard Clause

```cpp
#include <iostream>
using namespace std;

int main() {
    int* ptr = nullptr;

    // First check: ptr is nullptr
    if (ptr != nullptr && *ptr > 50) {
        cout << "Valid and greater than 50\n";
    } else {
        cout << "Safety triggered: pointer is null or value <= 50\n";
    }

    // Second check: assign ptr to a valid address
    int val = 85;
    ptr = &val;

    if (ptr != nullptr && *ptr > 50) {
        cout << "Valid and greater than 50\n";
    } else {
        cout << "Safety triggered: pointer is null or value <= 50\n";
    }

    return 0;
}
```

**Explanation:**  
In `ptr != nullptr && *ptr > 50`, if `ptr` is `nullptr`, the left-hand expression is `false`. C++ **short-circuits** — it immediately stops evaluating and never evaluates `*ptr > 50`. This prevents a fatal crash (access violation / segfault).

---

## easy_02.cpp — switch Statement with Fallthrough

```cpp
#include <iostream>
using namespace std;

void checkSeverity(char severity) {
    cout << "Testing \'" << severity << "\': ";
    switch (severity) {
        case 'E':
        case 'e':
            cout << "Level: ERROR (Critical action required)\n";
            break;

        case 'W':
        case 'w':
            cout << "Level: WARNING (Check system logs)\n";
            break;

        case 'I':
        case 'i':
            cout << "Level: INFO (Normal status)\n";
            break;

        default:
            cout << "Level: UNKNOWN\n";
            break;
    }
}

int main() {
    checkSeverity('W');
    checkSeverity('e');
    checkSeverity('X');
    return 0;
}
```

**Explanation:**  
Omitting `break;` between `case 'E':` and `case 'e':` lets execution fall through to the same code block. This allows multiple values to share the same handler without code duplication.

---

## easy_03.cpp — while vs do-while Loop Guarantee

```cpp
#include <iostream>
using namespace std;

int main() {
    int count = 10;

    // while loop: checks condition BEFORE execution
    while (count < 5) {
        cout << "While running: " << count << "\n";
    }

    // do-while loop: executes body FIRST, then checks condition
    do {
        cout << "Do-While running: " << count << "\n";
    } while (count < 5);

    cout << "Explanation: do-while executes the body at least once before checking the condition.\n";

    return 0;
}
```

**Explanation:**  
`while` evaluates the condition first; since `10 < 5` is false, it executes 0 times. `do-while` runs the block once, prints `10`, and then checks `10 < 5`, terminating immediately.

---

## medium_01.cpp — Loop Control with break & continue

```cpp
#include <iostream>
using namespace std;

int main() {
    int readings[] = { 12, -1, 45, 0, 999, 28, -5, 60, 999, 10 };
    const int SIZE = 10;

    int validTotal = 0;
    int validCount = 0;

    for (int i = 0; i < SIZE; ++i) {
        int r = readings[i];

        if (r < 0) {
            cout << "Skipping corrupted sensor reading: " << r << "\n";
            continue; // Skip the rest of this iteration
        }

        if (r == 999) {
            cout << "CRITICAL FAILURE (999) detected! Aborting loop.\n";
            break; // Stop loop completely
        }

        cout << "Valid reading: " << r << "\n";
        validTotal += r;
        validCount++;
    }

    cout << "Processed valid count: " << validCount << "\n";
    cout << "Valid total: " << validTotal << "\n";
    if (validCount > 0) {
        cout << "Valid average: " << (validTotal / validCount) << "\n";
    }

    return 0;
}
```

**Explanation:**  
`continue` jumps straight to the loop increment (`++i`) for the next cycle. `break` terminates the entire loop immediately, ignoring all subsequent array elements.

---

## medium_02.cpp — State Machine using switch in a while loop

```cpp
#include <iostream>
using namespace std;

enum class State {
    IDLE,
    CONNECTING,
    CONNECTED,
    DISCONNECTED
};

int main() {
    State currentState = State::IDLE;
    int attempts = 0;

    while (currentState != State::DISCONNECTED) {
        switch (currentState) {
            case State::IDLE:
                cout << "System IDLE. Initiating connection...\n";
                currentState = State::CONNECTING;
                break;

            case State::CONNECTING:
                attempts++;
                cout << "Attempting connection (attempt " << attempts << ")...\n";
                if (attempts == 3) {
                    currentState = State::CONNECTED;
                }
                break;

            case State::CONNECTED:
                cout << "Connection established! Performing work...\n";
                cout << "Work finished. Disconnecting...\n";
                currentState = State::DISCONNECTED;
                break;

            case State::DISCONNECTED:
                break;
        }
    }

    cout << "Session closed gracefully.\n";
    return 0;
}
```

**Explanation:**  
State machines drive games, network protocols, and parser engines. An `enum class` combined with a `switch` statement makes each state transition explicit and clean.

---

## medium_03.cpp — Guard Clauses & Early Returns

```cpp
#include <iostream>
using namespace std;

bool canProcessTransaction(int accountBalance, int amount, bool isAccountActive, bool isFraudFlagged) {
    if (!isAccountActive) {
        cout << "Declined: Account inactive\n";
        return false;
    }
    if (isFraudFlagged) {
        cout << "Declined: Fraud alert triggered\n";
        return false;
    }
    if (amount <= 0) {
        cout << "Declined: Invalid transaction amount\n";
        return false;
    }
    if (amount > accountBalance) {
        cout << "Declined: Insufficient funds\n";
        return false;
    }

    cout << "Approved: Transaction processed!\n";
    return true;
}

int main() {
    cout << "Test A: ";
    canProcessTransaction(500, 100, false, false);

    cout << "Test B: ";
    canProcessTransaction(500, 800, true, false);

    cout << "Test C: ";
    canProcessTransaction(500, 250, true, false);

    return 0;
}
```

**Explanation:**  
Failing early with guard clauses eliminates deep indentation. The happy path stays unindented at the bottom.

---

## hard_01.cpp — Multi-Level Loop Break with goto

```cpp
#include <iostream>
using namespace std;

int main() {
    int grid[3][3] = {
        { 11, 12, 13 },
        { 21, 99, 23 },
        { 31, 32, 33 }
    };

    int target = 99;
    int foundRow = -1, foundCol = -1;

    for (int r = 0; r < 3; ++r) {
        for (int c = 0; c < 3; ++c) {
            cout << "Inspecting [" << r << "][" << c << "]: " << grid[r][c] << "\n";
            if (grid[r][c] == target) {
                foundRow = r;
                foundCol = c;
                goto found; // Escape both loops immediately!
            }
        }
    }

    cout << "Target not found.\n";
    return 0;

found:
    cout << "Found target 99 at [" << foundRow << "][" << foundCol << "]! Exited immediately.\n";
    return 0;
}
```

**Explanation:**  
In C++, `break` only exits the innermost loop. `goto` allows an instant jump out of arbitrarily deeply nested loops without introducing helper boolean flags or function wrappers.

---

## hard_02.cpp — Error Cleanup Pattern via goto (Linux Kernel Pattern)

```cpp
#include <iostream>
using namespace std;

bool initializeService(bool failStage1, bool failStage2, bool failStage3) {
    // Stage 1
    if (failStage1) {
        cout << "Failed Stage 1\n";
        goto cleanup_none;
    }
    cout << "Stage 1 Acquired (Network Buffer)\n";

    // Stage 2
    if (failStage2) {
        cout << "Failed Stage 2\n";
        goto cleanup_stage1;
    }
    cout << "Stage 2 Acquired (DB Connection)\n";

    // Stage 3
    if (failStage3) {
        cout << "Failed Stage 3\n";
        goto cleanup_stage2;
    }
    cout << "Stage 3 Acquired (User Auth)\n";

    cout << "Service Started Successfully!\n";
    return true;

// Unwind in reverse order of acquisition
cleanup_stage2:
    cout << "Releasing Stage 2 (DB Connection)\n";
cleanup_stage1:
    cout << "Releasing Stage 1 (Network Buffer)\n";
cleanup_none:
    cout << "Initialization aborted.\n";
    return false;
}

int main() {
    initializeService(false, true, false);
    return 0;
}
```

**Explanation:**  
In kernel programming and driver code without C++ exceptions/RAII, acquiring resources sequentially requires unwinding in exact reverse order if any intermediate step fails. The stacked `goto` labels execute fallthrough cleanups cleanly without duplicating cleanup code.

---

## hard_03.cpp — Bubble Sort Invariants & Early Termination

```cpp
#include <iostream>
using namespace std;

int main() {
    int arr[] = { 4, 1, 8, 3, 9, 2, 7 };
    const int N = 7;

    int passes = 0;

    for (int i = 0; i < N - 1; ++i) {
        passes++;
        bool swapped = false;

        // Invariant: after pass i, the last i elements are already in their final sorted positions!
        for (int j = 0; j < N - i - 1; ++j) {
            if (arr[j] > arr[j + 1]) {
                int temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
                swapped = true;
            }
        }

        // Optimization: if no two elements swapped, array is already sorted!
        if (!swapped) {
            break;
        }
    }

    cout << "Passes taken: " << passes << "\n";
    cout << "Sorted array:";
    for (int i = 0; i < N; ++i) {
        cout << " " << arr[i];
    }
    cout << "\n";

    return 0;
}
```

**Explanation:**  
1. **Loop Bound `N - i - 1`:** After each outer pass `i`, the largest remaining element bubbles up to the end. So the last `i` elements are guaranteed to be sorted. Checking past `N - i - 1` would cause `arr[j + 1]` to access either already sorted items or out-of-bounds memory.
2. **Early exit `!swapped`:** Prevents unnecessary $O(N^2)$ comparisons on partially or fully sorted data, achieving $O(N)$ best-case.

---

## easy_04.cpp — Bitmasking with if-statements

```cpp
#include <iostream>
using namespace std;

int main() {
    const int EXECUTE = 1; // 001
    const int WRITE   = 2; // 010
    const int READ    = 4; // 100

    int myPermission = 5; // 101 (READ and EXECUTE)

    // Bitwise AND (&) extracts the exact bit.
    // If the result is non-zero, the condition is true.
    if (myPermission & READ) {
        cout << "Has READ permission\n";
    }
    if (myPermission & WRITE) {
        cout << "Has WRITE permission\n";
    }
    if (myPermission & EXECUTE) {
        cout << "Has EXECUTE permission\n";
    }

    return 0;
}
```

---

## medium_04.cpp — Setting, Clearing, and Toggling Bits

```cpp
#include <iostream>
using namespace std;

int main() {
    unsigned char state = 0;
    const unsigned char ALARM = 8; // 00001000

    // SET bit using OR
    state |= ALARM;
    if (state & ALARM) {
        cout << "Alarm is ON\n";
    }

    // TOGGLE bit using XOR
    state ^= ALARM;
    if (!(state & ALARM)) {
        cout << "Alarm is OFF\n";
    }

    // SET again, then CLEAR using AND NOT
    state |= ALARM;
    state &= ~ALARM; // ~ALARM flips 00001000 to 11110111. ANDing leaves all other bits intact but kills the ALARM bit.
    if (!(state & ALARM)) {
        cout << "Alarm successfully cleared\n";
    }

    return 0;
}
```

---

## hard_04.cpp — Shift Operators and while loops (Hamming Weight)

```cpp
#include <iostream>
using namespace std;

int main() {
    unsigned int num = 29; // 11101
    int count = 0;

    while (num > 0) {
        if (num & 1) { // Check if the right-most bit is a 1
            count++;
        }
        num >>= 1; // Shift everything to the right by 1
    }

    cout << "Number of set bits: " << count << "\n";

    return 0;
}
```
