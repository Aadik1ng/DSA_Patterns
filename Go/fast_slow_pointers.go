/*
╔════════════════════════════════════════════════════════╗
║         Fast & Slow Pointers (Cycle Detection)        ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Use two pointers moving at different speeds—slow      ║
║ moves 1 step, fast moves 2 steps. In a cycle, they   ║
║ must eventually meet.                                 ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   a -> b -> c                                         ║
║         ↑    ↓                                        ║
║         |____| ← fast loops back                        ║
║   slow: a→b→c→b→c→…                                    ║
║   fast: a→c→b→c→b→…                                    ║
║   => slow == fast                                       ║
╚════════════════════════════════════════════════════════╝
*/

package main

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

// Returns true if cycle exists
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next      // 1 step
        fast = fast.Next.Next // 2 steps
        if slow == fast {
            return true
        }
    }
    return false
}

func main() {
    // Build a cycle: a->b->c->b...
    a := &ListNode{Val: 1}
    b := &ListNode{Val: 2}
    c := &ListNode{Val: 3}
    a.Next = b
    b.Next = c
    c.Next = b

    if hasCycle(a) {
        fmt.Println("Cycle detected")
    } else {
        fmt.Println("No cycle")
    }
} 