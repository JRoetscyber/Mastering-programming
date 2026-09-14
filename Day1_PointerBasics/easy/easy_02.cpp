/*
 * FILE: easy_02.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Easy
 *
 * TASK: Use the address-of operator (&) on multiple types
 * ---------------------------------------------------------
 * 1. Declare the following variables:
 *      char   letter = 'A';
 *      int    count  = 7;
 *      double price  = 3.99;
 * 2. Print the address of EACH variable using the & operator.
 *    (Cast char* and double* to void* so std::cout prints them as
 *     hex addresses rather than characters or numbers.)
 *
 * EXPECTED OUTPUT (addresses will differ):
 *   Address of letter : 0x...
 *   Address of count  : 0x...
 *   Address of price  : 0x...
 *
 * HINT: To print the address of a char without printing the char itself,
 *       cast it: static_cast<void*>(&letter)
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.
    char letter = 'A';
    int count = 7;
    double price = 3.99;

    cout << "Address of letter: " << (void*)&letter << "\n";
    cout << "Address of count: " << &count << "\n";
    cout << "Address of price: " << &price << "\n";

    return 0;
}

/*so letter are a pain in the ass not going to lei to you look like int, doubles and floats are great letter from what i read is consider c
string and the std::cout doesn't know what to do with that why i got Address of letter: A0ⁿc+ then i leard about casting it with (void*) still don't 100% 
how this makes it work thought*/
