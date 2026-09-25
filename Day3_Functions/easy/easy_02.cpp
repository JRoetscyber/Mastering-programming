/*
 * FILE: easy_02.cpp
 * TOPIC: Functions Deep Dive - Day 3
 * DIFFICULTY: Easy
 *
 * TASK: Function Overloading (Static Polymorphism)
 * ------------------------------------------------
 * Function overloading allows multiple functions to share the exact same name,
 * as long as their parameter types or counts are different. The compiler figures
 * out which one to call at compile-time (Static Polymorphism).
 *
 * 1. Write `int multiply(int a, int b)`
 * 2. Write `double multiply(double a, double b)`
 * 3. Write `int multiply(int a, int b, int c)`
 *
 * 4. In main, test all three functions.
 *
 * EXPECTED OUTPUT:
 *   Int multiply (2x3): 6
 *   Double multiply (2.5x3.0): 7.5
 *   Int multiply (2x3x4): 24
 */

#include <iostream>
using namespace std;

// TODO: Write your 3 overloaded multiply functions here
int multiply(int a, int b) {
    int answer = a * b;
    std::cout << "Int multiply (" << a << "x" << b << "): " << answer << "\n";
    return answer;
}
double multiply(double a, double b) {
    double answer = a * b;
    std::cout << "Double multiply (" << a << "x" << b << "): " << answer << "\n";
    return answer;
}
int multiply(int a, int b, int c) {
    int answer = a * b * c;
    std::cout << "Int multiply (" << a << "x" << b << "x" << c << "): " << answer << "\n";
    return answer;
}


int main() {
    // TODO: Write your solution here.
    multiply(2, 45);
    multiply(2.32, 5.37);
    multiply(2354235, 2235, 1);

    return 0;
}
