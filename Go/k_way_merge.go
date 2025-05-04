/*
╔════════════════════════════════════════════════════════╗
║                K-way Merge Pattern                    ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Merge K sorted lists using a min-heap to always       ║
║ pick the smallest element from all lists.             ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   lists = [[1,4,5],[1,3,4],[2,6]]                    ║
║                                                      ║
║   Heap: [1,1,2] → pop 1                              ║
║   Heap: [1,2,4] → pop 1                              ║
║   Heap: [2,3,4] → pop 2                              ║
║   Heap: [3,4,6] → pop 3                              ║
║   Heap: [4,4,6] → pop 4                              ║
║   Heap: [4,5,6] → pop 4                              ║
║   Heap: [5,6] → pop 5                                ║
║   Heap: [6] → pop 6                                  ║
║                                                      ║
║   Result: [1,1,2,3,4,4,5,6]                         ║
╚════════════════════════════════════════════════════════╝
*/

package main

import (
    "container/heap"
    "fmt"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func mergeKLists(lists []*ListNode) *ListNode {
    minHeap := &MinHeap{}
    heap.Init(minHeap)

    // Push first node of each list into heap
    for _, list := range lists {
        if list != nil {
            heap.Push(minHeap, list)
        }
    }

    dummy := &ListNode{}
    current := dummy

    for minHeap.Len() > 0 {
        node := heap.Pop(minHeap).(*ListNode)
        current.Next = node
        current = current.Next

        if node.Next != nil {
            heap.Push(minHeap, node.Next)
        }
    }

    return dummy.Next
}

func main() {
    // Create sample lists
    list1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
    list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
    list3 := &ListNode{Val: 2, Next: &ListNode{Val: 6}}

    lists := []*ListNode{list1, list2, list3}
    merged := mergeKLists(lists)

    // Print merged list
    fmt.Print("Merged list: ")
    for merged != nil {
        fmt.Printf("%d ", merged.Val)
        merged = merged.Next
    }
    fmt.Println()
} 