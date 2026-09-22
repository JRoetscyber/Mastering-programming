/*
 * FILE: easy_04.cpp
 * TOPIC: Control Flow & Bitwise - Day 2
 * DIFFICULTY: Easy
 *
 * TASK: Bitmasking with if-statements
 * -----------------------------------
 * Bitwise AND (`&`) is used to check if a specific bit is set.
 * This is incredibly common in systems programming for "flags".
 *
 * 1. Declare three constants:
 *      const int EXECUTE = 1; // Binary: 001
 *      const int WRITE   = 2; // Binary: 010
 *      const int READ    = 4; // Binary: 100
 *
 * 2. A file has permissions set to 5 (Binary: 101).
 *      int myPermission = 5;
 *
 * 3. Write three separate `if` statements using the bitwise AND operator (`&`) 
 *    to check if `myPermission` has READ, WRITE, and EXECUTE.
 *      - If (myPermission & READ) is true, print "Has READ permission"
 *      - etc.
 *
 * EXPECTED OUTPUT:
 *   Has READ permission
 *   Has EXECUTE permission
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.


    return 0;
}
