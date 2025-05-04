"""
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
"""
from collections import deque

def topo_sort(V, adj):
    indegree = [0]*V
    order = []
    q = deque()
    for u in range(V):
        for v in adj[u]:
            indegree[v] += 1
    for i in range(V):
        if indegree[i] == 0:
            q.append(i)
    while q:
        u = q.popleft()
        order.append(u)
        for v in adj[u]:
            indegree[v] -= 1
            if indegree[v] == 0:
                q.append(v)
    return order

def main():
    V = 6
    graph = [
        [1,2], [3], [3,4], [5], [5], []
    ]
    order = topo_sort(V, graph)
    print(*order)

if __name__ == "__main__":
    main() 