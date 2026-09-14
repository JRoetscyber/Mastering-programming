/*
 * FILE: hard_02.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Hard
 *
 * TASK: Null pointers and detecting an invalid pointer
 * ----------------------------------------------------
 * A null pointer holds the address 0 (nullptr in modern C++).
 * Dereferencing it is undefined behaviour - your program will crash.
 * This exercise teaches you to GUARD against that.
 *
 * 1. Declare an int pointer `p` and initialise it to nullptr.
 * 2. Write an if-statement that checks whether `p` is null.
 *    If null: print "Pointer is null - cannot dereference safely."
 *    If not null: print the value it points to.
 * 3. Now assign `p` to point to a local int variable `x = 77`.
 * 4. Run the same if-check again and print the value this time.
 *
 * EXPECTED OUTPUT:
 *   Pointer is null - cannot dereference safely.
 *   Pointer is valid. Value = 77
 *
 * BONUS QUESTION (think about it, no code needed):
 *   Why is it important to always initialise pointers to nullptr
 *   rather than leaving them uninitialised?
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.
    int* p;
    int x = 77;
    int num = 0;

    while (num < 2) {
        if (num == 0) {
            if (p == nullptr) {
                cout << "Pointer is null - cannot dereference safely." << "\n";
            };
        };
        if (num == 1) {
            p = &x;
            if (p != nullptr) {
                cout << "Pointer is valid. Value = " << *p << "\n";
            };
        };
        num++;
    };

    return 0;
}
