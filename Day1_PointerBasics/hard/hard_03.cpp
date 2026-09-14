/*
 * FILE: hard_03.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Hard
 *
 * TASK: Pass a pointer into a function to modify the original variable
 * --------------------------------------------------------------------
 * You may NOT use references (&) in the function signature - only raw pointers.
 *
 * 1. Write a function:
 *      void doubleIt(int* ptr)
 *    that multiplies the value at `ptr` by 2.
 *
 * 2. Write a function:
 *      void addFive(int* ptr)
 *    that adds 5 to the value at `ptr`.
 *
 * 3. In main():
 *      a) Declare int n = 10.
 *      b) Call doubleIt(&n).   Print n.
 *      c) Call addFive(&n).    Print n.
 *      d) Call doubleIt(&n).   Print n.
 *
 * EXPECTED OUTPUT:
 *   After doubleIt: 20
 *   After addFive : 25
 *   After doubleIt: 50
 */

#include <iostream>
using namespace std;

// TODO: Write your function declarations and definitions here.

void doubleIt(int* ptr) {
    *ptr = *ptr * 2;
}

void addFive(int* ptr) {
    *ptr = *ptr + 5;
}

int main() {
    // TODO: Write your solution here.
    int n = 10;
    doubleIt(&n);
    cout << "After doubleIt:" << n << "\n";
    addFive(&n);
    cout << "After addFive :" << n << "\n";
    doubleIt(&n);
    cout << "After doubleIt:" << n << "\n";
    return 0;
}
