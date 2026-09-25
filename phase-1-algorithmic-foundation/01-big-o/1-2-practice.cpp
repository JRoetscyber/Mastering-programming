// This is a O(1) time complexity function because it retrieves a value from an unordered_map using a key, which has an average time complexity of O(1) for lookups. The function takes an integer id and an unordered_map that maps integers to strings, and it returns the string associated with the given id in constant time.
string getUser(int id, unordered_map<int, string>& database) {
    return database[id];
}
// This can not be improved because the time complexity of O(1) is already optimal for this operation.

// feedback:
// Correct. O(1), and correctly identified it's already optimal. You're not just pattern-matching here, you understood why it can't get faster that's good depth.