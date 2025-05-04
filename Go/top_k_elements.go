/*
╔════════════════════════════════════════════════════════╗
║                Top K Elements Pattern                 ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Use a min-heap to keep track of the K largest        ║
║ elements. When heap size exceeds K, remove smallest.  ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   nums = [3,1,5,12,2,11], K=3                         ║
║                                                      ║
║   Add 3: [3]                                          ║
║   Add 1: [1,3]                                        ║
║   Add 5: [1,3,5]                                      ║
║   Add 12: [3,5,12] (removed 1)                       ║
║   Add 2: [3,5,12] (2 < 3, ignored)                   ║
║   Add 11: [5,11,12] (removed 3)                      ║
║                                                      ║
║   Final top 3: [5,11,12]                             ║
╚════════════════════════════════════════════════════════╝
*/

package main

import (
    "container/heap"
    "fmt"
)

// MinHeap implementation
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func topKFrequent(nums []int, k int) []int {
    minHeap := &MinHeap{}
    heap.Init(minHeap)

    for _, num := range nums {
        heap.Push(minHeap, num)
        if minHeap.Len() > k {
            heap.Pop(minHeap)
        }
    }

    result := make([]int, k)
    for i := k - 1; i >= 0; i-- {
        result[i] = heap.Pop(minHeap).(int)
    }
    return result
}

func main() {
    nums := []int{3, 1, 5, 12, 2, 11}
    k := 3
    result := topKFrequent(nums, k)
    fmt.Printf("Top %d elements: %v\n", k, result)
} 