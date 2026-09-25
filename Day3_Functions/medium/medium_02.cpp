/*
 * FILE: medium_02.cpp
 * TOPIC: Functions Deep Dive - Day 3
 * DIFFICULTY: Medium
 *
 * TASK: Lambda Captures
 * ---------------------
 * Lambdas can "capture" variables from outside their scope!
 *   [var]  <- Capture by value (read-only copy)
 *   [&var] <- Capture by reference (can modify original)
 *
 * 1. Declare `int multiplier = 5;`.
 * 2. Use `std::transform` and a lambda that captures `multiplier` by value.
 * 3. The lambda should take an `int`, multiply it by `multiplier`, and return it.
 *
 * EXPECTED OUTPUT:
 *   Original: 1 2 3 4 5
 *   Multiplied: 5 10 15 20 25
 */

#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int main() {
    vector<int> numbers = { 1, 2, 3, 4, 5 };
    int multiplier = 5;

    // TODO: Use std::transform with a capturing lambda to multiply all numbers in place.
    // std::transform(numbers.begin(), numbers.end(), numbers.begin(), ...lambda...);


    return 0;
}
