/*
 * FILE: hard_01.cpp
 * TOPIC: Functions Deep Dive - Day 3
 * DIFFICULTY: Hard
 *
 * TASK: Operator Overloading
 * --------------------------
 * Operator overloading lets you teach C++ how to use math operators (+, -, ==)
 * on custom data types like Structs and Classes.
 *
 * 1. Define a struct `Vector2 { float x, y; };`
 * 2. Write an operator overload for `+` inside the struct:
 *      Vector2 operator+(const Vector2& other) const { ... }
 * 3. Write an operator overload for `==` that returns true if both x and y match.
 *
 * 4. In main, create `v1 = {1.5, 2.0}` and `v2 = {3.0, 4.0}`.
 * 5. Add them together: `Vector2 v3 = v1 + v2;`
 *
 * EXPECTED OUTPUT:
 *   v3 coordinates: 4.5, 6
 *   v1 equals v2? 0 (false)
 */

#include <iostream>
using namespace std;

// TODO: Define Vector2 struct and overload + and ==
struct Vector2 {
    float x, y;
};

int main() {
    // TODO: Write your solution here.

    return 0;
}
