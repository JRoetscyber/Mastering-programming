pair<int, int> firstAndLast(vector<int> arr) {
    sort(arr.begin(), arr.end());
    return { arr.front(), arr.back() };
}

// What does firstAndLast need to return ? Just the smallest value and the largest value in the array
// nothing else.The current code sorts the entire array just to grab those two values off the ends.That's 
// like reorganizing your whole bookshelf alphabetically just to find the shortest and tallest book.
// Here's the key insight: you can find the min and max with a single pass, no sorting at all.

// Big O Notation of O(n)
pair<int, int> firstAndLast(vector<int> arr) {
    int minVal = arr[0];
    int maxVal = arr[0];

    for (int i = 1; i < arr.size(); i++) {
        if (arr[i] < minVal) minVal = arr[i];
        if (arr[i] > maxVal) maxVal = arr[i];
    }

    return { minVal, maxVal };
}
// feedback:
// This is the one to focus on. You said O(n), but sort() is O(n log n), not O(n). That's the actual complexity of the function as written.