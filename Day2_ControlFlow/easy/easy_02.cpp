/*
 * FILE: easy_02.cpp
 * TOPIC: Control Flow Deep Dive - Day 2
 * DIFFICULTY: Easy
 *
 * TASK: switch Statement with intentional fallthrough vs break
 * -------------------------------------------------------------
 * A `switch` evaluates an integral expression and jumps directly to a matching `case:`.
 * If you omit `break;`, execution "falls through" to subsequent cases!
 * In modern C++, intentional fallthrough can be annotated with `[[fallthrough]];`
 * to silence compiler warnings.
 *
 * 1. Declare a char `severity = 'W'` ('E' = Error, 'W' = Warning, 'I' = Info).
 * 2. Write a function or switch block that takes `severity`:
 *    - Case 'E' or 'e': print "Level: ERROR (Critical action required)"
 *    - Case 'W' or 'w': print "Level: WARNING (Check system logs)"
 *    - Case 'I' or 'i': print "Level: INFO (Normal status)"
 *    - default: print "Level: UNKNOWN"
 * 3. Notice how grouping 'E' and 'e' without a break is intentional fallthrough
 *    to handle both uppercase and lowercase letters!
 * 4. Test it with 'W', 'e', and an invalid char like 'X'.
 *
 * EXPECTED OUTPUT:
 *   Testing 'W': Level: WARNING (Check system logs)
 *   Testing 'e': Level: ERROR (Critical action required)
 *   Testing 'X': Level: UNKNOWN
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.



    return 0;
}
