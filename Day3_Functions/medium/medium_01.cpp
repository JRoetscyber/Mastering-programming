/*
 * FILE: medium_01.cpp
 * TOPIC: Functions Deep Dive - Day 3
 * DIFFICULTY: Medium
 *
 * TASK: Lambdas Basics & std::count_if
 * -------------------------------------
 * A lambda is an anonymous, inline function:  [] (int x) { return x > 5; }
 *
 * 1. You are given a vector of test scores.
 * 2. Use `std::count_if` along with a lambda function to count how many
 *    scores are greater than or equal to 80 (Passing grade).
 *
 * EXPECTED OUTPUT:
 *   Number of passing grades: 4
 */

#include <iostream>
#include <vector>
#include <algorithm> // For count_if
using namespace std;

int main() {
    vector<int> scores = { 45, 88, 92, 70, 80, 55, 100, 32 };

    // TODO: Use std::count_if and a lambda to count scores >= 80
    int passingCount = std::count_if([&scores](int i) {return i > 80:})


    return 0;
}
