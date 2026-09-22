/*
 * FILE: hard_04.cpp
 * TOPIC: Control Flow & Bitwise - Day 2
 * DIFFICULTY: Hard
 *
 * TASK: Shift Operators and while loops (Hamming Weight)
 * ------------------------------------------------------
 * The right-shift operator (`>>`) moves all bits to the right.
 * e.g., `13 >> 1` changes `1101` (13) into `0110` (6).
 *
 * TASK:
 * Count how many '1' bits exist in the binary representation of a number.
 *
 * 1. Declare: unsigned int num = 29; // Binary: 11101 (Has four 1s)
 * 2. Declare: int count = 0;
 * 3. Create a `while` loop that runs as long as `num > 0`:
 *      a) If the lowest bit is 1 (`num & 1`), increment `count`.
 *      b) Shift `num` right by 1 bit (`num >>= 1`).
 * 4. Print the final count.
 *
 * EXPECTED OUTPUT:
 *   Number of set bits: 4
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.


    return 0;
}
