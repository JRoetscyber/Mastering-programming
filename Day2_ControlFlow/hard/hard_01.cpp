/*
 * FILE: hard_01.cpp
 * TOPIC: Control Flow Deep Dive - Day 2
 * DIFFICULTY: Hard
 *
 * TASK: Real-World Use of goto: Breaking Out of Deeply Nested Loops
 * ------------------------------------------------------------------
 * In C++, `break;` only breaks out of the INNERMOST loop.
 * If you have 3 nested loops (e.g. searching a 3D volume or 2D matrix),
 * a simple `break;` leaves you stuck in the outer loops unless you create
 * messy boolean flags.
 *
 * `goto` provides an instant exit label out of multi-level loops!
 *
 * SCENARIO:
 * You have a 3x3 grid:
 *   int grid[3][3] = {
 *       { 11, 12, 13 },
 *       { 21, 99, 23 },
 *       { 31, 32, 33 }
 *   };
 *
 * TASK:
 * 1. Search for the target value `99` using nested `for` loops (row and col).
 * 2. As you check each element, print: "Inspecting [row][col]: <val>"
 * 3. When `99` is found:
 *    - Jump directly to a label `found:` using `goto found;`
 * 4. At `found:`, print: "Found target 99 at [row][col]! Exited immediately."
 * 5. If loop finishes without finding target, print "Target not found."
 *
 * EXPECTED OUTPUT:
 *   Inspecting [0][0]: 11
 *   Inspecting [0][1]: 12
 *   Inspecting [0][2]: 13
 *   Inspecting [1][0]: 21
 *   Inspecting [1][1]: 99
 *   Found target 99 at [1][1]! Exited immediately.
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.



    return 0;
}
