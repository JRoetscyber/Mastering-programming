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
    int rows = 10, cols = 10;
    int low = 0, high = rows * cols - 1;
    int target = 99;

    int grid[10][10] = {
    {  1,  2,  3,  4,  5,  6,  7,  8,  9, 10 },
    { 11, 12, 13, 14, 15, 16, 17, 18, 19, 20 },
    { 21, 22, 23, 24, 25, 26, 27, 28, 29, 30 },
    { 31, 32, 33, 34, 35, 36, 37, 38, 39, 40 },
    { 41, 42, 43, 44, 45, 46, 47, 48, 49, 50 },
    { 51, 52, 53, 54, 55, 56, 57, 58, 59, 60 },
    { 61, 62, 63, 64, 65, 66, 67, 68, 99, 70 },
    { 71, 72, 73, 74, 75, 76, 77, 78, 79, 80 },
    { 81, 82, 83, 84, 85, 86, 87, 88, 89, 90 },
    { 91, 92, 93, 94, 95, 96, 97, 98, 69, 100 }
    };

    int foundRow = -1, foundCol = -1;

    for (int row = 0; row < rows; row++) {
        cout << "Binary searching row " << row << "...\n";

        int low = 0, high = cols - 1;

        while (low <= high) {
            int mid = low + (high - low) / 2;
            int val = grid[row][mid];

            cout << "  Inspecting [" << row << "][" << mid << "]: " << val << "\n";

            if (val == target) {
                foundRow = row;
                foundCol = mid;
                goto found;
            }
            if (val < target) {
                low = mid + 1;
            }
            else {
                high = mid - 1;
            }
        }
    }

    cout << "Target not found.\n";
    return 0;

found:
    std::cout << "Found " << target << " at [" << foundRow << "][" << foundCol << "]\n";
    return 0;
}
