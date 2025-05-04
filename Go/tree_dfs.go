/*
╔════════════════════════════════════════════════════════╗
║             Tree Depth‑First Search (DFS)             ║
╠════════════════════════════════════════════════════════╣
║ Description:                                           ║
║ Pre‑order traversal via recursion (or a stack):       ║
║ visit node, then recurse into each child.            ║
╠════════════════════════════════════════════════════════╣
║ Flow Diagram (ASCII):                                  ║
║      1                                               ║
║     / \                                              ║
║    2   3 → 5                                        ║
║                                                      ║
║ Visit order: 1 → 2 → (back) → 3 → 5                  ║
╚════════════════════════════════════════════════════════╝
*/

package main

import "fmt"

type TreeNode struct {
    Val      int
    Children []*TreeNode
}

func dfs(node *TreeNode) {
    if node == nil {
        return
    }
    fmt.Printf("%d ", node.Val) // visit
    for _, child := range node.Children {
        dfs(child) // recurse
    }
}

func main() {
    // Build tree: 1→{2,3}, 3→{5}
    root := &TreeNode{Val: 1}
    root.Children = []*TreeNode{
        {Val: 2},
        {Val: 3},
    }
    root.Children[1].Children = []*TreeNode{
        {Val: 5},
    }

    dfs(root)
    fmt.Println()
} 