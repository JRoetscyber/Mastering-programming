// The slow findDuplicate function has a time complexity of O(n^2) because it uses two nested loops to compare each element with every other element in the array. This results in a quadratic number of comparisons as the size of the input array increases.
bool findDuplicate(vector<int>& arr) {
    for (int i = 0; i < arr.size(); i++) {
        for (int j = i + 1; j < arr.size(); j++) {
            if (arr[i] == arr[j]) return true;
        }
    }
    return false;
}

// The optimized findDuplicate function has a time complexity of O(n) because it uses a set to store unique elements from the array. The set allows for average O(1) time complexity for insertions and lookups, resulting in a linear number of operations as the size of the input array increases.
bool findDuplicate(vector<int>& arr) {
	set<int> Myset(arr.begin(), arr.end());
    if (arr.size() != Myset.size()) {
		return true;
    }
	return false;
}

// feedback:
// Correct. O(n²) → your set solution brings it to O(n) with O(n) extra space. Good tradeoff reasoning, and you correctly clocked set's average O(1) insert/lookup. (Small note: std::set is actually a red-black tree with O(log n) operations — you'd want unordered_set for true average O(1). Worth filing away, but your complexity instinct was right either way.)
// C++ has two different "set" types, and they behave very differently under the hood :
// std::set(what you used) is implemented as a red - black tree — a self - balancing binary search tree.Every element is kept in sorted order internally.Because of that, inserting an element means walking down the tree to find its correct sorted position, which takes O(log n) time.Do that for all n elements, and building the set costs O(n log n) total — not O(n) like you said in your explanation.
// std::unordered_set is implemented as a hash table instead.No sorting, no tree - walking — it just hashes the value to figure out where to put it.That gives you average O(1) insertion and lookup, so building it from n elements costs average O(n) total.
// So for your findDuplicate rewrite :
// cppbool findDuplicate(vector<int>&arr) {
//     set<int> Myset(arr.begin(), arr.end()); // O(n log n) — red-black tree
//     if (arr.size() != Myset.size()) return true;
//     return false;
// }
// Swapping set for unordered_set gets you the true O(n) average case you were aiming for:
//     cppbool findDuplicate(vector<int>&arr) {
//         unordered_set<int> Myset(arr.begin(), arr.end()); // average O(n) — hash table
//         if (arr.size() != Myset.size()) return true;
//         return false;
//     }
//     Your instinct to reach for a set - based dedup was completely correct — that's the right strategy. The only nuance was picking the variant that actually delivers O(n) rather than O(n log n). This is exactly the kind of "knowing what's under the hood" knowledge that separates pattern-recognition from true mastery, so good catch wanting to understand it.