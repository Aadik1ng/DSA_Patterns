/*
╔════════════════════════════════════════════════════════╗
║              Two Heaps Pattern (Median)               ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Maintain a max‑heap for the lower half and a min‑heap ║
║ for the upper half of the stream to get median in O(log N) per insert. ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║ Stream: 5,15,1,3                                      ║
║                                                      ║
║ Insert 5 → maxH:[5], minH:[] → median=5               ║
║ Insert 15→ maxH:[5], minH:[15] → median=(5+15)/2=10   ║
║ Insert 1 → maxH:[5,1], minH:[15] → median=5           ║
║ Insert 3 → maxH:[3,1], minH:[5,15] → median=(3+5)/2=4 ║
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

// MaxHeap implementation
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func main() {
    maxH := &MaxHeap{} // lower half
    minH := &MinHeap{} // upper half
    heap.Init(maxH)
    heap.Init(minH)

    nums := []int{5, 15, 1, 3}
    for _, x := range nums {
        // 1) Add new number to one of the heaps
        if maxH.Len() == 0 || x < (*maxH)[0] {
            heap.Push(maxH, x)
        } else {
            heap.Push(minH, x)
        }

        // 2) Rebalance to ensure size difference ≤1
        if maxH.Len() > minH.Len()+1 {
            heap.Push(minH, heap.Pop(maxH))
        } else if minH.Len() > maxH.Len() {
            heap.Push(maxH, heap.Pop(minH))
        }

        // 3) Compute median
        if maxH.Len() == minH.Len() {
            fmt.Printf("Median: %.1f\n", float64((*maxH)[0]+(*minH)[0])/2.0)
        } else {
            fmt.Printf("Median: %d\n", (*maxH)[0])
        }
    }
} 