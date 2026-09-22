/*
 * FILE: medium_01.cpp
 * TOPIC: Control Flow Deep Dive - Day 2
 * DIFFICULTY: Medium
 *
 * TASK: Loop Control with break, continue, and nested loops
 * -----------------------------------------------------------
 * You are processing a 1D stream of sensor readings:
 *   int readings[] = { 12, -1, 45, 0, 999, 28, -5, 60, 999, 10 };
 *   const int SIZE = 10;
 *
 * Rules to implement in a single loop:
 * 1. If reading is NEGATIVE (e.g. -1, -5):
 *    - Print: "Skipping corrupted sensor reading: <val>"
 *    - Use `continue;` to immediately skip to the next iteration.
 *
 * 2. If reading is 999:
 *    - Print: "CRITICAL FAILURE (999) detected! Aborting loop."
 *    - Use `break;` to immediately stop processing further elements.
 *
 * 3. Otherwise:
 *    - Add the reading to a running `validTotal` sum and count valid readings.
 *    - Print: "Valid reading: <val>"
 *
 * 4. After the loop, print the total sum and the average of valid readings.
 *
 * EXPECTED OUTPUT:
 *   Valid reading: 12
 *   Skipping corrupted sensor reading: -1
 *   Valid reading: 45
 *   Valid reading: 0
 *   CRITICAL FAILURE (999) detected! Aborting loop.
 *   Processed valid count: 3
 *   Valid total: 57
 *   Valid average: 19
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.



    return 0;
}
