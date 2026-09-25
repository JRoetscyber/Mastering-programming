/*
 * FILE: hard_02.cpp
 * TOPIC: Functions Deep Dive - Day 3
 * DIFFICULTY: Hard
 *
 * TASK: Templates (Compile-Time Static Polymorphism)
 * ---------------------------------------------------
 * Function overloading (easy_02) is great, but writing the exact same logic
 * for ints, floats, and strings is tedious. Templates allow you to write the
 * logic ONCE, and the compiler generates the overloads for you automatically!
 *
 * 1. Write a template function `getMax`:
 *      template <typename T>
 *      T getMax(T a, T b) { return (a > b) ? a : b; }
 *
 * 2. In main, test it with:
 *      - Two ints (5, 10)
 *      - Two floats (3.14, 2.71)
 *      - Two strings ("Apple", "Zebra")
 *
 * EXPECTED OUTPUT:
 *   Max int: 10
 *   Max float: 3.14
 *   Max string: Zebra
 */

#include <iostream>
#include <string>
using namespace std;

// TODO: Write your template function here


int main() {
    // TODO: Write your solution here.

    return 0;
}
