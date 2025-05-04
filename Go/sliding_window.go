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

package main

import "fmt"

func slidingWindow() {
    arr := []int{1,3,2,6,-1,4,1,8,2}
    K := 5 // Window size
    n := len(arr)
    if n < K {
        return
    }

    // 1) Sum of first K elements
    windowSum := 0
    for i := 0; i < K; i++ {
        windowSum += arr[i] // Build initial window
    }

    // 2) Print average of first window
    fmt.Printf("Avg [0..%d] = %.1f\n", K-1, float64(windowSum)/float64(K))

    // 3) Slide the window one element at a time
    for end := K; end < n; end++ {
        // Subtract outgoing, add incoming
        windowSum += arr[end] - arr[end-K]
        fmt.Printf("Avg [%d..%d] = %.1f\n", end-K+1, end, float64(windowSum)/float64(K))
    }
} 