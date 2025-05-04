#include <iostream>
#include <vector>
using namespace std;

/*
╔════════════════════════════════════════════════════════╗
║                  Subsets Pattern (BFS)                ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Start with empty subset, then for each number, take   ║
║ all existing subsets, append the number, and add to  ║
║ the list (BFS‑style build).                           ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ nums=[1,2,3]                                          ║
║ subsets=[[]]                                          ║
║ → for 1: new=[[1]]      → subsets=[[],[1]]            ║
║ → for 2: new=[[2],[1,2]]→ subsets=[[],[1],[2],[1,2]]  ║
║ → for 3: new=[[3],[1,3],[2,3],[1,2,3]] → final       ║
╚════════════════════════════════════════════════════════╝
*/

int main() {
    vector<int> nums = {1,2,3};
    vector<vector<int>> subsets = {{}};

    for (int x : nums) {
        int sz = subsets.size();
        for (int i = 0; i < sz; ++i) {
            auto curr = subsets[i]; // copy
            curr.push_back(x);
            subsets.push_back(curr);
        }
    }

    // Print all subsets
    for (auto &s : subsets) {
        cout << "[";
        for (int v : s) cout << v;
        cout << "] ";
    }
    cout << "\n";
    return 0;
} 