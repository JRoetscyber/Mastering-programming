/*
 * FILE: medium_03.cpp
 * TOPIC: Control Flow Deep Dive - Day 2
 * DIFFICULTY: Medium
 *
 * TASK: Multi-Condition Guard Clauses & Early Returns
 * ----------------------------------------------------
 * Deeply nested if/else statements (the "Arrow Anti-Pattern" or "Pyramid of Doom")
 * make code hard to read and debug.
 *
 * BAD (nested pyramid):
 *   if (user != nullptr) {
 *       if (user->age >= 18) {
 *           if (user->hasLicense) {
 *               // do something
 *           }
 *       }
 *   }
 *
 * CLEAN (Guard clauses with early return):
 *   if (user == nullptr) return false;
 *   if (user->age < 18) return false;
 *   if (!user->hasLicense) return false;
 *   // do something!
 *
 * 1. Write a function:
 *      bool canProcessTransaction(int accountBalance, int amount, bool isAccountActive, bool isFraudFlagged)
 *
 * 2. Implement it using ONLY guard clauses (check failure conditions FIRST and return false immediately):
 *      - If account is NOT active: print "Declined: Account inactive", return false.
 *      - If fraud IS flagged: print "Declined: Fraud alert triggered", return false.
 *      - If amount <= 0: print "Declined: Invalid transaction amount", return false.
 *      - If amount > accountBalance: print "Declined: Insufficient funds", return false.
 *      - If all checks pass: print "Approved: Transaction processed!", return true.
 *
 * 3. In main(), test with 3 scenarios:
 *      - Scenario A: balance = 500, amount = 100, active = false, fraud = false
 *      - Scenario B: balance = 500, amount = 800, active = true, fraud = false
 *      - Scenario C: balance = 500, amount = 250, active = true, fraud = false
 *
 * EXPECTED OUTPUT:
 *   Test A: Declined: Account inactive
 *   Test B: Declined: Insufficient funds
 *   Test C: Approved: Transaction processed!
 */

#include <iostream>
using namespace std;

// TODO: Write canProcessTransaction here using guard clauses.
bool canProcessTransaction(int accountBalance, int amount, bool isAccountActive, bool isFraudFlagged) {
    if (!isAccountActive) {
        std::cout << "Declined: Account inactive" << "\n";
        return false;
    }
    if (isFraudFlagged) {
        std::cout << "Declined: Fraud alert triggered" << "\n";
        return false;
    }
    if (amount <= 0) {
        std::cout << "Declined: Invalid transaction amount" << "\n";
        return false;
    }
    if (amount > accountBalance) {
        std::cout << "Declined: Insufficient funds" << "\n";
        return false;
    }

    std::cout << "Approved: Transaction processed!" << "\n";
    return true;
}

int main() {
    // TODO: Write your solution here.
    //-Scenario A : balance = 500, amount = 100, active = false, fraud = false
    //-Scenario B : balance = 500, amount = 800, active = true, fraud = false
    //-Scenario C : balance = 500, amount = 250, active = true, fraud = false

    canProcessTransaction(500, 100, false, false);
    canProcessTransaction(500, 800, true, false);
    canProcessTransaction(500, 250, true, false);

    return 0;
}
