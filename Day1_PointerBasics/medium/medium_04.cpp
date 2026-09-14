/*
 * FILE: medium_04.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Medium
 *
 * TASK: Pointer to a specific element inside an array
 * ----------------------------------------------------
 * 1. Declare an int array:  int scores[] = {10, 20, 30, 40, 50};
 * 2. Declare an int pointer `p` that points to the THIRD element (index 2).
 * 3. Print the value at `p` (should be 30).
 * 4. Use `p` to change the third element to 99.
 * 5. Print the ENTIRE array using a regular for-loop to confirm the change.
 *
 * EXPECTED OUTPUT:
 *   Pointed-to value: 30
 *   Array after edit: 10 20 99 40 50
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.
    int scores[] = { 10, 20, 30, 40, 50 };

    int* p = &scores[2];

    cout << *p << "\n";

    *p = 99;

    for (int num : scores) {
        cout << num << "\n";
    }

    return 0;
}
