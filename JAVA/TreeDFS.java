import java.util.*;

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

class TreeNodeDFS {
    int val; List<TreeNodeDFS> children = new ArrayList<>();
    TreeNodeDFS(int v) { val = v; }
}

public class TreeDFS {
    // Recursive pre-order traversal
    public static void dfs(TreeNodeDFS node) {
        if (node == null) return;
        System.out.print(node.val + " ");
        for (TreeNodeDFS ch : node.children)
            dfs(ch);
    }
    public static void main(String[] args) {
        // Build tree: 1→{2,3}, 3→{5}
        TreeNodeDFS root = new TreeNodeDFS(1);
        root.children.add(new TreeNodeDFS(2));
        root.children.add(new TreeNodeDFS(3));
        root.children.get(1).children.add(new TreeNodeDFS(5));
        dfs(root);
        System.out.println();
    }
} 