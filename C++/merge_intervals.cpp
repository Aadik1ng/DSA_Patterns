#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

/*
╔════════════════════════════════════════════════════════╗
║               Merge Intervals Pattern                 ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Given a list of intervals, sort by start time and     ║
║ merge overlapping ones by comparing the current end   ║
║ with the next start.                                   ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   Input: [[1,3],[2,6],[8,10],[15,18]]                 ║
║                                                      ║
║   After Sort: [[1,3],[2,6],[8,10],[15,18]]            ║
║     merged = [1,3]                                    ║
║     next = [2,6]  → overlap? 3 >= 2 → merge end=max(6,3)=6 ║
║     merged = [1,6]                                    ║
║     next = [8,10] → 6 < 8 → push [8,10]                ║
║     …                                                  ║
║   Output: [[1,6],[8,10],[15,18]]                      ║
╚════════════════════════════════════════════════════════╝
*/

int main() {
    vector<pair<int,int>> intervals = {{1,3},{2,6},{8,10},{15,18}};
    // 1) Sort by start
    sort(intervals.begin(), intervals.end());

    vector<pair<int,int>> merged;
    for (auto &in : intervals) {
        if (merged.empty() || merged.back().second < in.first) {
            // no overlap
            merged.push_back(in);
        } else {
            // overlap → extend end
            merged.back().second = max(merged.back().second, in.second);
        }
    }
    // 2) Print result
    for (auto &in : merged)
        cout << "[" << in.first << "," << in.second << "] ";
    cout << "\n";
    return 0;
} 