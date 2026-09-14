/*
 * FILE: medium_03.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Medium
 *
 * TASK: Swap two variables using pointers (no std::swap)
 * ------------------------------------------------------
 * Write a swap WITHOUT using std::swap or references.
 * Use ONLY raw pointers.
 *
 * 1. Declare two ints: `a = 100`, `b = 200`.
 * 2. Print them BEFORE the swap.
 * 3. Perform the swap using pointers and a temporary variable:
 *      int* pa = &a;
 *      int* pb = &b;
 *      ... (your swap logic using *pa and *pb)
 * 4. Print `a` and `b` AFTER the swap.
 *
 * EXPECTED OUTPUT:
 *   Before: a = 100, b = 200
 *   After : a = 200, b = 100
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.
    int a = 100;
    int b = 200;

    int* pa = &a;
    int* pb = &b;

    cout << "Before: a = " << *pa << ", b = " << *pb << "\n";

    int temp = *pa;
    
    *pa = *pb;
    *pb = temp;

    cout << "After: a = " << *pa << ", b = " << *pb << "\n";
    return 0;
}
