"""
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
"""

class TreeNodeDFS:
    def __init__(self, val):
        self.val = val
        self.children = []

def dfs(node):
    if not node:
        return
    print(node.val, end=' ')
    for ch in node.children:
        dfs(ch)

def main():
    # Build tree: 1→{2,3}, 3→{5}
    root = TreeNodeDFS(1)
    root.children = [TreeNodeDFS(2), TreeNodeDFS(3)]
    root.children[1].children = [TreeNodeDFS(5)]
    dfs(root)
    print()

if __name__ == "__main__":
    main() 