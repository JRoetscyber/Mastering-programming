/*
 * FILE: easy_03.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Easy
 *
 * TASK: Pointer to a double - declare, point, and dereference
 * ------------------------------------------------------------
 * 1. Declare a double called `temperature` with the value 36.6.
 * 2. Declare a double pointer `pTemp` that points to `temperature`.
 * 3. Print the value of `temperature` DIRECTLY (using the variable itself).
 * 4. Print the value of `temperature` INDIRECTLY (using the pointer).
 * 5. Print the SIZE (in bytes) of the pointer itself using sizeof(pTemp).
 *
 * EXPECTED OUTPUT:
 *   Direct value   : 36.6
 *   Via pointer    : 36.6
 *   Size of pointer: 8   <- (will be 4 on 32-bit builds)
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.
    double temp = 36.6;
    double* pTemp = &temp;
    cout << "Direct value:" << temp << "\n";
    cout << "Via pointer:" << *pTemp << "\n";
    cout << "Size of pointer: " << sizeof(pTemp) << "\n";


    return 0;
}
