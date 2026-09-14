/*
 * FILE: medium_02.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Medium
 *
 * TASK: Pointer to a pointer (double indirection)
 * -----------------------------------------------
 * 1. Declare an int `x = 42`.
 * 2. Declare an int pointer  `p  = &x`.
 * 3. Declare an int** pointer `pp = &p`  (a pointer to the pointer).
 * 4. Print the value of `x` in THREE different ways:
 *      a) directly: x
 *      b) single dereference: *p
 *      c) double dereference: **pp
 * 5. Print the address of `x` in TWO ways:
 *      a) &x
 *      b) *pp   (dereference pp once - you get the address stored in p)
 *
 * EXPECTED OUTPUT (addresses will differ):
 *   x = 42  |  *p = 42  |  **pp = 42
 *   &x = 0x...  |  *pp = 0x...  (both should be the same address)
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.
    int x = 42;
    cout << "x = " << x << "\n";
    cout << "&x = " << &x << "\n";

    int* p = &x;
    cout << "*p = " << *p << "\n";

    int** pp = &p;
    cout << "**pp = " << **pp << "\n";
    cout << "**pp = " << *pp << "\n";

    return 0;
} /* i just pointer at the adress didn't i made a pointer that returns the the adress */
