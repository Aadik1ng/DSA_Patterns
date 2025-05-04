/*
╔════════════════════════════════════════════════════════╗
║               Tree Breadth‑First Search               ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Level‑order traversal using a FIFO queue: push root,  ║
║ then for each node pop→visit→enqueue its children.    ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║      1                                               ║
║     / \                                              ║
║    2   3                                             ║
║   /                                               ║
║  4                                                ║
║                                                      ║
║ Queue evolves:                                       ║
║  [1] → visit 1, enqueue [2,3]                        ║
║  [2,3] → visit 2, enqueue [3,4]                      ║
║  [3,4] → visit 3, enqueue [4]                        ║
║  [4] → visit 4, enqueue [] → done                    ║
╚════════════════════════════════════════════════════════╝
*/

package main

import (
    "container/list"
    "fmt"
)

type TreeNode struct {
    Val      int
    Children []*TreeNode
}

func bfs(root *TreeNode) {
    if root == nil {
        return
    }

    queue := list.New()
    queue.PushBack(root)

    for queue.Len() > 0 {
        // Pop front
        node := queue.Remove(queue.Front()).(*TreeNode)
        fmt.Printf("%d ", node.Val)

        // Enqueue children
        for _, child := range node.Children {
            queue.PushBack(child)
        }
    }
}

func main() {
    // Build tree: 1→{2,3}, 2→{4}
    root := &TreeNode{Val: 1}
    root.Children = []*TreeNode{
        {Val: 2},
        {Val: 3},
    }
    root.Children[0].Children = []*TreeNode{
        {Val: 4},
    }

    bfs(root)
    fmt.Println()
} 