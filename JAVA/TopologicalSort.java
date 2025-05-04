import java.util.*;

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

public class TopologicalSort {
    public static List<Integer> topoSort(int V, List<List<Integer>> adj) {
        int[] indegree = new int[V];
        List<Integer> order = new ArrayList<>();
        Queue<Integer> q = new LinkedList<>();
        for (int u = 0; u < V; ++u)
            for (int v : adj.get(u))
                ++indegree[v];
        for (int i = 0; i < V; ++i)
            if (indegree[i] == 0)
                q.add(i);
        while (!q.isEmpty()) {
            int u = q.poll();
            order.add(u);
            for (int v : adj.get(u)) {
                if (--indegree[v] == 0)
                    q.add(v);
            }
        }
        return order;
    }
    public static void main(String[] args) {
        int V = 6;
        List<List<Integer>> graph = Arrays.asList(
            Arrays.asList(1,2), Arrays.asList(3), Arrays.asList(3,4), Arrays.asList(5), Arrays.asList(5), new ArrayList<>()
        );
        List<Integer> order = topoSort(V, graph);
        for (int u : order)
            System.out.print(u + " ");
        System.out.println();
    }
} 