/*
 * FILE: easy_04.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Easy
 *
 * TASK: Pointer to a char - print the character via dereferencing
 * ---------------------------------------------------------------
 * 1. Declare a char variable `grade` with the value 'B'.
 * 2. Declare a char pointer `pGrade` pointing to `grade`.
 * 3. Print the character stored in `grade` by DEREFERENCING the pointer.
 * 4. Print whether the pointer address matches the address of `grade`
 *    (print "Match" if pGrade == &grade, otherwise "No match").
 *
 * EXPECTED OUTPUT:
 *   Grade via pointer: B
 *   Address check    : Match
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.
    char grade = 'B';
    char* pGrade = &grade;
    cout << "Grade via pointer :" << *pGrade << "\n";
    if (&grade != pGrade) {
        cout << "Address check :" << "not a match" << "\n";
    } else {
        cout << "Address check : " << "match" << "\n";
    }


    return 0;
}
