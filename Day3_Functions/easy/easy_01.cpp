/*
 * FILE: easy_01.cpp
 * TOPIC: Functions Deep Dive - Day 3
 * DIFFICULTY: Easy
 *
 * TASK: Pass by Value vs Pass by Reference
 * -----------------------------------------
 * 1. Write a function `void healByValue(int health)` that adds 50 to health.
 * 2. Write a function `void healByReference(int& health)` that adds 50 to health.
 * 3. In main, create `int playerHealth = 100;`.
 * 4. Call `healByValue(playerHealth)` and print the health. (Notice it doesn't change!)
 * 5. Call `healByReference(playerHealth)` and print the health. (Notice it changes!)
 *
 * EXPECTED OUTPUT:
 *   Health after healByValue: 100
 *   Health after healByReference: 150
 */

#include <iostream>
using namespace std;

// TODO: Write healByValue here
void healByValue(int health) {
    health = health + 50;
}

// TODO: Write healByReference here
void healByReference(int& health) {
    health = health + 50;
}

int main() {
    // TODO: Write your solution here.
    int playerHealth = 100;

    healByValue(playerHealth);
    std::cout << "Health after healByValue: " << playerHealth << "\n";

    healByReference(playerHealth);
    std::cout << "Health after healByValue: " << playerHealth << "\n";

    return 0;
}
