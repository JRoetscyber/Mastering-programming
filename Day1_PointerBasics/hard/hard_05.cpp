/*
 * FILE: hard_05.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Hard
 *
 * TASK: Pointer arithmetic to find the maximum element in an array
 * ----------------------------------------------------------------
 * Use ONLY pointer arithmetic (no array indexing []) to find the
 * maximum value in an integer array.
 *
 * 1. Declare: int data[] = {15, 3, 72, 8, 44, 61, 5};
 *             const int SIZE = 7;
 * 2. Declare an int pointer `p` starting at data (first element).
 * 3. Declare an int pointer `maxPtr` also starting at data.
 * 4. Loop through the array using only pointer arithmetic:
 *      - If the current element (*p) is greater than *maxPtr, update maxPtr.
 *      - Increment p each iteration.
 * 5. Print the maximum value and its position (index) in the array.
 *    (Calculate index from pointer subtraction: maxPtr - data)
 *
 * EXPECTED OUTPUT:
 *   Maximum value : 72
 *   Found at index: 2
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.
    int data[] = { 15, 3, 72, 8, 44, 61, 5 };
    int* max_ptr = data;

    for (int* p = data; p < end(data); p++) {
        if (*p > *max_ptr) {
            max_ptr = p;
        };
    };

    int index = max_ptr - data;
    cout << "Maximum value :" << *max_ptr << "\n";
    cout << "Found at index:" << index << "\n";

    return 0;
}
