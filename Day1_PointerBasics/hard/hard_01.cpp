/*
 * FILE: hard_01.cpp
 * TOPIC: Pointer Basics - Day 1
 * DIFFICULTY: Hard
 *
 * TASK: Pointer arithmetic - walk through an array using only a pointer
 * ---------------------------------------------------------------------
 * 1. Declare: int nums[] = {2, 4, 6, 8, 10};
 * 2. Declare an int pointer `p` pointing to the FIRST element.
 * 3. Using ONLY pointer arithmetic (++p, p+n, etc.) and NO array indexing
 *    (no nums[i] or p[i]), print every element of the array on one line,
 *    separated by spaces.
 * 4. After printing all elements, move the pointer back to the first element
 *    and print the address difference between the last and first element
 *    in terms of ELEMENTS (i.e., how many ints apart they are).
 *
 * EXPECTED OUTPUT:
 *   Elements: 2 4 6 8 10
 *   Element span: 4
 */

#include <iostream>
using namespace std;

int main() {
    // TODO: Write your solution here.
    int nums[] = { 2, 4, 6, 8, 10};
    
    cout << "Elements: ";
    for (int* p = nums; p < end(nums); p++) {
        cout << " " << *p;
    };
    cout << "\n";

    int* first = nums;
    int* last = &nums[4];
    cout << "Element span: " << (last - first) << "\n";
    return 0;
}
