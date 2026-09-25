/*
 * FILE: hard_03.cpp
 * TOPIC: Control Flow Deep Dive - Day 2
 * DIFFICULTY: Hard
 *
 * TASK: Complex Nested Loop Invariant & Off-by-One Debugging
 * -----------------------------------------------------------
 * A common source of security vulnerabilities and bugs in C/C++ is the
 * "off-by-one" error in loops, or failing to maintain loop invariants.
 *
 * TASK:
 * 1. Declare a 1D array of 7 integers:
 *      int arr[] = { 4, 1, 8, 3, 9, 2, 7 };
 *      const int N = 7;
 *
 * 2. Implement Bubble Sort to sort `arr` in ascending order using TWO nested `for` loops.
 *    Pay close attention to your loop bounds:
 *      - Outer loop `i` from 0 to N - 1
 *      - Inner loop `j` from 0 to (N - i - 1)   <-- Why subtract i and 1?
 *
 * 3. In the inner loop:
 *      if arr[j] > arr[j + 1], swap them!
 *
 * 4. Add an optimization flag: `bool swapped = false;`
 *    If no swaps occurred during an entire pass of the inner loop, `break;` early!
 *
 * 5. Print the sorted array and count how many total passes the outer loop took.
 *
 * EXPECTED OUTPUT:
 *   Passes taken: 5   (or fewer depending on data)
 *   Sorted array: 1 2 3 4 7 8 9
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.
    int arr[] = { 4, 1, 8, 3, 9, 2, 7 };
    const int N = 7;
    bool swapped;
    int passes = 0;

    for (int i = 0; i < N - 1; i++) {
        swapped = false;
        passes++;
        for (int j = 0; j < N - i - 1; j++) {
            if (arr[j] > arr[j + 1]) {
                std::swap(arr[j], arr[j + 1]);
                swapped = true;
            }
        }
        if (!swapped) {
            break;
        }
        
    }
    std::cout << "Passes taken: " << passes << "\n";

    std::cout << "Sorted array: ";
    for (int num : arr) {
        std::cout << num << " ";
    }
    std::cout << "\n";

    return 0;
}
