/*
╔════════════════════════════════════════════════════════╗
║                Subsets Pattern (BFS)                  ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Generate all subsets by iteratively adding each       ║
║ number to existing subsets, creating new ones.        ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   nums = [1,2,3]                                      ║
║                                                      ║
║   Start: [[]]                                         ║
║   Add 1: [[],[1]]                                    ║
║   Add 2: [[],[1],[2],[1,2]]                          ║
║   Add 3: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]] ║
╚════════════════════════════════════════════════════════╝
*/

package main

import "fmt"

func subsets(nums []int) [][]int {
    result := [][]int{{}} // Start with empty subset

    for _, num := range nums {
        n := len(result)
        for i := 0; i < n; i++ {
            // Create new subset by adding current number
            newSubset := make([]int, len(result[i]))
            copy(newSubset, result[i])
            newSubset = append(newSubset, num)
            result = append(result, newSubset)
        }
    }

    return result
}

func main() {
    nums := []int{1, 2, 3}
    result := subsets(nums)
    
    fmt.Println("All subsets:")
    for _, subset := range result {
        fmt.Println(subset)
    }
} 