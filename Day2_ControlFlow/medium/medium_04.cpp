/*
 * FILE: medium_04.cpp
 * TOPIC: Control Flow & Bitwise - Day 2
 * DIFFICULTY: Medium
 *
 * TASK: Setting, Clearing, and Toggling Bits
 * ------------------------------------------
 *   - Bitwise OR (`|`) sets a bit to 1.
 *   - Bitwise XOR (`^`) toggles a bit (flips 1 to 0, or 0 to 1).
 *   - Bitwise AND with NOT (`& ~`) clears a bit to 0.
 *
 * 1. Declare: unsigned char state = 0; // All bits 00000000
 * 2. Define a flag: const unsigned char ALARM = 8; // Binary: 00001000
 *
 * 3. SET the ALARM bit on `state` using `|=`.
 * 4. Write an if-statement checking if ALARM is on. Print "Alarm is ON".
 *
 * 5. TOGGLE the ALARM bit using `^=`.
 * 6. Write an if-statement checking if ALARM is off (using `&`). Print "Alarm is OFF".
 *
 * 7. SET the ALARM bit again (`|=`).
 * 8. CLEAR the ALARM bit using `&= ~ALARM`.
 * 9. Verify it is off by printing "Alarm successfully cleared" if it's off.
 *
 * EXPECTED OUTPUT:
 *   Alarm is ON
 *   Alarm is OFF
 *   Alarm successfully cleared
 */

#include <iostream>
#include <bitset>

int main() {
    // TODO: Write your solution here.
    unsigned char state = 0;
    const unsigned char ALARM = 8;

    state |= ALARM;

    if (state & ALARM) {
        std::cout << "Alarm is ON" << "\n";
    }

    state ^= ALARM;

    if (!(state & ALARM)) {
        std::cout << "Alarm is OFF" << "\n";
    }

    state |= ALARM;

    state &= ~ALARM;

    if (!(state & ALARM)) {
        std::cout << "Alarm successfully cleared" << "\n";
    }

    return 0;
}
