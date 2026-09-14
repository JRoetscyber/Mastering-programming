/*
 * FILE: easy_05.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Easy
 *
 * TASK: Re-point a pointer to a different variable
 * -------------------------------------------------
 * A pointer is just a variable - you can change what it points to.
 *
 * 1. Declare two ints: `a = 10` and `b = 20`.
 * 2. Declare an int pointer `p` and initially point it to `a`.
 * 3. Print: "p points to: " followed by *p.
 * 4. Now re-assign `p` so it points to `b` instead.
 * 5. Print: "p now points to: " followed by *p.
 *
 * EXPECTED OUTPUT:
 *   p points to    : 10
 *   p now points to: 20
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.
    int a = 10;
    int b = 20;
    int* p = &a;

    cout << "p points to: " << *p << "\n";
    
    p = &b;

    cout << "p points to: " << *p << "\n";


    return 0;
}
