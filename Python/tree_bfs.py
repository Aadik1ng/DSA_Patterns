"""
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
"""

from collections import deque

class TreeNode:
    def __init__(self, val):
        self.val = val
        self.children = []

def bfs(root):
    if not root:
        return
    q = deque([root])
    while q:
        node = q.popleft()
        print(node.val, end=' ')
        for ch in node.children:
            q.append(ch)

def main():
    # Build tree: 1→{2,3}, 2→{4}
    root = TreeNode(1)
    root.children = [TreeNode(2), TreeNode(3)]
    root.children[0].children = [TreeNode(4)]
    bfs(root)
    print()

if __name__ == "__main__":
    main() 