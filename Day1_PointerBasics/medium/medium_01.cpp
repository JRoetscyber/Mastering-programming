/*
 * FILE: medium_01.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Medium
 *
 * TASK: Modify a variable's value THROUGH its pointer
 * ---------------------------------------------------
 * 1. Declare an int `lives = 3`.
 * 2. Print the original value of `lives`.
 * 3. Declare an int pointer `p` pointing to `lives`.
 * 4. Use the pointer to subtract 1 from `lives` (write through `p`).
 * 5. Print the new value of `lives` (read it directly, NOT through the pointer).
 * 6. Use the pointer again to set `lives` to 0.
 * 7. Print `lives` one more time.
 *
 * EXPECTED OUTPUT:
 *   lives = 3
 *   lives = 2
 *   lives = 0
 *
 * NOTE: The change made through the pointer must be visible when reading `lives`
 *       directly - this confirms they share the same memory location.
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.
    int lives = 3;
    cout << "lives = " << lives << "\n";

    int* p = &lives;
    *p = *p - 1;
    cout << "lives = " << lives << "\n";

    *p = 0;
    cout << "lives = " << lives << "\n";



    return 0;
}
