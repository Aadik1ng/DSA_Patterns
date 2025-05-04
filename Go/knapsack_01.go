/*
╔════════════════════════════════════════════════════════╗
║                0/1 Knapsack Pattern                   ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Dynamic programming solution to maximize value while  ║
║ staying within weight capacity. dp[i][w] = max value  ║
║ using first i items with weight w.                    ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   weights = [1,2,3], values = [6,10,12], W=5         ║
║                                                      ║
║   dp[0][w] = 0 for all w                             ║
║   dp[i][0] = 0 for all i                             ║
║                                                      ║
║   i=1: w=1 → dp[1][1]=6                              ║
║   i=2: w=2 → dp[2][2]=max(6,10)=10                   ║
║   i=3: w=3 → dp[3][3]=max(12,dp[2][3])=12            ║
║   i=3: w=5 → dp[3][5]=max(12+dp[2][2],dp[2][5])=22   ║
║                                                      ║
║   Final answer: dp[3][5]=22                          ║
╚════════════════════════════════════════════════════════╝
*/

package main

import "fmt"

func knapsack(weights []int, values []int, W int) int {
    n := len(weights)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, W+1)
    }

    for i := 1; i <= n; i++ {
        for w := 1; w <= W; w++ {
            if weights[i-1] <= w {
                // Can include current item
                dp[i][w] = max(values[i-1]+dp[i-1][w-weights[i-1]], dp[i-1][w])
            } else {
                // Cannot include current item
                dp[i][w] = dp[i-1][w]
            }
        }
    }

    return dp[n][W]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    weights := []int{1, 2, 3}
    values := []int{6, 10, 12}
    W := 5

    maxValue := knapsack(weights, values, W)
    fmt.Printf("Maximum value: %d\n", maxValue)
} 