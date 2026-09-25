/*
 * FILE: medium_02.cpp
 * TOPIC: Control Flow Deep Dive - Day 2
 * DIFFICULTY: Medium
 *
 * TASK: State Machine using switch inside a while loop
 * -----------------------------------------------------
 * A common pattern in game development, networking, and robotics is a
 * Finite State Machine (FSM) implemented with an enum class + switch.
 *
 * 1. Define an enum class for states:
 *      enum class State { IDLE, CONNECTING, CONNECTED, DISCONNECTED };
 *
 * 2. In main(), start with:
 *      State currentState = State::IDLE;
 *      int attempts = 0;
 *
 * 3. Run a `while (currentState != State::DISCONNECTED)` loop:
 *      Inside the loop, use a `switch (currentState)`:
 *        - Case IDLE:
 *            print "System IDLE. Initiating connection..."
 *            transition to CONNECTING
 *
 *        - Case CONNECTING:
 *            increment attempts
 *            print "Attempting connection (attempt " << attempts << ")..."
 *            if attempts < 3: remain in CONNECTING
 *            if attempts == 3: transition to CONNECTED
 *
 *        - Case CONNECTED:
 *            print "Connection established! Performing work..."
 *            print "Work finished. Disconnecting..."
 *            transition to DISCONNECTED
 *
 * 4. After exiting the loop, print: "Session closed gracefully."
 *
 * EXPECTED OUTPUT:
 *   System IDLE. Initiating connection...
 *   Attempting connection (attempt 1)...
 *   Attempting connection (attempt 2)...
 *   Attempting connection (attempt 3)...
 *   Connection established! Performing work...
 *   Work finished. Disconnecting...
 *   Session closed gracefully.
 */

#include <iostream>
using namespace std;

enum class State { IDLE, CONNECTING, CONNECTED, DISCONNECTED };

int main() {
    // TODO: Write your solution here.
    State currentState = State::IDLE;
    int attempts = 0;

    while (currentState != State::DISCONNECTED) {
        switch (currentState) {
            case State::IDLE:
                std::cout << "System IDLE. Initiating connection..." << "\n";
                currentState = State::CONNECTING;
                break;
            case State::CONNECTING:
                attempts++;
                std::cout << "Attempting connection (attempt " << attempts << ")..." << "\n";
                if (attempts < 3) {
                    currentState = State::CONNECTING;
                }
                if (attempts == 3) {
                    currentState = State::CONNECTED;
                }
                break;
            case State::CONNECTED:
                std::cout << "Connection established! Performing work..." << "\n";
                std::cout << "Work finished. Disconnecting..." << "\n";
                currentState = State::DISCONNECTED;
                break;
        }
    }
    std::cout << "Session closed gracefully." << "\n";

    return 0;
}
