/*
 * FILE: hard_04.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Hard
 *
 * TASK: const pointer vs pointer to const
 * ----------------------------------------
 * There are THREE distinct pointer+const combinations. Understand each:
 *
 *   (A) int* const p        -- constant pointer to int
 *                              p cannot be re-assigned; *p CAN be changed.
 *   (B) const int* p        -- pointer to constant int
 *                              p CAN be re-assigned; *p cannot be changed.
 *   (C) const int* const p  -- constant pointer to constant int
 *                              neither p nor *p can change.
 *
 * 1. Declare:  int a = 1, b = 2;
 *
 * 2. Case A: int* const pA = &a;
 *      - Try to change *pA (change a's value) - this should COMPILE fine.
 *      - Do NOT attempt to re-assign pA itself (that would be a compile error).
 *      - Print *pA after the change.
 *
 * 3. Case B: const int* pB = &a;
 *      - Re-assign pB to point to b - this should COMPILE fine.
 *      - Do NOT attempt to modify *pB (that would be a compile error).
 *      - Print *pB after re-assignment.
 *
 * 4. Case C: const int* const pC = &b;
 *      - Just print *pC.
 *      - Do NOT attempt to change either pC or *pC.
 *
 * EXPECTED OUTPUT:
 *   Case A (*pA after write)  : 10
 *   Case B (*pB after re-point): 2
 *   Case C (*pC)              : 2
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.

    int a = 1, b = 2;

    int* const pA = &a;
    *pA = 10;
    cout << "(*pA after write) :" << *pA << "\n";

    const int* pB = &a;
    pB = &b;
    cout << "(*pB after re-point): " << *pB << "\n";

    const int* const pC = &b;
    cout << "(*pC) :" << *pC << "\n";


    return 0;
}
