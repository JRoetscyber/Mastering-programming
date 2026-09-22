/*
 * FILE: easy_01.cpp
 * TOPIC: Control Flow Deep Dive - Day 2
 * DIFFICULTY: Easy
 *
 * TASK: Short-Circuit Evaluation & Guard Clause with if/else
 * -----------------------------------------------------------
 * In C++, logical AND (&&) stops immediately if the left condition is FALSE.
 * Logical OR (||) stops immediately if the left condition is TRUE.
 *
 * This prevents crashes! For example:
 *   if (ptr != nullptr && *ptr > 10)  <-- *ptr is NEVER dereferenced if ptr is null!
 *
 * 1. Declare an int pointer `ptr = nullptr`.
 * 2. Write an if-statement that checks:
 *      if `ptr` is not null AND the value at `ptr` is greater than 50:
 *          print "Valid and greater than 50"
 *      else:
 *          print "Safety triggered: pointer is null or value <= 50"
 * 3. Verify that running this does NOT crash your program.
 * 4. Next, assign `ptr = &val` where `int val = 85`.
 * 5. Run the same check again.
 *
 * EXPECTED OUTPUT:
 *   Safety triggered: pointer is null or value <= 50
 *   Valid and greater than 50
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.



    return 0;
}
