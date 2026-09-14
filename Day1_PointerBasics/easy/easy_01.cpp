/*
 * FILE: easy_01.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Easy
 *
 * TASK: Declare a pointer to an integer
 * ----------------------------------------
 * 1. Declare an int variable called `score` and assign it the value 95.
 * 2. Declare an int pointer called `p` and make it point to `score`.
 * 3. Print the VALUE of `score` using the pointer (dereference it).
 * 4. Print the MEMORY ADDRESS stored in `p`.
 *
 * EXPECTED OUTPUT (addresses will differ on your machine):
 *   Value of score via pointer: 95
 *   Address stored in p: 0x...
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.
    int score = 95;
    int* p = &score;
    cout << *p << endl;
    cout << p << endl;
    return 0;
}

/*ok the way i understand this is that the value is created ten the pointer is create the pointer is used to 
point to a value in a scerten memory address so score is created auto asigned the memory adress 0xc8c8fff8a4
then we create a pointer and point it to the score address (0xc8c8fff8a4) to print the value i need to 
to use * (astrickts) before the pointer name and to print the adress we use the & (amprisant) after playing a round 
 with the & it is not needed to print a pointer adress it is just used to asign a adress */
