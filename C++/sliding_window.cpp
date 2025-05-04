#include <iostream>
#include <vector>
using namespace std;

/*
╔════════════════════════════════════════════════════════╗
║            Sliding Window Pattern (Size K)            ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Maintain a running sum of a K‑sized window and slide   ║
║ across the array in O(N) time by subtracting the      ║
║ element that "goes out" and adding the new one "in".   ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ Array: [1,3,2,6,-1,4,1,8,2], K=5                       ║
║                                                      ║
║   Initial Window:                                      ║
║   Indexes:   0   1   2   3    4    5    6    7   8     ║
║   Elements: [1] – [3] – [2] – [6] – [ -1 ]   4    1    8  2 ║
║              |--------------------------------|         ║
║                                                      ║
║   Slide Right by 1:                                    ║
║   Remove arr[0], Add arr[5]                            ║
║   [3] – [2] – [6] – [ -1 ] – [4]                        ║
║                 |--------------------------------|       ║
╚════════════════════════════════════════════════════════╝
*/

int main() {
    vector<int> arr = {1,3,2,6,-1,4,1,8,2};
    int K = 5;                      // Window size
    int n = arr.size();
    if (n < K) return 0;

    // 1) Sum of first K elements
    int window_sum = 0;
    for (int i = 0; i < K; ++i)
        window_sum += arr[i];       // Build initial window

    // 2) Print average of first window
    cout << "Avg [0.." << (K-1) << "] = "
         << (window_sum / double(K)) << "\n";

    // 3) Slide the window one element at a time
    for (int end = K; end < n; ++end) {
        //   Subtract outgoing, add incoming
        window_sum += arr[end] - arr[end - K];
        cout << "Avg [" << (end-K+1) << ".." << end << "] = "
             << (window_sum / double(K)) << "\n";
    }
    return 0;
} 