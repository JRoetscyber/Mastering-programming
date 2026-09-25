/*
 * FILE: easy_03.cpp
 * TOPIC: Control Flow Deep Dive - Day 2
 * DIFFICULTY: Easy
 *
 * TASK: while vs do-while Loop Guarantee
 * ----------------------------------------
 * A `while` loop checks the condition BEFORE running the body (may run 0 times).
 * A `do while` loop runs the body FIRST, then checks the condition (guaranteed >= 1 run).
 *
 * 1. Declare: int count = 10;
 * 2. Write a `while (count < 5)` loop that prints "While running: " << count.
 * 3. Write a `do while (count < 5)` loop that prints "Do-While running: " << count.
 * 4. Observe the output: which loop executed and which one was skipped?
 * 5. Print a concluding message explaining why the do-while executed even though
 *    10 is NOT less than 5.
 *
 * EXPECTED OUTPUT:
 *   Do-While running: 10
 *   Explanation: do-while executes the body at least once before checking the condition.
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.
    int count = 10;

    while (count < 5) {
        std::cout << "While running : " << count << "\n";
    };

    do {
        std::cout << "Do-While running: " << count << "\n";
    } while (count < 5);

    return 0;
}
