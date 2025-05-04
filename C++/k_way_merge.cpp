#include <iostream>
#include <vector>
#include <queue>
using namespace std;

struct Item {
    int val, listIdx, elemIdx;
    bool operator>(const Item& o) const { return val > o.val; }
};

/*
╔════════════════════════════════════════════════════════╗
║                 K‑way Merge Pattern                  ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Push the first element of each of the K sorted lists ║
║ into a min‑heap. Pop the smallest, then push its      ║
║ successor from the same list, repeating until empty. ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ lists=[[1,4,7],[2,5,8],[3,6,9]]                       ║
║                                                       ║
║ Heap init: [1@L0, 2@L1, 3@L2]                         ║
║ Pop 1@L0 → output 1, push 4@L0                         ║
║ Heap: [2@L1,3@L2,4@L0]                                 ║
║ Pop 2@L1 → output 2, push 5@L1                         ║
║ …                                                     ║
║ Final output: 1,2,3,4,5,6,7,8,9                       ║
╚════════════════════════════════════════════════════════╝
*/

int main() {
    vector<vector<int>> lists = {{1,4,7},{2,5,8},{3,6,9}};
    priority_queue<Item, vector<Item>, greater<Item>> minH;

    // 1) Initialize heap with first elem of each list
    for (int i = 0; i < lists.size(); ++i)
        if (!lists[i].empty())
            minH.push({lists[i][0], i, 0});

    vector<int> result;
    // 2) Extract‑and‑push loop
    while (!minH.empty()) {
        auto cur = minH.top(); minH.pop();
        result.push_back(cur.val);

        // push next from same list
        if (cur.elemIdx + 1 < lists[cur.listIdx].size())
            minH.push({ 
                lists[cur.listIdx][cur.elemIdx+1],
                cur.listIdx,
                cur.elemIdx+1
            });
    }

    // 3) Print merged result
    for (int x : result) cout << x << " ";
    cout << "\n";
    return 0;
} 