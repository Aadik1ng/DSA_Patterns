/*
╔════════════════════════════════════════════════════════╗
║          In‑Place LinkedList Reversal Pattern         ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Reverse the `next` pointers by walking through the    ║
║ list and, for each node, redirecting its next to the  ║
║ previous node.                                        ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║   Input: 1 → 2 → 3 → null                              ║
║                                                      ║
║   prev=null, cur=1                                   ║
║      cur->next = prev  → 1→null                       ║
║      prev=1, cur=2                                   ║
║                                                      ║
║   prev=1→null, cur=2                                 ║
║      2→next = 1→null  → 2→1→null                      ║
║      prev=2, cur=3                                   ║
║                                                      ║
║   prev=2→1→null, cur=3                               ║
║      3→next = 2→1→null  → 3→2→1→null                  ║
║   Done: new head = prev (3)                          ║
╚════════════════════════════════════════════════════════╝
*/

package main

import "fmt"

type Node struct {
    Val  int
    Next *Node
}

// Reverse the linked list in-place
func reverseList(head *Node) *Node {
    var prev *Node
    cur := head

    for cur != nil {
        next := cur.Next // save next
        cur.Next = prev  // reverse link
        prev = cur       // advance prev
        cur = next       // advance cur
    }
    return prev // new head
}

func main() {
    // Build list: 1->2->3
    head := &Node{Val: 1}
    head.Next = &Node{Val: 2}
    head.Next.Next = &Node{Val: 3}

    rev := reverseList(head)
    for rev != nil {
        fmt.Printf("%d ", rev.Val)
        rev = rev.Next
    }
    fmt.Println()
} 