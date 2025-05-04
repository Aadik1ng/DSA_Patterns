#include <iostream>
#include <vector>
#include <queue>
using namespace std;

/*
╔════════════════════════════════════════════════════════╗
║         Topological Sort (Kahn's Algorithm)           ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ 1) Compute indegree of each node.                     ║
║ 2) Enqueue all zero‑indegree nodes.                   ║
║ 3) Pop, add to order, decrement neighbors' indegree. ║
║ 4) Enqueue new zero‑indegree until done.              ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ Graph:                                                 ║
║ 0→1,0→2; 1→3;2→3,2→4;3→5;4→5                            ║
║ indegree: [0,1,1,2,1,2]                                ║
║ Queue init: [0] → order=[ ]                            ║
║ Pop 0 → order=[0] → indegree[1]-- →0 , [2]-- →0 → enqueue[1,2] ║
║ Pop 1 → order=[0,1] → indegree[3]-- →1 → …             ║
║ → Final order: [0,1,2,3,4,5]                           ║
╚════════════════════════════════════════════════════════╝
*/

vector<int> topoSort(int V, const vector<vector<int>>& adj) {
    vector<int> indegree(V, 0), order;
    queue<int> q;

    // 1) Count indegrees
    for (int u = 0; u < V; ++u)
        for (int v : adj[u])
            ++indegree[v];

    // 2) Enqueue zeros
    for (int i = 0; i < V; ++i)
        if (indegree[i] == 0)
            q.push(i);

    // 3) Process
    while (!q.empty()) {
        int u = q.front(); q.pop();
        order.push_back(u);
        for (int v : adj[u]) {
            if (--indegree[v] == 0)
                q.push(v);
        }
    }
    return order;
}

int main() {
    int V = 6;
    vector<vector<int>> graph = {
        {1,2}, {3}, {3,4}, {5}, {5}, {}
    };
    auto order = topoSort(V, graph);
    for (int u : order)
        cout << u << " ";
    cout << "\n";
    return 0;
} 