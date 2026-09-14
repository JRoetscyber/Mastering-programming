/*
 * FILE: medium_05.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Medium
 *
 * TASK: Read and write through a pointer-to-pointer chain
 * -------------------------------------------------------
 * 1. Declare:  int value = 5;
 * 2. Declare:  int*  p1  = &value;
 * 3. Declare:  int** p2  = &p1;
 * 4. Using ONLY `p2` (no direct use of `value` or `p1` after this point),
 *    do the following:
 *      a) Print the current value (double-dereference p2).
 *      b) Multiply `value` by 3 through the pointer chain.
 *      c) Print the new value again through p2.
 * 5. Finally, print `value` directly to confirm it changed.
 *
 * EXPECTED OUTPUT:
 *   Via p2 before: 5
 *   Via p2 after : 15
 *   value direct : 15
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.

    int value = 5;
    int* p1 = &value;
    int** p2 = &p1;

    cout << **p2 << "\n";

    **p2 = **p2 * 3;

    cout << **p2 << "\n";
    return 0;
}
