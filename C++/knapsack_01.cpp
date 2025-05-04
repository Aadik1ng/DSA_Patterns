#include <iostream>
#include <vector>
using namespace std;

/*
╔════════════════════════════════════════════════════════╗
║          0/1 Knapsack Pattern (Dynamic Programming)   ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Build a 1D DP array dp[w] = max value for capacity w. ║
║ For each item, iterate w from W→weight to avoid reuse.║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ weights=[1,3,4], values=[15,50,60], W=8               ║
║ dp initially [0,0,0,0,0,0,0,0,0]                     ║
║ Item0(w=1,v=15):                                     ║
║  dp[8]=max(dp[8],15+dp[7]) → … → dp[1]=15              ║
║ Item1(w=3,v=50):                                     ║
║  …                                                   ║
║ Item2(w=4,v=60):                                     ║
║  …                                                   ║
║ Result=dp[8]=80 (items 1 and 2)                       ║
╚════════════════════════════════════════════════════════╝
*/

int knapsack(int W, const vector<int>& wt, const vector<int>& val) {
    int n = wt.size();
    vector<int> dp(W+1, 0);

    // For each item
    for (int i = 0; i < n; ++i) {
        // Traverse capacities downward
        for (int w = W; w >= wt[i]; --w) {
            dp[w] = max(dp[w], val[i] + dp[w - wt[i]]);
        }
    }
    return dp[W];
}

int main() {
    vector<int> weights = {1,3,4};
    vector<int> values  = {15,50,60};
    int capacity = 8;
    cout << "Max value = " << knapsack(capacity, weights, values) << "\n";
    return 0;
} 