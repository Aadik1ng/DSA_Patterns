#include <iostream>
#include <vector>
using namespace std;

/*
╔════════════════════════════════════════════════════════╗
║            Two‑Pointer Pattern (Sorted Array)          ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Use two indices—one at start, one at end—to find a    ║
║ pair whose sum equals the target in O(N) time.        ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   arr = [1,2,3,4,6], target=6                          ║
║                                                      ║
║     left=0           right=4                           ║
║       ↓                ↑                               ║
║       1  2  3  4  6                                      ║
║       |--------------|   sum=1+6=7>6 → decrement right ║
║                                                      ║
║     left=0         right=3                             ║
║       ↓              ↑                                 ║
║       1  2  3  4  6    sum=1+4=5<6 → increment left    ║
║                                                      ║
║     left=1       right=3                               ║
║        ↓           ↑                                   ║
║       1  2  3  4  6      sum=2+4=6 → found             ║
╚════════════════════════════════════════════════════════╝
*/

int main() {
    vector<int> arr = {1,2,3,4,6};
    int target = 6;
    int left = 0, right = arr.size() - 1;

    // Move inward until pointers meet
    while (left < right) {
        int sum = arr[left] + arr[right];
        if (sum == target) {
            cout << "Pair: (" << arr[left] << "," << arr[right] << ")\n";
            return 0;
        }
        if (sum < target) ++left;   // Need larger sum
        else --right;               // Need smaller sum
    }

    cout << "No valid pair\n";
    return 0;
} 