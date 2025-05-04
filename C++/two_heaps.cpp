#include <iostream>
#include <queue>
using namespace std;

/*
╔════════════════════════════════════════════════════════╗
║              Two Heaps Pattern (Median)               ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Maintain a max‑heap for the lower half and a min‑heap ║
║ for the upper half of the stream to get median in O(log N) per insert. ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ Stream: 5,15,1,3                                      ║
║                                                      ║
║ Insert 5 → maxH:[5], minH:[] → median=5               ║
║ Insert 15→ maxH:[5], minH:[15] → median=(5+15)/2=10   ║
║ Insert 1 → maxH:[5,1], minH:[15] → median=5           ║
║ Insert 3 → maxH:[3,1], minH:[5,15] → median=(3+5)/2=4 ║
╚════════════════════════════════════════════════════════╝
*/

int main() {
    priority_queue<int>      maxH;   // lower half
    priority_queue<int, vector<int>, greater<int>> minH; // upper half

    int nums[] = {5, 15, 1, 3};
    for (int x : nums) {
        // 1) Add new number to one of the heaps
        if (maxH.empty() || x < maxH.top())
            maxH.push(x);
        else
            minH.push(x);

        // 2) Rebalance to ensure size difference ≤1
        if (maxH.size() > minH.size() + 1) {
            minH.push(maxH.top());
            maxH.pop();
        } else if (minH.size() > maxH.size()) {
            maxH.push(minH.top());
            minH.pop();
        }

        // 3) Compute median
        if (maxH.size() == minH.size())
            cout << "Median: " << ((maxH.top() + minH.top())/2.0) << "\n";
        else
            cout << "Median: " << maxH.top() << "\n";
    }
    return 0;
} 