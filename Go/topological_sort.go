/*
╔════════════════════════════════════════════════════════╗
║           Topological Sort (Kahn's Algorithm)         ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Sort vertices in a DAG such that for every directed   ║
║ edge u→v, u comes before v. Use in-degree count and  ║
║ queue to process nodes with no incoming edges.        ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   Graph: 5→2→3→1                                      ║
║           ↓   ↑                                       ║
║           4 → 0                                       ║
║                                                      ║
║   in-degree: [1,1,1,0,0]                             ║
║   queue: [4,5] → process 4 → queue: [5,0]            ║
║   queue: [5,0] → process 5 → queue: [0,2]            ║
║   queue: [0,2] → process 0 → queue: [2,3]            ║
║   queue: [2,3] → process 2 → queue: [3]              ║
║   queue: [3] → process 3 → queue: [1]                ║
║   queue: [1] → process 1 → queue: []                 ║
║                                                      ║
║   Result: [4,5,0,2,3,1]                             ║
╚════════════════════════════════════════════════════════╝
*/

package main

import (
    "container/list"
    "fmt"
)

func topologicalSort(numVertices int, edges [][]int) []int {
    // Build adjacency list and in-degree count
    adj := make([][]int, numVertices)
    inDegree := make([]int, numVertices)

    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
        inDegree[v]++
    }

    // Initialize queue with vertices of in-degree 0
    queue := list.New()
    for i := 0; i < numVertices; i++ {
        if inDegree[i] == 0 {
            queue.PushBack(i)
        }
    }

    result := make([]int, 0, numVertices)
    for queue.Len() > 0 {
        u := queue.Remove(queue.Front()).(int)
        result = append(result, u)

        // Decrease in-degree of neighbors
        for _, v := range adj[u] {
            inDegree[v]--
            if inDegree[v] == 0 {
                queue.PushBack(v)
            }
        }
    }

    // Check for cycle
    if len(result) != numVertices {
        return nil // Cycle detected
    }

    return result
}

func main() {
    numVertices := 6
    edges := [][]int{
        {5, 2},
        {5, 0},
        {4, 0},
        {4, 1},
        {2, 3},
        {3, 1},
    }

    result := topologicalSort(numVertices, edges)
    if result == nil {
        fmt.Println("Cycle detected in graph")
    } else {
        fmt.Println("Topological order:", result)
    }
} 